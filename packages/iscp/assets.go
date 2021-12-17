package iscp

import (
	"bytes"
	"fmt"
	"math/big"

	"golang.org/x/xerrors"

	"github.com/iotaledger/hive.go/marshalutil"
	"github.com/iotaledger/hive.go/serializer/v2"
	iotago "github.com/iotaledger/iota.go/v3"
	"github.com/iotaledger/wasp/packages/kv"
	"github.com/iotaledger/wasp/packages/kv/dict"
)

// Assets is used as assets in the UTXO and as tokens in transfer
type Assets struct {
	Iotas  uint64
	Tokens iotago.NativeTokens
}

var IotaAssetID = []byte{}

func NewEmptyAssets() *Assets {
	return &Assets{
		Tokens: make([]*iotago.NativeToken, 0),
	}
}

func NewAssets(iotas uint64, tokens iotago.NativeTokens) *Assets {
	return &Assets{
		Iotas:  iotas,
		Tokens: tokens,
	}
}

func NewAssetsFromDict(d dict.Dict) (*Assets, error) {
	ret := NewEmptyAssets()
	for key, val := range d {
		if IsIota([]byte(key)) {
			ret.Iotas = new(big.Int).SetBytes(d.MustGet(kv.Key(IotaAssetID))).Uint64()
			continue
		}
		id, err := NativeTokenIDFromBytes([]byte(key))
		if err != nil {
			return nil, xerrors.Errorf("NewAssetsFromDict: %w", err)
		}
		token := &iotago.NativeToken{
			ID:     id,
			Amount: new(big.Int).SetBytes(val),
		}
		ret.Tokens = append(ret.Tokens, token)
	}
	return ret, nil
}

func NewAssetsFromNativeTokenSum(iotas uint64, tokens iotago.NativeTokenSum) *Assets {
	ret := NewEmptyAssets()
	ret.Iotas = iotas
	for id, val := range tokens {
		ret.Tokens = append(ret.Tokens, &iotago.NativeToken{
			ID:     id,
			Amount: val,
		})
	}
	return ret
}

func AssetsFromOutput(iotago.Output) *Assets {
	panic("TODO implement")
}

func NativeTokenIDFromBytes(data []byte) (iotago.NativeTokenID, error) {
	if len(data) != iotago.NativeTokenIDLength {
		return iotago.NativeTokenID{}, xerrors.New("NativeTokenIDFromBytes: wrong data length")
	}
	var tokenID iotago.NativeTokenID
	copy(tokenID[:], data)
	return tokenID, nil
}

func MustNativeTokenIDFromBytes(data []byte) iotago.NativeTokenID {
	ret, err := NativeTokenIDFromBytes(data)
	if err != nil {
		panic(xerrors.Errorf("MustNativeTokenIDFromBytes: %w", err))
	}
	return ret
}

func (a *Assets) AmountNativeToken(tokenID *iotago.NativeTokenID) *big.Int {
	for _, t := range a.Tokens {
		if t.ID == *tokenID {
			return t.Amount
		}
	}
	return big.NewInt(0)
}

func (a *Assets) String() string {
	return fmt.Sprintf("Assets: iotas: %d tokens len: %d", a.Iotas, len(a.Tokens))
}

func (a *Assets) Bytes() []byte {
	mu := marshalutil.New()
	a.WriteToMarshalUtil(mu)
	return mu.Bytes()
}

func (a *Assets) Equals(b *Assets) bool {
	if a.Iotas != b.Iotas {
		return false
	}
	if len(a.Tokens) != len(b.Tokens) {
		return false
	}
	bTokensSet := b.Tokens.MustSet()
	for _, token := range a.Tokens {
		if token.Amount.Cmp(bTokensSet[token.ID].Amount) != 0 {
			return false
		}
	}
	return true
}

func (a *Assets) Add(b *Assets) *Assets {
	a.Iotas += b.Iotas
	resultTokens := a.Tokens.MustSet()
	for _, token := range b.Tokens {
		if resultTokens[token.ID] != nil {
			resultTokens[token.ID].Amount.Add(
				resultTokens[token.ID].Amount,
				token.Amount,
			)
			continue
		}
		resultTokens[token.ID] = token
	}
	a.Tokens = nativeTokensFromSet(resultTokens)
	return a
}

func (a *Assets) IsEmpty() bool {
	return a == nil || a.Iotas == 0 && len(a.Tokens) == 0
}

func (a *Assets) AddToken(tokenID iotago.NativeTokenID, amount *big.Int) *Assets {
	b := NewAssets(0, iotago.NativeTokens{
		&iotago.NativeToken{
			ID:     tokenID,
			Amount: amount,
		},
	})
	return a.Add(b)
}

//func (a *Assets) AddAsset(assetID []byte, amount *big.Int) *Assets {
//	switch len(assetID) {
//	case iotago.NativeTokenIDLength:
//		return a.AddToken(NativeTokenIDFromBytes(assetID), amount)
//	// TODO implement add NFTs
//	case len(IotaAssetID):
//		return a.AddIotas(amount.Uint64())
//	}
//	return a
//}

func (a *Assets) AddIotas(amount uint64) *Assets {
	a.Iotas += amount
	return a
}

func (a *Assets) ToDict() dict.Dict {
	ret := dict.New()
	ret.Set(kv.Key(IotaAssetID), new(big.Int).SetUint64(a.Iotas).Bytes())
	for _, token := range a.Tokens {
		ret.Set(kv.Key(token.ID[:]), token.Amount.Bytes())
	}
	return ret
}

func nativeTokensFromSet(set iotago.NativeTokensSet) iotago.NativeTokens {
	ret := make(iotago.NativeTokens, len(set))
	i := 0
	for _, token := range set {
		ret[i] = token
		i++
	}
	return ret
}

// IsIota return whether a given tokenID represents native Iotas
func IsIota(tokenID []byte) bool {
	return bytes.Equal(tokenID, IotaAssetID)
}

var NativeAssetsSerializationArrayRules = iotago.NativeTokenArrayRules()

func (a *Assets) WriteToMarshalUtil(mu *marshalutil.MarshalUtil) {
	mu.WriteUint64(a.Iotas)
	tokenBytes, err := serializer.NewSerializer().WriteSliceOfObjects(&a.Tokens, serializer.DeSeriModePerformLexicalOrdering, nil, serializer.SeriLengthPrefixTypeAsUint16, &NativeAssetsSerializationArrayRules, func(err error) error {
		return fmt.Errorf("unable to serialize alias output native tokens: %w", err)
	}).Serialize()
	if err != nil {
		panic(fmt.Errorf("unexpected error serializing native tokens: %w", err))
	}
	mu.WriteUint16(uint16(len(tokenBytes)))
	mu.WriteBytes(tokenBytes)
}

// TODO this could be refactored to use `AmountNativeToken`
// ToMap creates respective map by summing up repetitive token IDs
func FindNativeTokenBalance(nts iotago.NativeTokens, id *iotago.NativeTokenID) *big.Int {
	for _, nt := range nts {
		if nt.ID == *id {
			return nt.Amount
		}
	}
	return nil
}

func NewAssetsFromMarshalUtil(mu *marshalutil.MarshalUtil) (*Assets, error) {
	ret := &Assets{
		Tokens: make(iotago.NativeTokens, 0),
	}
	var err error
	if ret.Iotas, err = mu.ReadUint64(); err != nil {
		return nil, err
	}
	tokenBytesLength, err := mu.ReadUint16()
	if err != nil {
		return nil, err
	}
	tokenBytes, err := mu.ReadBytes(int(tokenBytesLength))
	if err != nil {
		return nil, err
	}
	_, err = serializer.NewDeserializer(tokenBytes).
		ReadSliceOfObjects(&ret.Tokens, serializer.DeSeriModePerformLexicalOrdering, nil, serializer.SeriLengthPrefixTypeAsUint16, serializer.TypeDenotationNone, &NativeAssetsSerializationArrayRules, func(err error) error {
			return fmt.Errorf("unable to deserialize native tokens for alias output: %w", err)
		}).Done()
	if err != nil {
		return nil, err
	}
	return ret, nil
}