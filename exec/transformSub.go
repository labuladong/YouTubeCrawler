package exec

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

//寻找 path 中的 vtt 字幕转化成 srt 并切除多余行
func TransformSub(path string) {
	//println(path)
	infos, err := ioutil.ReadDir(path)
	if err != nil {
		log.Println("转化字幕格式：打开目标位置出错！", err)
		return
	}

	var srtFile string
	//借助 ffmpeg 转化成 srt
	for _, info := range infos {
		file := info.Name()
		if strings.HasSuffix(file, "vtt") {
			parts := strings.Split(file, ".")
			parts[len(parts)-1] = "srt"
			srtFile = strings.Join(parts, ".")

			if !checkFileIsExist(path + srtFile) {
				// 用 ffmpeg 开始转化
				cmd := fmt.Sprintf(`cd %s && ffmpeg -i "%s" "%s"`, path, file, srtFile)
				//println(cmd)
				_, err := Shell(cmd)
				if err != nil {
					log.Println("转化字幕格式：ffmpeg 命令出错！", err)
					return
				}
			}

			//if checkFileIsExist(path + "new.srt") {
			//	log.Println(srtFile, "已经转换过！")
			//	return
			//}
			removeDuplicate(path + srtFile)
		}
	}

}

func removeDuplicate(path string) {
	reader, err := os.OpenFile(path, os.O_RDONLY, os.ModeAppend)
	if err != nil {
		log.Println("创建字幕文件失败！", err)
		return
	}
	defer reader.Close()

	sep := string(filepath.Separator)
	parts := strings.Split(path, sep)
	parts[len(parts)-1] = "new.srt"
	newPath := strings.Join(parts, sep)
	fileObj, err := os.OpenFile(newPath, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		log.Println("创建去重文件失败！", err)
	}
	defer fileObj.Close()

	seen := make(map[string]bool)
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			WriteWithIo(fileObj, "\n")
		} else if seen[text] == false {
			seen[text] = true
			WriteWithIo(fileObj, text)
		}
	}
	log.Println(parts[len(parts) - 2], "去重成功！")
}

//使用io.WriteString()函数进行数据的写入
func WriteWithIo(file *os.File, text string) {
	_, err := file.WriteString(text + "\n")
	if err != nil {
		log.Println("去重字幕写入失败：", text, err)
	}
}

/**
 * 判断文件是否存在  存在返回 true 不存在返回false
 */
func checkFileIsExist(filename string) (bool) {
	var exist = true;
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		exist = false;
	}
	return exist;
}
