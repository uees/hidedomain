package utils

import (
	"log"
	"os"
	"path"
)

func InitLoger() error {
	file := path.Join(BaseDir(), Conf.LogFile)
	logFile, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0766)
	if err != nil {
		return err
	}
	log.SetOutput(logFile) // 将文件设置为log输出的文件
	// log.SetPrefix("[LOG]")
	log.SetFlags(log.LstdFlags | log.Lshortfile | log.LUTC)
	return nil
}
