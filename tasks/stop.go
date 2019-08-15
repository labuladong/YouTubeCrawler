package tasks

import (
	"log"
	"youtube/interfaces"
	"youtube/utils"
)

type Finish struct {
	path string
}

func (t Finish) Execute() (error, interfaces.Task) {
	return nil, nil
}

func (t Finish) Report() {
	log.Println(utils.LastOfDirs(t.path) + "执行完成！")
}


