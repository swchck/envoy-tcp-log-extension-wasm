package main

import (
	"github.com/tetratelabs/proxy-wasm-go-sdk/proxywasm"
	"github.com/tetratelabs/proxy-wasm-go-sdk/proxywasm/types"
)

func main() {
	proxywasm.SetVMContext(&vmContext{})
}

type vmContext struct {
	types.DefaultVMContext
}

// NewPluginContext Override types.DefaultVMContext.
func (*vmContext) NewPluginContext(contextID uint32) types.PluginContext {
	return &pluginContext{}
}

type pluginContext struct {
	types.DefaultPluginContext
}

// NewTcpContext Override types.DefaultPluginContext.
func (ctx *pluginContext) NewTcpContext(contextID uint32) types.TcpContext {
	return &networkContext{}
}

type networkContext struct {
	types.DefaultTcpContext
}

// OnNewConnection Override types.DefaultTcpContext.
func (ctx *networkContext) OnNewConnection() types.Action {
	sourceAddress, _ := proxywasm.GetProperty([]string{"source", "address"})
	proxywasm.LogInfof("New Connection with address: %s!", sourceAddress)
	return types.ActionContinue
}

// OnDownstreamClose Override types.DefaultTcpContext.
func (ctx *networkContext) OnDownstreamClose(types.PeerType) {
	sourceAddress, _ := proxywasm.GetProperty([]string{"source", "address"})
	proxywasm.LogInfof("Downstream connection with address: %s closed!", sourceAddress)
	return
}

// OnStreamDone Override types.DefaultTcpContext.
func (ctx *networkContext) OnStreamDone() {
	sourceAddress, _ := proxywasm.GetProperty([]string{"source", "address"})
	proxywasm.LogInfof("Connection with address: %s closed!", sourceAddress)
}
