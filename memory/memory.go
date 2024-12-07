// Package memory provides a FileSystem implementation backed by a in-memory storage mechanism
// that does not require any access to the underlying OS. All File IO operations happen in memory
// making this FileSystem ideal for testing.
package memory

import (
	"fmt"
	"io/fs"
	"path/filepath"

	"github.com/FollowTheProcess/fsx"
	"github.com/FollowTheProcess/fsx/internal/flag"
)

// Compile time interface checks, these will fail to compile if we ever violate our interfaces.
var (
	_ fsx.File       = &file{}
	_ fsx.FileSystem = &InMemory{}
)

// InMemory is an implementation of an [fsx.FileSystem] backed by an in-memory
// storage backend.
type InMemory struct {
	// The files in the filesystem, mapped by their full path (with os-specific separator).
	files map[string]*file
}

// New creates and returns a new [fsx.FileSystem] backed by an in-memory storage system
// so that file io operations can be performed as normal but without touching the real OS filesystem.
func New() fsx.FileSystem {
	return &InMemory{
		// Create the root
		files: map[string]*file{
			fsx.Separator: {
				name:   fsx.Separator,
				mode:   fsx.DefaultDirPermissions | fs.ModeDir,
				exists: true,
			},
		},
	}
}

// Create creates a new named file (truncating it if it already exists).
func (i *InMemory) Create(name string) (fsx.File, error) {
	name = filepath.Clean(filepath.FromSlash(name))
	path := filepath.Join(fsx.Separator, name)
	f, exists := i.files[path]
	if exists {
		// If it exists, truncate it
		if err := f.Truncate(0); err != nil {
			return nil, fmt.Errorf("could not truncate %s: %w", path, err)
		}
		i.files[path] = f
		return f, nil
	}

	f = &file{
		name:   name,
		flag:   flag.ReadWrite | flag.Create | flag.Truncate,
		mode:   fsx.DefaultFilePermissions,
		exists: true,
	}
	i.files[path] = f

	return f, nil
}

// file represents a single named file in an in-memory filesystem.
type file struct {
	name    string      // The name of the file
	content []byte      // The contents of the file
	flag    int         // The OS file flags
	mode    fs.FileMode // The file's mode and permission bits
	exists  bool        // Whether the file has been created into the filesystem
}

// Name returns the name of the file.
func (f *file) Name() string {
	return f.name
}

// Exists reports whether the file exists in the filesystem.
func (f *file) Exists() bool {
	return f.exists
}

// Truncate truncates the file to a specific size.
func (f *file) Truncate(size int64) error {
	if size < int64(len(f.content)) {
		f.content = f.content[:size]
	} else if more := int(size) - len(f.content); more > 0 {
		f.content = append(f.content, make([]byte, more)...)
	}

	return nil
}
