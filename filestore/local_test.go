/*
 * @PackageName: filestore
 * @Description:
 * @Author: limuzhi
 * @Date: 2022/12/2 13:58
 */

package filestore

import (
	"io"
	"os"
	"reflect"
	"testing"
)

func TestLocalOption_Copy(t *testing.T) {
	type fields struct {
		config *LocalConfig
	}
	type args struct {
		srcKey  string
		destKey string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &LocalOption{
				config: tt.fields.config,
			}
			if err := o.Copy(tt.args.srcKey, tt.args.destKey); (err != nil) != tt.wantErr {
				t.Errorf("Copy() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestLocalOption_Delete(t *testing.T) {
	type fields struct {
		config *LocalConfig
	}
	type args struct {
		key string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &LocalOption{
				config: tt.fields.config,
			}
			if err := o.Delete(tt.args.key); (err != nil) != tt.wantErr {
				t.Errorf("Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestLocalOption_Exists(t *testing.T) {
	type fields struct {
		config *LocalConfig
	}
	type args struct {
		key string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    bool
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &LocalOption{
				config: tt.fields.config,
			}
			got, err := o.Exists(tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("Exists() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Exists() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLocalOption_Get(t *testing.T) {
	type fields struct {
		config *LocalConfig
	}
	type args struct {
		key string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    io.ReadCloser
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &LocalOption{
				config: tt.fields.config,
			}
			got, err := o.Get(tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Get() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLocalOption_GetDiskName(t *testing.T) {
	type fields struct {
		config *LocalConfig
	}
	tests := []struct {
		name   string
		fields fields
		want   DiskName
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &LocalOption{
				config: tt.fields.config,
			}
			if got := o.GetDiskName(); got != tt.want {
				t.Errorf("GetDiskName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLocalOption_Put(t *testing.T) {
	type fields struct {
		config *LocalConfig
	}
	type args struct {
		key         string
		r           io.Reader
		dataLength  int64
		contentType string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &LocalOption{
				config: tt.fields.config,
			}
			if err := o.Put(tt.args.key, tt.args.r, tt.args.dataLength, tt.args.contentType); (err != nil) != tt.wantErr {
				t.Errorf("Put() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestLocalOption_PutBase64(t *testing.T) {
	type fields struct {
		config *LocalConfig
	}
	type args struct {
		key         string
		r           []byte
		contentType string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &LocalOption{
				config: tt.fields.config,
			}
			if err := o.PutBase64(tt.args.key, tt.args.r, tt.args.contentType); (err != nil) != tt.wantErr {
				t.Errorf("PutBase64() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestLocalOption_PutFile(t *testing.T) {
	type fields struct {
		config *LocalConfig
	}
	type args struct {
		key         string
		localFile   string
		contentType string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &LocalOption{
				config: tt.fields.config,
			}
			if err := o.PutFile(tt.args.key, tt.args.localFile, tt.args.contentType); (err != nil) != tt.wantErr {
				t.Errorf("PutFile() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestLocalOption_Rename(t *testing.T) {
	type fields struct {
		config *LocalConfig
	}
	type args struct {
		srcKey  string
		destKey string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &LocalOption{
				config: tt.fields.config,
			}
			if err := o.Rename(tt.args.srcKey, tt.args.destKey); (err != nil) != tt.wantErr {
				t.Errorf("Rename() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestLocalOption_Size(t *testing.T) {
	type fields struct {
		config *LocalConfig
	}
	type args struct {
		key string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int64
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &LocalOption{
				config: tt.fields.config,
			}
			got, err := o.Size(tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("Size() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Size() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLocalOption_Url(t *testing.T) {
	type fields struct {
		config *LocalConfig
	}
	type args struct {
		key string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &LocalOption{
				config: tt.fields.config,
			}
			if got := o.Url(tt.args.key); got != tt.want {
				t.Errorf("Url() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLocalOption_getPath(t *testing.T) {
	type fields struct {
		config *LocalConfig
	}
	type args struct {
		key string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &LocalOption{
				config: tt.fields.config,
			}
			if got := o.getPath(tt.args.key); got != tt.want {
				t.Errorf("getPath() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLocalOption_openAsReadOnly(t *testing.T) {
	type fields struct {
		config *LocalConfig
	}
	type args struct {
		key string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *os.File
		want1   os.FileInfo
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &LocalOption{
				config: tt.fields.config,
			}
			got, got1, err := o.openAsReadOnly(tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("openAsReadOnly() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("openAsReadOnly() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("openAsReadOnly() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestNewLocalOption(t *testing.T) {
	type args struct {
		cfg *LocalConfig
	}
	tests := []struct {
		name    string
		args    args
		want    Storage
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewLocalOption(tt.args.cfg)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewLocalOption() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewLocalOption() got = %v, want %v", got, tt.want)
			}
		})
	}
}
