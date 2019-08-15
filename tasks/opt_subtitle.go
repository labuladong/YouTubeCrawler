package tasks

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
	"youtube/interfaces"
	"youtube/utils"
)

type OptSubtitle struct {
	path string
	msg  string
}

func NewOptSubtitle(path string) *OptSubtitle {
	return &OptSubtitle{path: path}
}

func (t OptSubtitle) Execute() (error, interfaces.Task) {
	infos, err := ioutil.ReadDir(t.path)
	if err != nil {
		err = errors.New("转化字幕格式：打开目标位置出错！\n" + err.Error())
		return err, nil
	}

	var srtFile string
	//借助 ffmpeg 转化成 srt
	for _, info := range infos {
		file := info.Name()
		if strings.HasSuffix(file, "vtt") {
			parts := strings.Split(file, ".")
			parts[len(parts)-1] = "srt"
			srtFile = strings.Join(parts, ".")

			if !utils.CheckFileIsExist(t.path + srtFile) {
				//先删除一下原有的文件，以免失败，以便后续任务执行
				_, _ = utils.Shell(fmt.Sprintf("cd %s && rm -f %s", t.path, srtFile))
				// 用 ffmpeg 开始转化
				cmd := fmt.Sprintf(`cd %s && ffmpeg -i "%s" "%s"`, t.path, file, srtFile)
				_, err := utils.Shell(cmd)
				if err != nil {
					err = errors.New("转化字幕格式：ffmpeg 命令出错！" + cmd)
					return err, nil
				}
			}
			err := utils.RemoveDuplicate(t.path + "/" + srtFile)
			if err != nil {
				err = errors.New("字幕文件去重出现问题" + err.Error())
				return err, nil
			}
		}
	}
	return nil, SetSubtitleTask{path: t.path, subname:srtFile}
}

func (t OptSubtitle) Report() {
	log.Println(t.msg)
}
