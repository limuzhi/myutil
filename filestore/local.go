/*
 * @PackageName: filestore
 * @Description: 本地存储
 * @Author: limuzhi
 * @Date: 2022/12/2 10:51
 */

package filestore

import (
	"encoding/base64"
	"io"
	"os"
	"path/filepath"
	"strings"
)

type LocalConfig struct {
	RootDir string
	AppUrl  string
}

type LocalOption struct {
	config *LocalConfig
}

func NewLocalOption(cfg *LocalConfig) (Storage, error) {
	return &LocalOption{config: cfg}, nil
}

func (o *LocalOption) getPath(key string) string {
	key = NormalizeKey(key)
	return filepath.Join(o.config.RootDir, key)
}

func (o *LocalOption) openAsReadOnly(key string) (*os.File, os.FileInfo, error) {
	fd, err := os.Open(key)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, nil, FileNotFoundErr
		}
		if os.IsPermission(err) {
			return nil, nil, FileNoPermissionErr
		}
		return nil, nil, err
	}

	stat, err := fd.Stat()
	if err != nil {
		return nil, nil, err
	}

	return fd, stat, nil
}

func (o *LocalOption) GetDiskName() DiskName {
	return Local
}

func (o *LocalOption) Put(key string, r io.Reader, dataLength int64, contentType string) error {
	path := o.getPath(key)
	dir, _ := filepath.Split(path)
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		return err
	}

	fd, err := os.OpenFile(path, os.O_CREATE|os.O_RDWR|os.O_TRUNC, os.ModePerm)
	if err != nil {
		if os.IsPermission(err) {
			return FileNoPermissionErr
		}
		return err
	}
	defer fd.Close()
	_, err = io.Copy(fd, r)
	return err
}

func (o *LocalOption) PutBase64(key string, r []byte, contentType string) error {
	//获取图片内容并base64解密
	base64Img := string(r)
	fileContentPosition := strings.Index(base64Img, ",")
	uploadBaseString := base64Img[fileContentPosition+1:]
	uploadString, _ := base64.StdEncoding.DecodeString(uploadBaseString)
	uploadReader := strings.NewReader(string(uploadString))
	if len(contentType) == 0 {
		contentType = base64Img[:fileContentPosition]
		contentType = strings.ReplaceAll(contentType, "data:", "")
		contentType = strings.ReplaceAll(contentType, ";base64", "")
	}
	return o.Put(key, uploadReader, int64(len(uploadString)), contentType)
}

func (o *LocalOption) PutFile(key string, localFile string, contentType string) error {
	path := o.getPath(localFile)

	fd, fileInfo, err := o.openAsReadOnly(path)
	if err != nil {
		return err
	}
	defer fd.Close()

	return o.Put(key, fd, fileInfo.Size(), contentType)
}

func (o *LocalOption) Get(key string) (io.ReadCloser, error) {
	path := o.getPath(key)

	fd, _, err := o.openAsReadOnly(path)
	if err != nil {
		return nil, err
	}
	return fd, nil
}

func (o *LocalOption) Rename(srcKey string, destKey string) error {
	srcPath := o.getPath(srcKey)
	ok, err := o.Exists(srcPath)
	if err != nil {
		return err
	}
	if !ok {
		return FileNotFoundErr
	}

	destPath := o.getPath(destKey)
	dir, _ := filepath.Split(destPath)
	if err = os.MkdirAll(dir, os.ModePerm); err != nil {
		return err
	}

	return os.Rename(srcPath, destPath)
}

func (o *LocalOption) Copy(srcKey string, destKey string) error {
	srcPath := o.getPath(srcKey)
	srcFd, _, err := o.openAsReadOnly(srcPath)
	if err != nil {
		return err
	}
	defer srcFd.Close()

	destPath := o.getPath(destKey)
	dir, _ := filepath.Split(destPath)
	if err = os.MkdirAll(dir, os.ModePerm); err != nil {
		return err
	}
	destFd, err := os.OpenFile(destPath, os.O_CREATE|os.O_RDWR|os.O_TRUNC, os.ModePerm)
	if err != nil {
		if os.IsPermission(err) {
			return FileNoPermissionErr
		}
		return err
	}
	defer destFd.Close()

	_, err = io.Copy(destFd, srcFd)
	return err
}

func (o *LocalOption) Exists(key string) (bool, error) {
	path := o.getPath(key)
	_, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}
		if os.IsPermission(err) {
			return false, FileNoPermissionErr
		}
		return false, err
	}
	return true, nil
}

func (o *LocalOption) Size(key string) (int64, error) {
	path := o.getPath(key)
	fileInfo, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return 0, FileNotFoundErr
		}
		if os.IsPermission(err) {
			return 0, FileNoPermissionErr
		}
		return 0, err
	}

	return fileInfo.Size(), nil
}

func (o *LocalOption) Delete(key string) error {
	path := o.getPath(key)
	err := os.Remove(path)
	if err != nil {
		if os.IsNotExist(err) {
			return FileNotFoundErr
		}
		if os.IsPermission(err) {
			return FileNoPermissionErr
		}
		return err
	}
	return nil
}

func (o *LocalOption) Url(key string) string {
	return o.config.AppUrl + "/" + NormalizeKey(key)
}
