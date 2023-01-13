/*
 * @PackageName: filestore
 * @Description:
 * @Author: limuzhi
 * @Date: 2022/12/2 13:03
 */

package filestore

import (
	"bytes"
	"github.com/minio/minio-go/v7"
	"io"
	"io/ioutil"
	"reflect"
	"testing"
)

func TestMinioOption_Copy(t *testing.T) {
	type fields struct {
		config *MinioConfig
		client *minio.Client
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
			o := &MinioOption{
				config: tt.fields.config,
				client: tt.fields.client,
			}
			if err := o.Copy(tt.args.srcKey, tt.args.destKey); (err != nil) != tt.wantErr {
				t.Errorf("Copy() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestMinioOption_Delete(t *testing.T) {
	type fields struct {
		config *MinioConfig
		client *minio.Client
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
			o := &MinioOption{
				config: tt.fields.config,
				client: tt.fields.client,
			}
			if err := o.Delete(tt.args.key); (err != nil) != tt.wantErr {
				t.Errorf("Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestMinioOption_Exists(t *testing.T) {
	type fields struct {
		config *MinioConfig
		client *minio.Client
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
			o := &MinioOption{
				config: tt.fields.config,
				client: tt.fields.client,
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

func TestMinioOption_Get(t *testing.T) {
	type fields struct {
		config *MinioConfig
		client *minio.Client
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
			o := &MinioOption{
				config: tt.fields.config,
				client: tt.fields.client,
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

func TestMinioOption_GetDiskName(t *testing.T) {
	type fields struct {
		config *MinioConfig
		client *minio.Client
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
			o := &MinioOption{
				config: tt.fields.config,
				client: tt.fields.client,
			}
			if got := o.GetDiskName(); got != tt.want {
				t.Errorf("GetDiskName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMinioOption_Put(t *testing.T) {
	//data, _ := ioutil.ReadFile("../tests/accounts.txt")
	//err := minioOss.Put("test_data/accounts.txt", bytes.NewReader(data), int64(len(data)), "text/plain")
	//if err != nil {
	//	t.Error(err.Error())
	//	return
	//}

	data, _ := ioutil.ReadFile("../tests/33.jpg")
	err := minioOss.Put("test_data/777.jpg", bytes.NewReader(data), int64(len(data)), "image/jpeg")
	if err != nil {
		t.Error(err.Error())
		return
	}
	t.Log("success")
}

func TestMinioOption_PutBase64(t *testing.T) {
	type fields struct {
		config *MinioConfig
		client *minio.Client
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
			o := &MinioOption{
				config: tt.fields.config,
				client: tt.fields.client,
			}
			if err := o.PutBase64(tt.args.key, tt.args.r, tt.args.contentType); (err != nil) != tt.wantErr {
				t.Errorf("PutBase64() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestMinioOption_PutFile(t *testing.T) {
	type fields struct {
		config *MinioConfig
		client *minio.Client
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
			o := &MinioOption{
				config: tt.fields.config,
				client: tt.fields.client,
			}
			if err := o.PutFile(tt.args.key, tt.args.localFile, tt.args.contentType); (err != nil) != tt.wantErr {
				t.Errorf("PutFile() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestMinioOption_Rename(t *testing.T) {
	type fields struct {
		config *MinioConfig
		client *minio.Client
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
			o := &MinioOption{
				config: tt.fields.config,
				client: tt.fields.client,
			}
			if err := o.Rename(tt.args.srcKey, tt.args.destKey); (err != nil) != tt.wantErr {
				t.Errorf("Rename() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestMinioOption_Size(t *testing.T) {
	type fields struct {
		config *MinioConfig
		client *minio.Client
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
			o := &MinioOption{
				config: tt.fields.config,
				client: tt.fields.client,
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

func TestMinioOption_Url(t *testing.T) {
	url := minioOss.Url("test_data/33.jpg")
	t.Log("url : " + url)
	t.Log("success")
}

func TestNewMinioOption(t *testing.T) {
	type args struct {
		cfg *MinioConfig
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
			got, err := NewMinioOption(tt.args.cfg)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewMinioOption() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewMinioOption() got = %v, want %v", got, tt.want)
			}
		})
	}
}
