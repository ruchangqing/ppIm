package services

import (
	"context"
	"errors"
	"fmt"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
)

type QiNiu struct {
	AccessKey string
	SecretKey string
	Bucket    string
	Domain    string
}

var QiNiuClient QiNiu

func (q QiNiu) Upload(localFile string, fileName string) error {
	putPolicy := storage.PutPolicy{
		Scope: q.Bucket,
	}
	mac := qbox.NewMac(q.AccessKey, q.SecretKey)
	upToken := putPolicy.UploadToken(mac)
	cfg := storage.Config{}
	// 空间对应的机房
	cfg.Zone = &storage.Zone_z2
	// 是否使用https域名
	cfg.UseHTTPS = true
	// 上传是否使用CDN上传加速
	cfg.UseCdnDomains = false
	// 构建表单上传的对象
	formUploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}
	// 可选配置
	putExtra := storage.PutExtra{
		Params: map[string]string{
			"x:name": "ppIm",
		},
	}

	err := formUploader.PutFile(context.Background(), &ret, upToken, fileName, localFile, &putExtra)
	if err != nil {
		fmt.Println("七牛云上传失败：" + err.Error())
		return errors.New("上传失败")
	}
	return nil
}

// 补全文件地址
func (q QiNiu) FullPath(filePath string) string {
	if filePath[0:4] != "http" {
		return QiNiuClient.Domain + "/" + filePath
	} else {
		return filePath
	}
}
