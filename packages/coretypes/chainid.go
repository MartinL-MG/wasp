package coretypes

import (
	"bytes"
	"io"

	"github.com/iotaledger/goshimmer/dapps/valuetransfers/packages/address"
	"github.com/mr-tron/base58"
)

const ChainIDLength = address.Length

// ChainID epresents global identifier of the chain
// Currently it is an alias for the chain address
// In it will be and alias for chain color
type ChainID address.Address

var NilChainID = ChainID{}

// NewChainIDFromBase58 constructor unmarshals string
func NewChainIDFromBase58(b58 string) (ret ChainID, err error) {
	var b []byte
	b, err = base58.Decode(b58)
	if err != nil {
		return
	}
	if len(b) != ChainIDLength {
		err = ErrWrongDataLength
		return
	}
	copy(ret[:], b)
	return
}

// NewChainIDFromBytes constructor unmarshals bytes
func NewChainIDFromBytes(data []byte) (ret ChainID, err error) {
	err = ret.Read(bytes.NewReader(data))
	return
}

// NewRandomChainID constructor creates a random chain ID.
func NewRandomChainID() ChainID {
	return (ChainID)(address.RandomOfType(address.VersionBLS))
}

// Bytes returns a serialized version of this ChainID.
func (chid ChainID) Bytes() []byte {
	return chid.Address().Bytes()
}

// String human readable form
func (chid ChainID) String() string {
	return chid.Address().String()
}

func (chid ChainID) Address() address.Address {
	return address.Address(chid)
}

// Write marshal
func (chid *ChainID) Write(w io.Writer) error {
	_, err := w.Write(chid[:])
	return err
}

// Read unmarshal
func (chid *ChainID) Read(r io.Reader) error {
	n, err := r.Read(chid[:])
	if err != nil {
		return err
	}
	if n != ChainIDLength {
		return ErrWrongDataLength
	}
	return nil
}