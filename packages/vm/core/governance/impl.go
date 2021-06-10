package governance

import (
	"github.com/iotaledger/wasp/packages/coretypes"
	"github.com/iotaledger/wasp/packages/coretypes/assert"
	"github.com/iotaledger/wasp/packages/kv/collections"
	"github.com/iotaledger/wasp/packages/kv/dict"
	"github.com/iotaledger/wasp/packages/kv/kvdecoder"
)

func initialize(ctx coretypes.Sandbox) (dict.Dict, error) {
	return nil, nil
}

// checkRotateCommitteeRequest the entry point is called when committee is about to be rotated to the new address
// If it fails, nothing happens and the state has trace of the failure in the state
// If it is successful VM takes over and replaces resulting transaction with
// governance transition. The state of the chain remains unchanged
func checkRotateCommitteeRequest(ctx coretypes.Sandbox) (dict.Dict, error) {
	a := assert.NewAssert(ctx.Log())
	a.RequireChainOwner(ctx, "checkRotateCommitteeRequest")
	par := kvdecoder.New(ctx.Params(), ctx.Log())
	addr := par.MustGetAddress(ParamStateAddress)
	amap := collections.NewMap(ctx.State(), StateVarAllowedCommitteeAddresses)
	a.Require(amap.MustHasAt(addr.Bytes()), "checkRotateCommitteeRequest: address is not allowed as next state address: %s", addr.Base58())
	return nil, nil
}

func addAllowedCommitteeAddress(ctx coretypes.Sandbox) (dict.Dict, error) {
	a := assert.NewAssert(ctx.Log())
	a.RequireChainOwner(ctx, "addAllowedCommitteeAddress")
	par := kvdecoder.New(ctx.Params(), ctx.Log())
	addr := par.MustGetAddress(ParamStateAddress)
	amap := collections.NewMap(ctx.State(), StateVarAllowedCommitteeAddresses)
	amap.MustSetAt(addr.Bytes(), []byte{0xFF})
	return nil, nil
}

func removeAllowedCommitteeAddress(ctx coretypes.Sandbox) (dict.Dict, error) {
	a := assert.NewAssert(ctx.Log())
	a.RequireChainOwner(ctx, "removeAllowedCommitteeAddress")
	par := kvdecoder.New(ctx.Params(), ctx.Log())
	addr := par.MustGetAddress(ParamStateAddress)
	amap := collections.NewMap(ctx.State(), StateVarAllowedCommitteeAddresses)
	amap.MustDelAt(addr.Bytes())
	return nil, nil
}

func isAllowedCommitteeAddress(ctx coretypes.SandboxView) (dict.Dict, error) {
	par := kvdecoder.New(ctx.Params(), ctx.Log())
	addr := par.MustGetAddress(ParamStateAddress)
	allowed := MustIsAllowedCommitteeAddress(ctx.State(), addr)
	ret := dict.New()
	if allowed {
		ret.Set(ParamIsAllowedAddress, []byte{0xFF})
	}
	return ret, nil
}

func moveToAddress(ctx coretypes.Sandbox) (dict.Dict, error) {
	ctx.Log().Panicf("moveToAddress: not implemented")
	return nil, nil
}