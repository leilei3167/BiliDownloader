package tool

import (
	"fmt"
	"os"
	"path/filepath"
)

// GetAidFileDownloadDir 创建文件夹,aid+title
func GetAidFileDownloadDir(aid int64, title string) string {
	curDir, err := os.Getwd() //TODO:学习获取当前的工作目录
	if err != nil {
		panic(err)
	}
	//组装成路径
	fullDirPath := filepath.Join(curDir, "download", fmt.Sprintf("%d_%s", aid, title))
	err = os.MkdirAll(fullDirPath, 0777)
	if err != nil {
		panic(err)
	}
	return fullDirPath
}
