/*
 * @PackageName: filestore
 * @Description:
 * @Author: limuzhi
 * @Date: 2022/12/1 16:47
 */

package filestore

import (
	"encoding/base64"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"io"
	"strconv"
	"strings"
)

type AliyunConfig struct {
	AccessKeyId     string
	AccessKeySecret string
	Bucket          string
	Endpoint        string
	IsSsl           bool
	IsPrivate       bool
}
type AliyunOssOption struct {
	config *AliyunConfig
	client *oss.Client
	bucket *oss.Bucket
}

func NewAliYunOssOption(cfg *AliyunConfig) (Storage, error) {
	var (
		err error
		o   = &AliyunOssOption{}
	)
	o.config = cfg
	o.client, err = oss.New(cfg.Endpoint, cfg.AccessKeyId, cfg.AccessKeySecret)
	if err != nil {
		return nil, err
	}
	o.bucket, err = o.client.Bucket(cfg.Bucket)
	return o, err
}

func (o *AliyunOssOption) GetDiskName() DiskName {
	return Oss
}

func (o *AliyunOssOption) Put(key string, r io.Reader, dataLength int64, contentType string) error {
	key = NormalizeKey(key)
	err := o.bucket.PutObject(key, r, oss.ContentType(contentType))
	if err != nil {
		return err
	}
	return nil
}

func (o *AliyunOssOption) PutBase64(key string, r []byte, contentType string) error {
	key = NormalizeKey(key)
	//获取图片内容并base64解密
	base64Img := string(r)
	fileContentPosition := strings.Index(base64Img, ",")
	uploadBaseString := base64Img[fileContentPosition+1:]
	uploadString, _ := base64.StdEncoding.DecodeString(uploadBaseString)
	err := o.bucket.PutObject(key, strings.NewReader(string(uploadString)), oss.ContentType(contentType))
	if err != nil {
		return err
	}
	return nil
}

func (o *AliyunOssOption) PutFile(key string, localFile string, contentType string) error {
	key = NormalizeKey(key)
	err := o.bucket.PutObjectFromFile(key, localFile, oss.ContentType(contentType))
	if err != nil {
		return err
	}
	return nil
}

func (o *AliyunOssOption) Get(key string) (io.ReadCloser, error) {
	key = NormalizeKey(key)

	body, err := o.bucket.GetObject(key)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func (o *AliyunOssOption) Rename(srcKey string, destKey string) error {
	srcKey = NormalizeKey(srcKey)
	destKey = NormalizeKey(destKey)

	_, err := o.bucket.CopyObject(srcKey, destKey)
	if err != nil {
		return err
	}

	err = o.Delete(srcKey)
	if err != nil {
		return err
	}

	return nil
}

func (o *AliyunOssOption) Copy(srcKey string, destKey string) error {
	srcKey = NormalizeKey(srcKey)
	destKey = NormalizeKey(destKey)

	_, err := o.bucket.CopyObject(srcKey, destKey)
	if err != nil {
		return err
	}
	return nil
}

func (o *AliyunOssOption) Exists(key string) (bool, error) {
	key = NormalizeKey(key)
	return o.bucket.IsObjectExist(key)
}

func (o *AliyunOssOption) Size(key string) (int64, error) {
	key = NormalizeKey(key)

	props, err := o.bucket.GetObjectDetailedMeta(key)
	if err != nil {
		return 0, err
	}

	size, err := strconv.ParseInt(props.Get("Content-Length"), 10, 64)
	if err != nil {
		return 0, err
	}
	return size, nil
}

func (o *AliyunOssOption) Delete(key string) error {
	key = NormalizeKey(key)
	err := o.bucket.DeleteObject(key)
	if err != nil {
		return err
	}
	return nil
}
func (o *AliyunOssOption) Url(key string) string {
	var prefix string
	key = NormalizeKey(key)

	if o.config.IsSsl {
		prefix = "https://"
	} else {
		prefix = "http://"
	}

	if o.config.IsPrivate {
		url, err := o.bucket.SignURL(key, oss.HTTPGet, 3600)
		if err == nil {
			return url
		}
	}
	return prefix + o.config.Bucket + "." + o.config.Endpoint + "/" + key
}
