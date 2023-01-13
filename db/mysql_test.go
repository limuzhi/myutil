/*
 * @PackageName: db
 * @Description:
 * @Author: limuzhi
 * @Date: 2022/12/1 11:11
 */

package db

import (
	"gorm.io/gorm"
	"testing"
	"time"
)

func TestNewMysql(t *testing.T) {
	type args struct {
		opts *MysqlOptions
	}
	tests := []struct {
		name    string
		args    args
		want    *gorm.DB
		wantErr bool
	}{
		{name: "ccc", args: args{opts: &MysqlOptions{
			Source:                "root:123456@tcp(127.0.0.1:3306)/affiliate_go?timeout=1s&readTimeout=1s&writeTimeout=1s&parseTime=true&loc=Local&charset=utf8mb4",
			MaxIdleConnections:    100,
			MaxOpenConnections:    100,
			MaxConnectionLifeTime: time.Hour,
		}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := NewMysql(tt.args.opts)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewMysql() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			//TODO
			//if !reflect.DeepEqual(got, tt.want) {
			//	t.Errorf("NewMysql() got = %v, want %v", got, tt.want)
			//}
		})
	}
}
