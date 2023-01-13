/*
 * @PackageName: oss
 * @Description:
 * @Author: limuzhi
 * @Date: 2022/12/1 15:23
 */

package filestore

import "io"

type Storage interface {
	Put(key string, r io.Reader, dataLength int64, contentType string) error
	PutBase64(key string, r []byte, contentType string) error
	PutFile(key string, localFile string, contentType string) error
	Get(key string) (io.ReadCloser, error)
	Rename(srcKey string, destKey string) error
	Copy(srcKey string, destKey string) error
	Exists(key string) (bool, error)
	Size(key string) (int64, error)
	Delete(key string) error
	Url(key string) string
	GetDiskName() DiskName
}
