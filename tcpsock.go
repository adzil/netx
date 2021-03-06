// Copyright (c) 2018 Fadhli Dzil Ikram. All rights reserved.
// This software is licensed under the MIT license. Please see LICENSE file on
// the project root for more information.

package netx

import (
	"context"
	"net"
	"time"
)

// DialTimeoutTCP acts like DialTCP but takes a timeout.
func DialTimeoutTCP(network, laddr, raddr string, timeout time.Duration) (*net.TCPConn, error) {
	// Filter for TCP network only.
	switch network {
	case "tcp", "tcp4", "tcp6":
	default:
		return nil, &net.OpError{Op: "dial", Net: network, Err: net.UnknownNetworkError(network)}
	}
	// Resolve local address if defined.
	var localAddr net.Addr
	if len(laddr) > 0 {
		var err error
		if localAddr, err = net.ResolveTCPAddr(network, laddr); err != nil {
			return nil, err
		}
	}
	// Use net.Dialer to Dial TCP.
	dialer := net.Dialer{Control: ReusePort, LocalAddr: localAddr, Timeout: timeout}
	conn, err := dialer.DialContext(context.Background(), network, raddr)
	if err != nil {
		return nil, err
	}
	// Return TCP connection.
	return conn.(*net.TCPConn), nil
}

// DialTCP acts like Dial for TCP networks.
//
// If laddr is empty, a local address is automatically chosen.
func DialTCP(network, laddr, raddr string) (*net.TCPConn, error) {
	return DialTimeoutTCP(network, laddr, raddr, 0)
}

// ListenTCP acts like Listen for TCP networks.
func ListenTCP(network, address string) (*net.TCPListener, error) {
	// Filter for TCP network only.
	switch network {
	case "tcp", "tcp4", "tcp6":
	default:
		return nil, &net.OpError{Op: "listen", Net: network, Err: net.UnknownNetworkError(network)}
	}
	// Use net.ListenConfig to listen TCP.
	listenConfig := net.ListenConfig{Control: ReusePort}
	conn, err := listenConfig.Listen(context.Background(), network, address)
	if err != nil {
		return nil, err
	}
	return conn.(*net.TCPListener), nil
}
