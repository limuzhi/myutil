/*
 * @PackageName: filestore
 * @Description:
 * @Author: limuzhi
 * @Date: 2022/12/1 16:02
 */

package filestore

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"testing"
)

var (
	disk, aliOss, minioOss Storage
)

func TestMain(t *testing.M) {
	qiniYun, _ := NewQiniuYunOption(&QiniuyunConfig{
		AccessKey: "xxx",
		Bucket:    "xx",
		Domain:    "xx.recomsale.com",
		SecretKey: "vxx",
		IsSsl:     true,
		IsPrivate: false,
	})
	aliyun, _ := NewAliYunOssOption(&AliyunConfig{
		AccessKeyId:     "",
		AccessKeySecret: "",
		Bucket:          "",
		Endpoint:        "",
		IsSsl:           true,
		IsPrivate:       false,
	})
	minios, err := NewMinioOption(&MinioConfig{
		AccessKeyId:     "NVmy1sHJEq0tNAOb",
		AccessKeySecret: "7F0rPDrTYJ2byCzWpbRDxzab4mCHRFvM",
		Bucket:          "demo",
		Endpoint:        "127.0.0.1:9000",
		IsSsl:           false,
		IsPrivate:       false,
	})
	if err != nil {
		fmt.Println("err:", err.Error())
	}
	oss := NewWithOptions(WithDisabled(false), WithStorage(qiniYun), WithStorage(aliyun), WithStorage(minios))
	oss.Register(KoDo, qiniYun)
	disk, _ = oss.GetStorage(KoDo)
	aliOss, _ = oss.GetStorage(Oss)
	minioOss, _ = oss.GetStorage(Minio)
	t.Run()
}

func TestQiniuYun_Copy(t *testing.T) {
	if err := disk.Copy("test_data/accounts.txt", "test_data/accounts2.txt"); err != nil {
		t.Errorf("Copy() error = %v", err)
	}
}

func TestQiniuYun_Delete(t *testing.T) {
	if err := disk.Delete("test_data/get_put_accounts.txt"); err != nil {
		t.Errorf("Delete() error = %v", err)
	}
}

func TestQiniuYun_Exists(t *testing.T) {
	exists, err := disk.Exists("test_data/accounts.txt")
	if err != nil {
		t.Error(err.Error())
		return
	}
	t.Logf("isExisted : %v", exists)
	t.Log("success")
}

func TestQiniuYun_Get(t *testing.T) {
	body, err := disk.Get("test_data/accounts2.txt")
	if _, ok := body.(io.Closer); ok {
		defer body.Close()
	}
	if err != nil {
		t.Error(err.Error())
		return
	}

	data, err := ioutil.ReadAll(body)
	t.Log(string(data))
	err = disk.Put("test_data/get_put_accounts.txt", bytes.NewReader(data), int64(len(data)), "text/plain")
	if err != nil {
		t.Error(err.Error())
		return
	}
	t.Log("success")
}

func TestQiniuYun_Put(t *testing.T) {
	data, _ := ioutil.ReadFile("../tests/accounts.txt")
	err := disk.Put("test_data/accounts.txt", bytes.NewReader(data), int64(len(data)), "text/plain")
	if err != nil {
		t.Error(err.Error())
		return
	}
	t.Log("success")
}

func TestQiniuYun_PutFile(t *testing.T) {
	err := disk.PutFile("test_data/accounts2.txt", "../tests/accounts.txt", "text/plain")
	if err != nil {
		t.Error(err.Error())
		return
	}
	t.Log("success")
}

func TestQiniuYun_Rename(t *testing.T) {
	err := disk.Rename("test_data/accounts.txt", "test_data/rename_accounts.txt")
	if err != nil {
		t.Error(err.Error())
		return
	}
	t.Log("success")
}

func TestQiniuYun_Size(t *testing.T) {
	size, err := disk.Size("test_data/accounts.txt")
	if err != nil {
		t.Log(err.Error())
		return
	}
	t.Logf("size : %d", size)
	t.Log("success")
}

func TestQiniuYun_Url(t *testing.T) {
	url := disk.Url("test_data/accounts.txt")
	t.Log("url : " + url)
	t.Log("success")
}
