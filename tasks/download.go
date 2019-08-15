package tasks

import (
	"errors"
	"fmt"
	"log"
	"youtube/interfaces"
	"youtube/utils"
)

type DownloadTask struct{
	path string
	url string
	msg string
}


func (t DownloadTask) Report() {
	log.Println("开始下载视频：" + utils.LastOfDirs(t.path))
}

func NewDownloadTask(path string, url string) *DownloadTask {
	return &DownloadTask{path: path, url: url}
}

func (t DownloadTask) Execute() (error, interfaces.Task) {
	downloadCmd := fmt.Sprintf(
		"cd %s && youtube-dl -i --write-auto-sub --write-thumbnail --proxy socks5://127.0.0.1:1080/ %s", t.path, t.url)
	//println(downloadCmd)
	_, err := utils.Shell(downloadCmd)
	if err != nil {
		return errors.New("下载视频出错！" + err.Error()), nil
	}
	return nil, OptSubtitle{path:t.path}
}
