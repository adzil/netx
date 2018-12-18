// Copyright (c) 2018 Fadhli Dzil Ikram. All rights reserved.
// This software is licensed under the MIT license. Please see LICENSE file on
// the project root for more information.

// +build aix darwin dragonfly freebsd linux netbsd openbsd

package netx

import (
	"syscall"

	"golang.org/x/sys/unix"
)

// reusePort enables reuse port on Unix.
func reusePort(network, address string, c syscall.RawConn) (err error) {
	// Enable SO_REUSEPORT on socket.
	if ctrlErr := c.Control(func(fd uintptr) {
		if err = syscall.SetsockoptInt(int(fd), syscall.SOL_SOCKET, unix.SO_REUSEPORT, 1); err != nil {
			return
		}
	}); ctrlErr != nil {
		return ctrlErr
	}
	return
}
