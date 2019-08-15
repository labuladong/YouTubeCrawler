package utils

import (
	"bufio"
	"errors"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func RemoveDuplicate(path string) (error){
	reader, err := os.OpenFile(path, os.O_RDONLY, os.ModeAppend)
	if err != nil {
		return errors.New("创建字幕文件失败！" + err.Error())
	}
	defer reader.Close()

	sep := string(filepath.Separator)
	parts := strings.Split(path, sep)
	parts[len(parts)-1] = "new.srt"
	newPath := strings.Join(parts, sep)
	fileObj, err := os.OpenFile(newPath, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return errors.New("创建去重文件失败！" + err.Error())
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
	return nil
}

//使用io.WriteString()函数进行数据的写入
func WriteWithIo(file *os.File, text string) {
	_, err := file.WriteString(text + "\n")
	if err != nil {
		log.Println("去重字幕写入失败：", text, err)
	}
}

/**
 * 判断文件是否存在  存在返回 true 不存在返回 false
 */
func CheckFileIsExist(filename string) (bool) {
	var exist = true;
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		exist = false;
	}
	return exist;
}
