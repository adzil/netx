// Copyright (c) 2018 Fadhli Dzil Ikram. All rights reserved.
// This software is licensed under the MIT license. Please see LICENSE file on
// the project root for more information.

package netx

import "syscall"

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
