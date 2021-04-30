package service

import (
	"os"
	"ppIm/lib"
	"ppIm/utils"
)

func UploadToQiNiu(uploadPath string, localPath string) error {
	// 七牛云上传
	err := utils.QiNiuClient.Upload(localPath, uploadPath)
	// 删除本地缓存
	os.Remove(localPath)

	if err != nil {
		lib.Logger.Debugf(err.Error())
		return err
	}
	return nil
}
