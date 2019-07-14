package traverse

import (
	"fmt"
	"io/ioutil"
	"learn/youtube/exec"
	"log"
	"os"
	"strings"
)

func Rename(root string) {
	files, err := ioutil.ReadDir(root)
	if err != nil {
		log.Println(err)
	}
	sep := string(os.PathSeparator)
	//parts := strings.Split(root, sep)
	//name := parts[len(parts)-1]

	for _, file := range files {
		if file.IsDir() {
			Rename(root + sep + file.Name())
		} else if strings.HasSuffix(file.Name(), "mp4") || strings.HasSuffix(file.Name(), "srt") || strings.HasSuffix(file.Name(), "vtt") {
			s := strings.Split(file.Name(), ".")
			suffix := "." + s[len(s) - 1]
			cmd := fmt.Sprintf(`cd %s && mv "%s" %s`, root, file.Name(), "name" + suffix)
			println(cmd)
			_, err := exec.Shell(cmd)
			if err != nil {
				log.Println(err)
			}
		}
	}
}
