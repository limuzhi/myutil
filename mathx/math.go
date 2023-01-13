/*
 * @PackageName: mathx
 * @Description:
 * @Author: limuzhi
 * @Date: 2022/12/1 14:37
 */

package mathx

import (
	"fmt"
	"github.com/shopspring/decimal"
	"math"
	"strconv"
)

// Round 四舍五入
func Round(val float64, places int) float64 {
	var t float64
	f := math.Pow10(places)
	x := val * f
	if math.IsInf(x, 0) || math.IsNaN(x) {
		return val
	}
	if x >= 0.0 {
		t = math.Ceil(x)
		if (t - x) > 0.50000000001 {
			t -= 1.0
		}
	} else {
		t = math.Ceil(-x)
		if (t + x) > 0.50000000001 {
			t -= 1.0
		}
		t = -t
	}
	x = t / f

	if !math.IsInf(x, 0) {
		return x
	}

	return t
}

func Keep2Decimal(value float64) float64 {
	value, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", value), 64)
	return value
}

func Keep8Decimal(value float64) float64 {
	value, _ = decimal.NewFromFloat(value).Round(8).Float64()
	return value
}
