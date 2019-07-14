package main

import (
	"learn/youtube/exec"
	"learn/youtube/traverse"
	"log"
	"os"
	"strings"
)

//var tokens = make(chan string, 4)
var done = make(chan bool)

func main() {
    urls, paths := traverse.FindURL("/media/fdl/ea0fb3dd-73b0-4979-8a74-9ac76271627c/已搬运/已完成/nqueens")
	tasks := make(chan [2]string)

	for i := 0; i < 5; i++ {
		go func() {
			for task := range tasks {
				url := task[0]
				path := task[1]

				dirs := strings.Split(path, string(os.PathSeparator))
				name := dirs[len(dirs)-2]
				log.Println(i, "# 开始：", name, url)
				exec.DownloadVideos(url, path)
				exec.TransformSub(path)
				exec.SetSubtitle(path)
				log.Println(i, "# 结束：", name, url)
				done <- true
			}

		}()
	}

	for i := 0; i < len(urls); i++ {
		go func(i int) {
			tasks <- [2]string{urls[i], paths[i]}
		}(i)
	}

	for i := 0; i < len(urls); i++ {
		<-done
	}

	//traverse.Rename("/media/fdl/ea0fb3dd-73b0-4979-8a74-9ac76271627c/已搬运")

	//exec.TransformSub("/media/fdl/ea0fb3dd-73b0-4979-8a74-9ac76271627c/youtube/01背包问题/")
}
