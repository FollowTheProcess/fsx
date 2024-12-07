// Package fsx provides a clean, simple filesystem abstraction allowing you to use a consistent
// API regardless of the underlying implementation whether it's the real OS filesystem or a memory
// backed test filesystem.
package fsx

import "path/filepath"

const (
	// DefaultDirPermissions are the default UNIX permission bits for a new directory, they
	// are the equivalent permissions to the result of running mkdir.
	DefaultDirPermissions = 0o755

	// DefaultFilePermissions are the default UNIX permission bits for a new file, they
	// are the equivalent permissions to the result of running touch.
	DefaultFilePermissions = 0o666

	// Separator is the OS-specific filepath separator.
	Separator = string(filepath.Separator)
)

// FileSystem represents an abstract filesystem in a storage-agnostic interface.
type FileSystem interface {
	// Create creates a new named file (truncating it if it already exists).
	Create(name string) (File, error)
}

// File represents a single named file in the filesystem.
type File interface {
	// TODO(@FollowTheProcess): Uncomment these and implement one at a time
	// io.Reader
	// io.ReaderAt
	// io.Writer
	// io.WriterAt
	// io.Seeker
	// io.Closer

	// TODO(@FollowTheProcess): Add the rest of the methods from os.File

	// Name returns the name of the file as presented to Open.
	//
	// It is safe to call Name after Close.
	Name() string

	// Exists reports whether the file exists in the filesystem.
	Exists() bool

	// Truncate truncates the file to a specific size.
	Truncate(size int64) error
}
