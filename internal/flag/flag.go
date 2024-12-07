// Package flag provides constants for OS file/dir flags.
//
// They are here in an internal package so we can reuse them across fsx but they are not
// exposed in the public API.
package flag

import "os"

const (
	Create    = os.O_CREATE
	Exclusive = os.O_EXCL
	Append    = os.O_APPEND
	Truncate  = os.O_TRUNC
	ReadWrite = os.O_RDWR
	ReadOnly  = os.O_RDONLY
	WriteOnly = os.O_WRONLY
)
