/*
 * @PackageName: filestore
 * @Description:
 * @Author: limuzhi
 * @Date: 2022/12/2 11:00
 */

package filestore

import "errors"

var (
	FileNotFoundErr     = errors.New("file not found")
	FileNoPermissionErr = errors.New("permission denied")
)
