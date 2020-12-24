// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package wasmimpl

import (
	"github.com/iotaledger/goshimmer/dapps/valuetransfers/packages/balance"
	"github.com/iotaledger/wasp/packages/coretypes"
	"github.com/iotaledger/wasp/packages/vm/wasmhost"
)

type ScTransfer struct {
	ScDict
	agent coretypes.AgentID
}

func (o *ScTransfer) Exists(keyId int32) bool {
	return o.GetTypeId(keyId) > 0
}

func (o *ScTransfer) GetTypeId(keyId int32) int32 {
	switch keyId {
	case wasmhost.KeyAgent:
		return wasmhost.OBJTYPE_BYTES //TODO OBJTYPE_AGENT
	}
	return wasmhost.OBJTYPE_INT
}

func (o *ScTransfer) SetBytes(keyId int32, value []byte) {
	var err error
	switch keyId {
	case wasmhost.KeyAgent:
		o.agent, err = coretypes.NewAgentIDFromBytes(value)
		if err != nil {
			panic("Invalid agent: " + err.Error())
		}
	default:
		o.ScDict.SetBytes(keyId, value)
	}
}

func (o *ScTransfer) SetInt(keyId int32, value int64) {
	switch keyId {
	case wasmhost.KeyLength:
		o.agent = coretypes.AgentID{}
	default:
		key := o.vm.GetKeyFromId(keyId)
		color, _, err := balance.ColorFromBytes(key)
		if err != nil {
			panic("Invalid color: " + err.Error())
		}
		o.Trace("TRANSFER #%d c'%s' a'%s'", value, color.String(), o.agent.String())
		if !o.vm.ctx.MoveTokens(o.agent, color, value) {
			panic("Failed to move tokens")
		}
	}
}

// \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\ // \\

type ScTransfers struct {
	ScDict
}

func (a *ScTransfers) GetObjectId(keyId int32, typeId int32) int32 {
	return GetArrayObjectId(a, keyId, typeId, func() WaspObject {
		return &ScTransfer{}
	})
}

func (a *ScTransfers) SetString(keyId int32, value string) {
	a.Panic("SetString: Invalid access")
}