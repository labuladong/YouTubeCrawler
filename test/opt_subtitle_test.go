package test

import (
	"testing"
	"youtube/tasks"
)

func TestOptSubtitle(t *testing.T) {
	task := tasks.NewOptSubtitle("/media/fdl/ea0fb3dd-73b0-4979-8a74-9ac76271627c/youtube/linux好用的命令")
	err, _ := task.Execute()
	if err != nil {
		task.Report()
		t.Error(err)
	}
	task.Report()
}
