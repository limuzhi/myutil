/*
 * @PackageName: utils
 * @Description: 类型转换
 * @Author: limuzhi
 * @Date: 2022/12/1 13:52
 */

package function

import (
	"reflect"
	"strconv"
)

func Int64ToInt(int64Value int64) int {
	strValue := strconv.FormatInt(int64Value, 10)
	intValue, _ := strconv.Atoi(strValue)
	return intValue
}

func IntToInt64(intValue int) int64 {
	strValue := strconv.Itoa(intValue)
	int64Value, _ := strconv.ParseInt(strValue, 10, 64)

	return int64Value
}

func IntToStr(intValue int) string {
	return strconv.Itoa(intValue)
}

func Int64ToStr(intValue int64) string {
	return strconv.FormatInt(intValue, 10)
}

func StrToInt(stringValue string) (int, error) {
	intValue, err := strconv.Atoi(stringValue)
	if err != nil {
		return 0, err
	}
	return intValue, nil
}

func StrToInt64(stringValue string) (int64, error) {
	int64Value, err := strconv.ParseInt(stringValue, 10, 64)
	if err != nil {
		return 0, err
	}
	return int64Value, nil
}

func Float64ToStr(float64Value float64) string {
	return strconv.FormatFloat(float64Value, 'f', -1, 64)
}

func StrToFloat64(stringValue string) (float64, error) {
	float64Value, err := strconv.ParseFloat(stringValue, 64)
	if err != nil {
		return 0, err
	}
	return float64Value, nil
}

func StrToBool(stringValue string) (value bool, err error) {
	value, err = strconv.ParseBool(stringValue)
	return
}

func Float64ToInt(float64Value float64) int {
	return int(float64Value)
}

/**
任意类型数组转为interface{}类型数组
*/
func ToInterfaceArr(arr interface{}) []interface{} {
	if reflect.TypeOf(arr).Kind() != reflect.Slice {
		return nil
	}

	arrValue := reflect.ValueOf(arr)
	retArr := make([]interface{}, arrValue.Len())
	for k := 0; k < arrValue.Len(); k++ {
		retArr[k] = arrValue.Index(k).Interface()
	}
	return retArr
}
