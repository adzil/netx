// Copyright (c) 2018 Fadhli Dzil Ikram. All rights reserved.
// This software is licensed under the MIT license. Please see LICENSE file on
// the project root for more information.

// Package netx enables port reuse feature on the network transport and provides
// transparent network address reflection interface through Session Traversal
// Utilities for NAT (STUN).
package netx

import (
	"net"
	"time"
)

// Listen announces on the local network address with port reuse enabled.
//
// The network must be "tcp", "tcp4", or "tcp6".
//
// See func net.Listen for further details.
func Listen(network, address string) (net.Listener, error) {
	return ListenTCP(network, address)
}

// ListenPacket announces on the local network address with port reuse enabled.
//
// The network must be "udp", "udp4", or "udp6".
//
// See func net.Listen for further details.
func ListenPacket(network, address string) (net.PacketConn, error) {
	return ListenUDP(network, address)
}

// DialTimeout acts like Dial but takes a timeout.
//
// The timeout includes name resolution, if required. When using TCP, and the
// host in the address parameter resolves to multiple IP addresses, the timeout
// is spread over each consecutive dial, such that each is given an appropriate
// fraction of the time to connect.
//
// See func net.Dial for a description of the network and address parameters.
func DialTimeout(network, localAddress, address string, timeout time.Duration) (net.Conn, error) {
	switch network {
	case "tcp", "tcp4", "tcp6":
		return DialTimeoutTCP(network, localAddress, address, timeout)
	case "udp", "udp4", "udp6":
		return DialTimeoutUDP(network, localAddress, address, timeout)
	}
	return nil, &net.OpError{Op: "listen", Net: network, Err: net.UnknownNetworkError(network)}
}

// Dial connects to the address on the named network with port reuse enabled.
//
// The network must be "tcp", "tcp4", "tcp6", "udp", "udp4", or "udp6".
//
// If localAddress is empty, a local address is automatically chosen.
//
// See func net.Dial for a description of the network and address parameters.
func Dial(network, localAddress, address string) (net.Conn, error) {
	return DialTimeout(network, localAddress, address, 0)
}
