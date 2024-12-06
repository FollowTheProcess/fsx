// Package fsx provides a clean, simple filesystem abstraction allowing you to use a consistent
// API regardless of the underlying implementation whether it's the real OS filesystem or a memory
// backed test filesystem.
package fsx

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

	// Name returns the name of the file as presented to Open.
	//
	// It is safe to call Name after Close.
	Name() string

	// Exists reports whether the file exists in the filesystem.
	Exists() bool

	// Truncate truncates the file to a specific size.
	Truncate(size int64) error
}
