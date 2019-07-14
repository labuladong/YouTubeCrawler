package exec

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
)

func DownloadVideos(url string, path string) {
	downloadCmd := fmt.Sprintf(
		"cd %s && youtube-dl -i --write-auto-sub --write-thumbnail --proxy socks5://127.0.0.1:1080/ %s", path, url)
	println(downloadCmd)
	out, err := Shell(downloadCmd)
	if err != nil {
		log.Println("下载视频出错！", err)
	}
	writeLog(out, path)
}

func writeLog(out, path string) {
	logCmd := fmt.Sprintf("cd %s && touch log && echo \"%s\" > log", path, out)
	_, err := Shell(logCmd)
	if err != nil {
		log.Println("生成 log 出错!", err)
	}
}

//阻塞式的执行外部shell命令的函数,等待执行完毕并返回标准输出
func Shell(s string) (string, error) {
	//函数返回一个*Cmd，用于使用给出的参数执行name指定的程序
	cmd := exec.Command("/bin/bash", "-c", s)

	//读取io.Writer类型的cmd.Stdout，再通过bytes.Buffer(缓冲byte类型的缓冲器)将byte类型转化为string类型(out.String():这是bytes类型提供的接口)
	var out bytes.Buffer
	cmd.Stdout = &out

	//Run执行c包含的命令，并阻塞直到完成。  这里stdout被取出，cmd.Wait()无法正确获取stdin,stdout,stderr，则阻塞在那了
	err := cmd.Run()

	return out.String(), err
}
