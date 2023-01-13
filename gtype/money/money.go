package money

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"
)

// Money 金钱类型
type Money float64

// func NewMoney(a float64) Money { return &Money(a) }
// 1.00499
func (m Money) String() string {
	// strValBytes := []byte(fmt.Sprintf("%.3f", m.Val()))
	// pointPos := 0
	// for i, ch := range strValBytes {
	// 	if ch == '.' {
	// 		pointPos = i
	// 	}
	// }
	// lastPos := pointPos + 3
	// if strValBytes[lastPos] == '5' {
	// 	strValBytes[lastPos] = strValBytes[lastPos] + 1
	// }
	// value, _ := strconv.ParseFloat(string(strValBytes), 64)
	value := m.Val()
	return fmt.Sprintf("%.2f", value)
}

// Add .
func (m Money) Add(a Money) Money {
	return m + a
}

// Sub .
func (m Money) Sub(a Money) Money {
	return m - a
}

// Mul .
func (m Money) Mul(a float64) Money {
	return m * Money(a)
}

// Div .
func (m Money) Div(a float64) Money {
	return m / Money(a)
}

// Val .
func (m Money) Val() float64 {
	return round(float64(m), 2)
}

// Less <
func (m Money) Less(a Money) bool {
	return m.Val() < a.Val()
}

// IsZero 是否为0
func (m Money) IsZero() bool {
	return m.Val() == 0.0
}

// Abs 求绝对值
func (m Money) Abs() Money {
	return Money(math.Abs(float64(m)))
}

// MarshalJSON .
func (m Money) MarshalJSON() ([]byte, error) {
	return []byte(m.String()), nil
}

// ToUpper 人民币大写
func (m Money) ToUpper() string {
	if m.Val() == 0 {
		return "零圆整"
	}
	unit := []string{"仟", "佰", "拾", "亿", "仟", "佰", "拾", "万", "仟", "佰", "拾", "圆", "角", "分"}
	upper := map[string]string{"0": "零", "1": "壹", "2": "贰", "3": "叁", "4": "肆", "5": "伍", "6": "陆", "7": "柒", "8": "捌", "9": "玖"}
	unitPrice := strconv.FormatFloat(m.Val()*100, 'f', 0, 64)
	s := unit[len(unit)-len(unitPrice) : len(unit)]
	str := ""
	for k, v := range unitPrice[:] {
		str = str + upper[string(v)] + s[k]
	}
	reg, _ := regexp.Compile(`零角零分$`)
	str = reg.ReplaceAllString(str, "整")

	reg, _ = regexp.Compile(`零角`)
	str = reg.ReplaceAllString(str, "零")

	reg, _ = regexp.Compile(`零分$`)
	str = reg.ReplaceAllString(str, "整")

	reg, _ = regexp.Compile(`零[仟佰拾]`)
	str = reg.ReplaceAllString(str, "零")

	reg, _ = regexp.Compile(`零{2,}`)
	str = reg.ReplaceAllString(str, "零")

	reg, _ = regexp.Compile(`零亿`)
	str = reg.ReplaceAllString(str, "亿")

	reg, _ = regexp.Compile(`零万`)
	str = reg.ReplaceAllString(str, "万")

	reg, _ = regexp.Compile(`零*圆`)
	str = reg.ReplaceAllString(str, "圆")

	reg, _ = regexp.Compile(`亿零{0, 3}万`)
	str = reg.ReplaceAllString(str, "^圆")

	reg, _ = regexp.Compile(`零圆`)
	str = reg.ReplaceAllString(str, "零")

	if strings.Contains(str, "角") || strings.Contains(str, "分") {
		str = strings.Replace(str, "整", "", -1)
	}
	return str

}

// round 四舍五入, "../mathx"
func round(val float64, places int) float64 {
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

func (m Money) Decimal() Money {
	val := float64(m)
	val = math.Trunc(val*1e2) * 1e-2
	val, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", val), 64)
	return Money(val)
}

func (m Money) Rounding() Money {
	val := float64(m)
	val = math.Trunc(val*1e2+0.5) * 1e-2
	return Money(val)
}
