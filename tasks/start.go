package tasks

import (
	"log"
	"youtube/interfaces"
	"youtube/utils"
)

type StartTask struct{
	path, url string
}

func NewStartTask(path string, url string) *StartTask {
	return &StartTask{path: path, url: url}
}

func (t StartTask) Execute() (error, interfaces.Task) {
	return nil, NewDownloadTask(t.path, t.url)
}

func (t StartTask) Report() {
	log.Println("初始化任务: " + utils.LastOfDirs(t.path))
}

