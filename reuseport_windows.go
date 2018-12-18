// Copyright (c) 2018 Fadhli Dzil Ikram. All rights reserved.
// This software is licensed under the MIT license. Please see LICENSE file on
// the project root for more information.

package netx

import (
	"syscall"
)

// reusePort enables reuse port on Windows.
func reusePort(network, address string, c syscall.RawConn) (err error) {
	// SO_REUSEADDR option is equivalent to SO_REUSEPORT in Unix.
	if ctrlErr := c.Control(func(fd uintptr) {
		if err = syscall.SetsockoptInt(syscall.Handle(fd), syscall.SOL_SOCKET, syscall.SO_REUSEADDR, 1); err != nil {
			return
		}
	}); ctrlErr != nil {
		return ctrlErr
	}
	return
}
