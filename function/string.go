/*
 * @PackageName: function
 * @Description:
 * @Author: limuzhi
 * @Date: 2022/12/2 14:37
 */

package function

import "strings"

//字符串拼接
func BuilderConcat(n int, str string) string {
	var builder strings.Builder
	builder.Grow(n * len(str))
	for i := 0; i < n; i++ {
		builder.WriteString(str)
	}
	return builder.String()
}
