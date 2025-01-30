package task_test

import (
	"io"
	"os"
	"testing"

	"github.com/go-task/task/v3"
	"github.com/go-task/task/v3/internal/filepathext"
)

func TestInit(t *testing.T) {
	t.Parallel()

	const dir = "testdata/init"
	file := filepathext.SmartJoin(dir, "Taskfile.yml")

	_ = os.Remove(file)
	if _, err := os.Stat(file); err == nil {
		t.Errorf("Taskfile.yml should not exist")
	}

	if err := task.InitTaskfile(io.Discard, file); err != nil {
		t.Error(err)
	}

	if _, err := os.Stat(file); err != nil {
		t.Errorf("Taskfile.yml should exist")
	}
	_ = os.Remove(file)
}
