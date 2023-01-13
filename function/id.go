/*
 * @PackageName: function
 * @Description:
 * @Author: limuzhi
 * @Date: 2022/12/3 9:44
 */

package function

import (
	"github.com/bwmarrin/snowflake"
	"github.com/google/uuid"
	"github.com/sony/sonyflake"
)

//	"github.com/gofrs/uuid"
//生成uuid

func GenerateUUID() string {
	return uuid.New().String()
}

func GetUUID() uuid.UUID {
	return uuid.New()
}

func GetSnowflakeId() (int64, error) {
	n, err := snowflake.NewNode(1)
	if err != nil {
		return 0, err
	}
	return n.Generate().Int64(), nil
}

func GetSonyflakeNextId() (uint64, error) {
	flake := sonyflake.NewSonyflake(sonyflake.Settings{})
	id, err := flake.NextID()
	return id, err
}
