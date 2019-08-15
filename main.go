package main

import (
	"youtube/config"
	"youtube/engine"
	"youtube/interfaces"
	"youtube/tasks"
	"youtube/utils"
)

func main() {
	urls, paths := utils.FindURL(config.Pwd)
	e := engine.NewSimpleEngine(3)
	var seeds []interfaces.Task;
	for i := 0; i < len(urls); i++ {
		seeds = append(seeds, tasks.NewStartTask(paths[i], urls[i]))
	}
	e.Run(seeds...)
}