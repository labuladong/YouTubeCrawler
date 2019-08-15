package utils

import (
	"io/ioutil"
	"log"
	"os"
	"strings"
	"youtube/config"
)

func LastOfDirs(path string) string {
	dirs := strings.Split(path, config.Sep)
	last := dirs[len(dirs) - 1]
	return last
}

func FindURL(root string) ([]string, []string) {
	files, err := ioutil.ReadDir(root)
	if err != nil {
		log.Println(err)
	}
	sep := string(os.PathSeparator)
	var urls []string
	var paths []string

	for _, file := range files {
		if file.IsDir() {
			tempUrls, tempPaths := FindURL(root + sep + file.Name())
			urls = append(urls, tempUrls...)
			paths = append(paths, tempPaths...)
		} else if file.Name() == "url" {
			bytes, err := ioutil.ReadFile(root + sep + file.Name())
			if err != nil {
				log.Println(err)
				continue
			}
			//log.Println("reach url file", root + sep)

			urls = append(urls, strings.TrimSpace(string(bytes)))
			paths = append(paths, root)
		}
	}
	return urls, paths
}
