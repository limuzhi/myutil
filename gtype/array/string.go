/*
 * @PackageName: utils
 * @Description:
 * @Author: limuzhi
 * @Date: 2022/12/1 14:00
 */

package array

import (
	"encoding/json"
	"fmt"
)

// Array 字符串数组
type Array string

func New(v ...string) Array {
	if len(v) == 0 {
		return Array("")
	}
	bs, _ := json.Marshal(v)
	return Array(bs)
}

// MarshalJSON .
func (i Array) MarshalJSON() ([]byte, error) {
	imgs := i.Slice()
	return json.Marshal(imgs)
}

// UnmarshalJSON .
func (i *Array) UnmarshalJSON(data []byte) error {
	it := Array(string(data))
	*i = it
	return nil
}

// Slice .
func (i Array) Slice() (imgs []string) {
	imgs = []string{}
	if i == "" {
		return
	}

	err := decodeJSON(string(i), &imgs)
	if err != nil {
		fmt.Printf("array %s decode failed, %s", i, err)
		return
	}
	return
}

func decodeJSON(bs string, v interface{}) error {
	return json.Unmarshal([]byte(bs), v)
}

// Contains 字符串是否包含在数组中
func Contains(arr []string, str string) bool {
	for _, v := range arr {
		if v == str {
			return true
		}
	}
	return false
}
