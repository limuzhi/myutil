/*
 * @PackageName: filestore
 * @Description: https://www.minio.org.cn/
 * @Author: limuzhi
 * @Date: 2022/12/2 8:48
 */

package filestore

import (
	"context"
	"crypto/tls"
	"encoding/base64"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type MinioConfig struct {
	AccessKeyId     string
	AccessKeySecret string
	Bucket          string
	Endpoint        string
	IsSsl           bool
	IsPrivate       bool
}

type MinioOption struct {
	config *MinioConfig
	client *minio.Client
}

func NewMinioOption(cfg *MinioConfig) (Storage, error) {
	var (
		o   = &MinioOption{}
		err error
	)
	o.config = cfg
	o.client, err = minio.New(cfg.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(cfg.AccessKeyId, cfg.AccessKeySecret, ""),
		Secure: cfg.IsSsl,
	})
	return o, err
}

func (o *MinioOption) GetDiskName() DiskName {
	return Minio
}

func (o *MinioOption) Put(key string, r io.Reader, dataLength int64, contentType string) error {
	key = NormalizeKey(key)
	_, err := o.client.PutObject(context.Background(), o.config.Bucket, key, r, dataLength,
		minio.PutObjectOptions{ContentType: contentType})
	return err
}

func (o *MinioOption) PutBase64(key string, r []byte, contentType string) error {
	key = NormalizeKey(key)
	//获取图片内容并base64解密
	base64Img := string(r)
	fileContentPosition := strings.Index(base64Img, ",")
	uploadBaseString := base64Img[fileContentPosition+1:]
	uploadString, _ := base64.StdEncoding.DecodeString(uploadBaseString)

	_, err := o.client.PutObject(context.Background(), o.config.Bucket, key, strings.NewReader(string(uploadString)), int64(len(uploadString)),
		minio.PutObjectOptions{ContentType: contentType})
	return err
}

func (o *MinioOption) PutFile(key string, localFile string, contentType string) error {
	key = NormalizeKey(key)
	_, err := o.client.FPutObject(context.Background(), o.config.Bucket, key, localFile, minio.PutObjectOptions{ContentType: contentType})
	return err
}

func (o *MinioOption) Get(key string) (io.ReadCloser, error) {
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

func (o *MinioOption) Rename(srcKey string, destKey string) error {
	err := o.Copy(srcKey, destKey)
	if err != nil {
		return err
	}
	err = o.Delete(srcKey)
	return err
}

//文档：https://min.io/docs/minio/linux/developers/go/API.html#CopyObject

func (o *MinioOption) Copy(srcKey string, destKey string) error {
	srcKey = NormalizeKey(srcKey)
	destKey = NormalizeKey(destKey)

	srcOpts := minio.CopySrcOptions{
		Bucket: o.config.Bucket,
		Object: srcKey,
	}

	// Destination object
	dstOpts := minio.CopyDestOptions{
		Bucket: o.config.Bucket,
		Object: destKey,
	}
	_, err := o.client.CopyObject(context.Background(), dstOpts, srcOpts)
	return err
}

func (o *MinioOption) Exists(key string) (bool, error) {
	key = NormalizeKey(key)
	_, err := o.client.GetObject(context.Background(), o.config.Bucket, key, minio.GetObjectOptions{})
	if err != nil {
		return false, err
	}
	return true, nil
}

func (o *MinioOption) Size(key string) (int64, error) {
	key = NormalizeKey(key)
	object, err := o.client.GetObject(context.Background(), o.config.Bucket, key, minio.GetObjectOptions{})
	if err != nil {
		return 0, err
	}
	var b []byte
	size, err := object.Read(b)
	if err != nil {
		return 0, err
	}
	return int64(size), nil
}

func (o *MinioOption) Delete(key string) error {
	key = NormalizeKey(key)
	opts := minio.RemoveObjectOptions{
		GovernanceBypass: true,
	}
	err := o.client.RemoveObject(context.Background(), o.config.Bucket, key, opts)

	return err
}

func (o *MinioOption) Url(key string) string {
	key = NormalizeKey(key)
	reqParams := make(url.Values)
	presignedURL, err := o.client.PresignedGetObject(context.Background(), o.config.Bucket, key, time.Second*24*60*60, reqParams)
	if err != nil {
		return ""
	}
	return presignedURL.String()
}
