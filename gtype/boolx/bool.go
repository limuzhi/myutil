/*
 * @PackageName: utils
 * @Description:
 * @Author: limuzhi
 * @Date: 2022/12/1 13:55
 */

package boolx

type Bool int32

const (
	// False .
	False Bool = 0
	// True .
	True Bool = 1
	// Any .
	Any Bool = 2
)

// BoolName 对应关系
var BoolName = map[Bool]string{
	False: "FALSE",
	True:  "TRUE",
	Any:   "ANY",
}

// String .
func (x Bool) String() string {
	if v, ok := BoolName[x]; ok {
		return v
	}
	return BoolName[False]
}
