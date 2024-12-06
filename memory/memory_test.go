package memory_test

import (
	"testing"

	"github.com/FollowTheProcess/fsx/memory"
	"github.com/FollowTheProcess/test"
)

func TestCreate(t *testing.T) {
	t.Run("new", func(t *testing.T) {
		mem := memory.New()
		file, err := mem.Create("test")
		test.Ok(t, err)

		test.Equal(t, file.Name(), "test")
		test.Equal(t, file.Exists(), true)
	})
	t.Run("exists", func(t *testing.T) {
		mem := memory.New()
		file, err := mem.Create("test")
		test.Ok(t, err)

		another, err := mem.Create("test")
		test.Ok(t, err)

		test.Equal(t, file.Name(), "test")
		test.Equal(t, file.Exists(), true)

		test.Equal(t, another.Name(), "test")
		test.Equal(t, another.Exists(), true)
	})
}
