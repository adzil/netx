// Copyright (c) 2018 Fadhli Dzil Ikram. All rights reserved.
// This software is licensed under the MIT license. Please see LICENSE file on
// the project root for more information.

package netx

import (
	"context"
	"net"
	"time"
)

// UDPConn is the remote bind address reflection interface for net.UDPConn.
type UDPConn struct {
	*udpConn                // The actual UDP connection.
	stun       *net.UDPConn // Connection to STUN server.
	remoteAddr *net.UDPAddr // Reflected remote address.
}

// RemoteAddr returns the reflected remote bind address. The Addr returned is
// shared by all invocations of RemoteAddr, so do not modify it.
func (u *UDPConn) RemoteAddr() net.Addr {
	return u.remoteAddr
}

// Close closes the connection.
func (u *UDPConn) Close() error {
	// Close underlying STUN connection to prevent leak.
	u.stun.Close()
	return u.udpConn.Close()
}

// DialTimeoutUDP acts like DialUDP but takes a timeout.
func DialTimeoutUDP(network, laddr, raddr string, timeout time.Duration) (*net.UDPConn, error) {
	// Filter for UDP network only.
	switch network {
	case "udp", "udp4", "udp6":
	default:
		return nil, &net.OpError{Op: "dial", Net: network, Err: net.UnknownNetworkError(network)}
	}
	// Resolve local address if defined.
	var localAddr net.Addr
	if len(laddr) > 0 {
		var err error
		if localAddr, err = net.ResolveUDPAddr(network, laddr); err != nil {
			return nil, err
		}
	}
	// Use net.Dialer to Dial UDP.
	dialer := net.Dialer{Control: ReusePort, LocalAddr: localAddr, Timeout: timeout}
	conn, err := dialer.DialContext(context.Background(), network, raddr)
	if err != nil {
		return nil, err
	}
	// Return UDP connection.
	return conn.(*net.UDPConn), nil
}

// DialUDP acts like Dial for UDP networks.
//
// If laddr is empty, a local address is automatically chosen.
func DialUDP(network, laddr, raddr string) (*net.UDPConn, error) {
	return DialTimeoutUDP(network, laddr, raddr, 0)
}

// ListenUDP acts like Listen for UDP networks.
func ListenUDP(network, address string) (*net.UDPConn, error) {
	// Filter for UDP network only.
	switch network {
	case "udp", "udp4", "udp6":
	default:
		return nil, &net.OpError{Op: "listen", Net: network, Err: net.UnknownNetworkError(network)}
	}
	// Use net.ListenConfig to listen UDP.
	listenConfig := net.ListenConfig{Control: ReusePort}
	conn, err := listenConfig.ListenPacket(context.Background(), network, address)
	if err != nil {
		return nil, err
	}
	return conn.(*net.UDPConn), nil
}

// ListenRemoteUDP acts like ListenUDP but with remote bind address reflection.
func ListenRemoteUDP(network, localAddress, stunAddress string) (*UDPConn, error) {
	// Create UDP listener on localAddress.
	conn, err := ListenUDP(network, localAddress)
	if err != nil {
		return nil, err
	}
	// Create separate connection to STUN server with same local address.
	stunConn, err := DialTimeoutUDP(network, conn.LocalAddr().String(), stunAddress, 5*time.Second)
	if err != nil {
		return nil, err
	}
	// TODO: Use stunConn to get mapped address from STUN server and spawn
	// heartbeat goroutine.
	_ = stunConn
	// Return active UDPConn.
	return &UDPConn{
		udpConn:    conn,
		stun:       stunConn,
		remoteAddr: nil,
	}, nil
}
