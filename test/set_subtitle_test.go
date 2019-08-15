package test

import (
	"testing"
	"youtube/tasks"
)

func TestSetSubtitle(t *testing.T) {
	task := tasks.NewSetSubtitleTask("/media/fdl/ea0fb3dd-73b0-4979-8a74-9ac76271627c/youtube/linux好用的命令", "new.srt")
	e, _ := task.Execute()
	if e != nil {
		task.Report()
		t.Error(e)
	}
	task.Report()
}
