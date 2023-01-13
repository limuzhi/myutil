/*
 * @PackageName: filestore
 * @Description:
 * @Author: limuzhi
 * @Date: 2022/12/1 15:28
 */

package filestore

type DiskName string

const (
	Local DiskName = "local" // 本地
	KoDo  DiskName = "kodo"  // 七牛云
	Oss   DiskName = "oss"   // 阿里云
	Minio DiskName = "minio" // minio
)
