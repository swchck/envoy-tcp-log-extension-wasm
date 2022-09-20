//go:build proxytest

package main

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/tetratelabs/proxy-wasm-go-sdk/proxywasm/proxytest"
	"github.com/tetratelabs/proxy-wasm-go-sdk/proxywasm/types"
)

func TestNetwork_OnNewConnection(t *testing.T) {
	opt := proxytest.NewEmulatorOption().WithVMContext(&vmContext{})
	host, reset := proxytest.NewHostEmulator(opt)
	defer reset()

	// Initialize plugin
	require.Equal(t, types.OnPluginStartStatusOK, host.StartPlugin())

	// OnNewConnection is called.
	_, action := host.InitializeConnection()
	require.Equal(t, types.ActionContinue, action)

	// Check Envoy logs.
	logs := host.GetInfoLogs()
	require.Contains(t, logs, "New Connection with address: !")
}

func TestNetwork_OnDownstreamClose(t *testing.T) {
	opt := proxytest.NewEmulatorOption().WithVMContext(&vmContext{})
	host, reset := proxytest.NewHostEmulator(opt)
	defer reset()

	// OnNewConnection is called.
	contextID, action := host.InitializeConnection()
	require.Equal(t, types.ActionContinue, action)

	// OnDownstreamClose is called.
	host.CloseDownstreamConnection(contextID)

	// Check Envoy logs.
	logs := host.GetInfoLogs()
	require.Contains(t, logs, "Downstream connection with address:  closed!")
}
