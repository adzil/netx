// Copyright (c) 2018 Fadhli Dzil Ikram. All rights reserved.
// This software is licensed under the MIT license. Please see LICENSE file on
// the project root for more information.

// +build !aix,!darwin,!dragonfly,!freebsd,!linux,!netbsd,!openbsd,!windows

package netx

import (
	"errors"
	"syscall"
)

// Disable reusePort on unsupported OS.
func reusePort(network, address string, c syscall.RawConn) error {
	return errors.New("netx: reuse port is unavailable")
}
