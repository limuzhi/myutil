/*
 * @PackageName: oss
 * @Description:
 * @Author: limuzhi
 * @Date: 2022/12/1 15:25
 */

package filestore

import (
	"fmt"
)

type OptionFunc func(*OssOptions)

type OssOptions struct {
	disabled bool
	storages map[DiskName]Storage
}

func NewWithOptions(opts ...OptionFunc) *OssOptions {
	o := &OssOptions{
		disabled: false,
		storages: make(map[DiskName]Storage, 0),
	}
	for _, opt := range opts {
		opt(o)
	}
	return o
}

func WithDisabled(disabled bool) OptionFunc {
	return func(o *OssOptions) {
		o.disabled = disabled
	}
}

func WithStorage(services ...Storage) OptionFunc {
	return func(o *OssOptions) {
		for _, s := range services {
			if s != nil {
				name := s.GetDiskName()
				if _, ok := o.storages[name]; !ok {
					o.storages[name] = s
				}
			}
		}
	}
}

func (o OssOptions) GetStorage(name DiskName) (Storage, error) {
	storage, ok := o.storages[name]
	if !ok {
		return nil, fmt.Errorf("storage: Unknown disk %q", name)
	}
	return storage, nil
}

func (o OssOptions) Register(name DiskName, disk Storage) {
	if disk == nil {
		panic("storage: Register disk is nil")
	}
	o.storages[name] = disk
}
