package real_test

import (
	"os"
	"testing"

	"github.com/FollowTheProcess/fsx/real"
	"github.com/FollowTheProcess/test"
)

func TestCreate(t *testing.T) {
	fs := real.New()

	file, err := fs.Create("test")
	test.Ok(t, err)
	defer os.RemoveAll("test")

	test.Equal(t, file.Name(), "test")
	test.Equal(t, file.Exists(), true)
}
