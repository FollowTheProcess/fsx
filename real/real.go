// Package real provides a FileSystem implementation backed by the OS for interacting
// with the real, physical file system in an OS agnostic way.
package real //nolint: predeclared // I've never used whatever 'real' is

import (
	"os"

	"github.com/FollowTheProcess/fsx"
)

// Compile time interface checks, these will fail to compile if we ever violate our interfaces.
var (
	_ fsx.File       = &file{}
	_ fsx.FileSystem = Real{}
)

// Real is an implementation of an [fsx.FileSystem] backed by the operating system.
type Real struct{}

// New creates and returns a new [fsx.FileSystem] backed by the operating system.
func New() fsx.FileSystem {
	return Real{}
}

// Create creates a new named file (truncating it if it already exists).
func (r Real) Create(name string) (fsx.File, error) {
	f, err := os.Create(name)
	if err != nil {
		return nil, err
	}

	file := &file{inner: f}

	return file, nil
}

// file is a thin wrapper around an *os.File to satisfy the [fsx.File] interface.
type file struct {
	inner *os.File
}

// Name returns the name of the file.
func (f *file) Name() string {
	return f.inner.Name()
}

// Exists reports whether the file exists in the filesystem.
func (f *file) Exists() bool {
	_, err := os.Stat(f.inner.Name())
	return err == nil
}

// Truncate truncates the file to a specific si ze.
func (f *file) Truncate(size int64) error {
	return f.inner.Truncate(size)
}
