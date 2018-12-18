// Copyright (c) 2018 Fadhli Dzil Ikram. All rights reserved.
// This software is licensed under the MIT license. Please see LICENSE file on
// the project root for more information.

package netx

import (
	"context"
	"net"
	"syscall"
)

// Internal port reusable feature status.
var isPortReusable bool

// Init function to test whether port reuse feature is available or not.
func init() {
	listenConfig := net.ListenConfig{Control: ReusePort}
	conn, err := listenConfig.ListenPacket(context.Background(), "udp", "")
	if err != nil {
		return
	}
	nconn, err := listenConfig.ListenPacket(context.Background(), "udp", conn.LocalAddr().String())
	conn.Close()
	if err != nil {
		return
	}
	nconn.Close()
	isPortReusable = true
}

// IsPortReusable checks whether port reuse feature is available or not.
func IsPortReusable() bool {
	return isPortReusable
}

// ReusePort used by netx to enable reuse port feature by default. It also can
// be used on the Control field in net.ListenConfig or net.Dialer to create
// custom ListenConfig or Dialer with reuse port capability.
//
//     lc := net.ListenConfig{Control: netx.ReusePort}
//     listener, err := lc.Listen("tcp", ":3000")
//
// ReusePort only available on TCP and UDP connection.
func ReusePort(network, address string, c syscall.RawConn) error {
	// Only control TCP and UDP socket.
	switch network {
	case "tcp", "tcp4", "tcp6", "udp", "udp4", "udp6":
	default:
		return nil
	}
	// Call OS-specific implementation function.
	return reusePort(network, address, c)
}
