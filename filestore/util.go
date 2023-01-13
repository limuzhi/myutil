/*
 * @PackageName: filestore
 * @Description:
 * @Author: limuzhi
 * @Date: 2022/12/1 15:49
 */

package filestore

import "strings"

func NormalizeKey(key string) string {
	key = strings.Replace(key, "\\", "/", -1)
	key = strings.Replace(key, " ", "", -1)
	key = filterNewLines(key)

	return key
}

func filterNewLines(s string) string {
	return strings.Map(func(r rune) rune {
		switch r {
		case 0x000A, 0x000B, 0x000C, 0x000D, 0x0085, 0x2028, 0x2029:
			return -1
		default:
			return r
		}
	}, s)
}
