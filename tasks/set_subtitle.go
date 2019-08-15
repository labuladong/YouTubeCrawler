package tasks

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
	"youtube/config"
	"youtube/interfaces"
	"youtube/utils"
)

type SetSubtitleTask struct {
	path string
	subname string
	msg  string
}

func NewSetSubtitleTask(path string, subname string) *SetSubtitleTask {
	return &SetSubtitleTask{path: path, subname: subname}
}

func (t SetSubtitleTask) Report() {
	log.Println(t.msg)
}

func (t SetSubtitleTask) Execute() (error, interfaces.Task) {
	dirs := strings.Split(t.path, config.Sep)
	videoName := dirs[len(dirs) - 1] + ".mp4"

	//if utils.CheckFileIsExist(t.path + config.Sep + videoName) {
	//	utils.Shell("rm -f " + t.path + config.Sep + videoName)
	//}

	subSrc := t.subname
	var videoSrc string

	infos, err := ioutil.ReadDir(t.path)
	if err != nil {
		err = errors.New("读取文件夹位置出错：" + err.Error())
		return err, nil
	}
	// 找到字幕文件和视频文件
	for _, info := range infos {
		name := info.Name()
		//文件的后缀
		suffix := name[strings.Index(name, ".") + 1:]
		if config.VideoType[suffix] {
			videoSrc = name
		}
	}

	if videoName == "" || subSrc == "" || videoSrc == "" {
		err = errors.New("未找到字幕或视频文件！")
		return err, nil
	}

	mergeCmd := fmt.Sprintf(`cd %s && ffmpeg -i "%s" -vf subtitles="%s" %s`, t.path, videoSrc, subSrc, videoName)
	_, err = utils.Shell(mergeCmd)

	if err != nil {
		err = errors.New("ffmepg 执行出错: " + mergeCmd)
		return err, nil
	}
	return nil, Finish{path:t.path}
}



