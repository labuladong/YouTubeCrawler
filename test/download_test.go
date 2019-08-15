package test

import (
	"testing"
	"youtube/tasks"
)
var path = "/media/fdl/ea0fb3dd-73b0-4979-8a74-9ac76271627c/youtube/test"

func TestDownloader(t *testing.T) {
	downloader := tasks.NewDownloadTask(path, "https://www.youtube.com/watch?v=mXYZi3AXcBc")
	e, _ := downloader.Execute()
	if e != nil {
		t.Error(e)
	}
	downloader.Report()
	//e, task = task.Execute()
}
