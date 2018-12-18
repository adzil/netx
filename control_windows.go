// Copyright (c) 2018 Fadhli Dzil Ikram. All rights reserved.
// This software is licensed under the MIT license. Please see LICENSE file on
// the project root for more information.

package netx

import (
	"syscall"
)

// controlFunc enables port reuse feature on Windows.
func controlFunc(network, address string, c syscall.RawConn) (err error) {
	// Enable SO_REUSEADDR bit.
	if ctrlErr := c.Control(func(fd uintptr) {
		if err = syscall.SetsockoptInt(syscall.Handle(fd), syscall.SOL_SOCKET, syscall.SO_REUSEADDR, 1); err != nil {
			return
		}
	}); ctrlErr != nil {
		return ctrlErr
	}
	return
}
