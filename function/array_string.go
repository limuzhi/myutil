/*
 * @PackageName: utils
 * @Description:
 * @Author: limuzhi
 * @Date: 2022/12/1 14:16
 */

package function

import (
	"sort"
)

//数组是否包含某个字符串
func ContainsString(arr []string, s string) bool {
	for _, v := range arr {
		if v == s {
			return true
		}
	}
	return false
}

//数组删除某个字符串
func RemoveString(arr []string, s string) []string {
	newArr := make([]string, 0)
	for _, v := range arr {
		if v != s {
			newArr = append(newArr, v)
		}
	}
	return newArr
}

func DuplicateString(l []string) []string {
	m := make(map[string]string, 0)
	arr := make([]string, 0)
	for _, v := range l {
		if _, ok := m[v]; !ok {
			m[v] = v
			arr = append(arr, v)
		}
	}
	sort.Strings(arr)
	return arr
}
