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

func init() {
	// Test whether port reuse feature is available or not.
	listenConfig := net.ListenConfig{Control: ControlFunc}
	conn, err := listenConfig.ListenPacket(context.Background(), "udp", "")
	if err != nil {
		return
	}
	// Create new listener on the same local address.
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

// ControlFunc is the default control function for netx to enable the port reuse
// feature. It can be directly embedded to net.ListenConfig or net.Dialer.
func ControlFunc(network, address string, c syscall.RawConn) error {
	// Only control TCP and UDP socket.
	switch network {
	case "tcp", "tcp4", "tcp6", "udp", "udp4", "udp6":
	default:
		return nil
	}
	// Call OS-specific function to enable port reuse.
	return controlFunc(network, address, c)
}
