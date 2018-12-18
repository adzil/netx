// Copyright (c) 2018 Fadhli Dzil Ikram. All rights reserved.
// This software is licensed under the MIT license. Please see LICENSE file on
// the project root for more information.

package netx

import (
	"context"
	"net"
	"time"
)

// DialTimeoutUDP acts like DialUDP but takes a timeout.
func DialTimeoutUDP(network, localAddress, address string, timeout time.Duration) (*net.UDPConn, error) {
	// Filter for UDP network only.
	switch network {
	case "udp", "udp4", "udp6":
	default:
		return nil, &net.OpError{Op: "dial", Net: network, Err: net.UnknownNetworkError(network)}
	}
	// Resolve local address if defined.
	var localAddr net.Addr
	if len(localAddress) > 0 {
		var err error
		if localAddr, err = net.ResolveUDPAddr(network, localAddress); err != nil {
			return nil, err
		}
	}
	// Use net.Dialer to Dial UDP.
	dialer := net.Dialer{Control: ControlFunc, LocalAddr: localAddr, Timeout: timeout}
	conn, err := dialer.DialContext(context.Background(), network, address)
	if err != nil {
		return nil, err
	}
	// Return UDP connection.
	return conn.(*net.UDPConn), nil
}

// DialUDP acts like Dial for UDP networks.
func DialUDP(network, localAddress, address string) (*net.UDPConn, error) {
	return DialTimeoutUDP(network, localAddress, address, 0)
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
	listenConfig := net.ListenConfig{Control: ControlFunc}
	conn, err := listenConfig.ListenPacket(context.Background(), network, address)
	if err != nil {
		return nil, err
	}
	return conn.(*net.UDPConn), nil
}
