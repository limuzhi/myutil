/*
 * @PackageName: filestore
 * @Description:
 * @Author: limuzhi
 * @Date: 2022/12/1 15:39
 */

package filestore

import (
	"context"
	"crypto/tls"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	qiniuStorage "github.com/qiniu/go-sdk/v7/storage"
	"io"
	"net/http"
	"time"
)

type QiniuyunConfig struct {
	AccessKey string
	Bucket    string
	Domain    string
	SecretKey string
	IsSsl     bool
	IsPrivate bool
}

type QiniuYunOption struct {
	config         *QiniuyunConfig
	putPolicy      *qiniuStorage.PutPolicy
	mac            *qbox.Mac
	formUploader   *qiniuStorage.FormUploader
	base64Uploader *qiniuStorage.Base64Uploader
	bucketManager  *qiniuStorage.BucketManager
}

func NewQiniuYunOption(cfg *QiniuyunConfig) (Storage, error) {
	var (
		o = &QiniuYunOption{}
	)
	o.config = cfg
	o.putPolicy = &qiniuStorage.PutPolicy{
		Scope: cfg.Bucket,
	}
	o.mac = qbox.NewMac(cfg.AccessKey, cfg.SecretKey)
	config := &qiniuStorage.Config{
		UseHTTPS:      cfg.IsSsl,
		UseCdnDomains: false,
	}
	o.formUploader = qiniuStorage.NewFormUploader(config)
	o.base64Uploader = qiniuStorage.NewBase64Uploader(config)
	o.bucketManager = qiniuStorage.NewBucketManager(o.mac, config)
	return o, nil
}

func (o *QiniuYunOption) GetDiskName() DiskName {
	return KoDo
}

func (o *QiniuYunOption) Put(key string, r io.Reader, dataLength int64, contentType string) error {
	key = NormalizeKey(key)
	upToken := o.putPolicy.UploadToken(o.mac)
	ret := qiniuStorage.PutRet{}
	putExtra := qiniuStorage.PutExtra{MimeType: contentType}
	err := o.formUploader.Put(context.Background(), &ret, upToken, key, r, dataLength, &putExtra)
	if err != nil {
		return err
	}
	return nil
}

func (o *QiniuYunOption) PutBase64(key string, r []byte, contentType string) error {
	key = NormalizeKey(key)

	upToken := o.putPolicy.UploadToken(o.mac)
	ret := qiniuStorage.PutRet{}
	putExtra := qiniuStorage.Base64PutExtra{MimeType: contentType}
	err := o.base64Uploader.Put(context.Background(), &ret, upToken, key, r, &putExtra)
	if err != nil {
		return err
	}
	return nil
}

func (o *QiniuYunOption) PutFile(key string, localFile string, contentType string) error {
	key = NormalizeKey(key)

	upToken := o.putPolicy.UploadToken(o.mac)
	ret := qiniuStorage.PutRet{}
	putExtra := qiniuStorage.PutExtra{MimeType: contentType}
	err := o.formUploader.PutFile(context.Background(), &ret, upToken, key, localFile, &putExtra)
	if err != nil {
		return err
	}

	return nil
}

func (o *QiniuYunOption) Get(key string) (io.ReadCloser, error) {
	key = NormalizeKey(key)

	req, err := http.NewRequest("GET", o.Url(key), nil)
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return resp.Body, nil
}

func (o *QiniuYunOption) Rename(srcKey string, destKey string) error {
	srcKey = NormalizeKey(srcKey)
	destKey = NormalizeKey(destKey)

	err := o.bucketManager.Move(o.config.Bucket, srcKey, o.config.Bucket, destKey, true)
	if err != nil {
		return err
	}

	return nil
}

func (o *QiniuYunOption) Copy(srcKey string, destKey string) error {
	srcKey = NormalizeKey(srcKey)
	destKey = NormalizeKey(destKey)

	err := o.bucketManager.Copy(o.config.Bucket, srcKey, o.config.Bucket, destKey, true)
	if err != nil {
		return err
	}

	return nil
}

func (o *QiniuYunOption) Exists(key string) (bool, error) {
	key = NormalizeKey(key)

	_, err := o.bucketManager.Stat(o.config.Bucket, key)
	if err != nil {
		if err.Error() == "no such file or directory" {
			err = nil
		}
		return false, err
	}
	return true, nil
}

func (o *QiniuYunOption) Size(key string) (int64, error) {
	key = NormalizeKey(key)

	fileInfo, err := o.bucketManager.Stat(o.config.Bucket, key)
	if err != nil {
		return 0, err
	}

	return fileInfo.Fsize, nil
}

func (o *QiniuYunOption) Delete(key string) error {
	key = NormalizeKey(key)

	err := o.bucketManager.Delete(o.config.Bucket, key)
	if err != nil {
		return err
	}

	return nil
}

func (o *QiniuYunOption) Url(key string) string {
	var prefix string

	key = NormalizeKey(key)

	if o.config.IsSsl {
		prefix = "https://"
	} else {
		prefix = "http://"
	}

	if o.config.IsPrivate {
		deadline := time.Now().Add(time.Second * 3600).Unix() // 1小时有效期
		return prefix + qiniuStorage.MakePrivateURL(o.mac, o.config.Domain, key, deadline)
	}

	return prefix + qiniuStorage.MakePublicURL(o.config.Domain, key)
}
