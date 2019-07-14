package exec

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func SetSubtitle(path string) {
	sep := string(os.PathSeparator)
	dirs := strings.Split(path, sep)
	var name string
	var subSrc string
	var videoSrc string

	// 目录名作为视频名字
	i := len(dirs) - 1
	for i >= 0 && dirs[i] == "" {
		i--
	}
	name = dirs[i]

	infos, err := ioutil.ReadDir(path)
	if err != nil {
		log.Println("压制字幕出错!", err)
		return
	}
	// 找到字幕文件和视频文件
	for _, info := range infos {
		suffix := getSuffix(info.Name())
		if (suffix == "mp4" || suffix == "mkv") && !strings.HasSuffix(name, "mp4"){
			name += "." + suffix
			videoSrc = info.Name()
		} else if suffix == "srt" {
			subSrc = info.Name()
		} else if suffix == "vtt" && subSrc == "" {
			subSrc = info.Name()
		}
	}

	if name == "" || subSrc == "" || videoSrc == "" {
		log.Println("setSubtitle: 相关文件查找出错！")
		return
	}

	mergeCmd := fmt.Sprintf(`cd %s && ffmpeg -i "%s" -vf subtitles=%s %s`, path, videoSrc, subSrc, name)
	println(mergeCmd)
	out, err := Shell(mergeCmd)
	if err != nil {
		log.Println("ffmepg 执行出错！", err)
	}
	writeLog(out, path)
}


func getSuffix(file string) string {
	parts := strings.Split(file, ".")
	if len(parts) < 2 {
		return ""
	}
	return parts[len(parts)-1]
}
