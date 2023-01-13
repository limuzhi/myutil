package function

// 切片去重
func RemoveRepeated(arr []string) (newArr []string) {
	newArr = make([]string, 0)
	for i := 0; i < len(arr); i++ {
		repeat := false
		for j := i + 1; j < len(arr); j++ {
			if arr[i] == arr[j] {
				repeat = true
				break
			}
		}
		if !repeat {
			newArr = append(newArr, arr[i])
		}
	}
	return
}

// 取交集的函数
func MixedSlice(s1 []string, s2 []string) (mixedS []string) {
	for i := 0; i < len(s1); i++ {
		for j := 0; j < len(s2); j++ {
			if s1[i] == s2[j] {
				mixedS = append(mixedS, s1[i])
			}
		}
	}
	mixedS = RemoveRepeated(mixedS)
	return
}

// 求差集：父集与子集的差集
func Difference(father []string, son []string) (diffedS []string) {
	// 去重
	newFather := RemoveRepeated(father)
	newSon := RemoveRepeated(son)
	lenFather := len(newFather)
	lenSon := len(newSon)
	if lenFather < lenFather {
		panic("not son of father...")
	} else if lenFather == lenSon {
		return
	} else {
		for i := 0; i < lenFather; i++ {
			if !isInArray(newSon, newFather[i]) {
				diffedS = append(diffedS, newFather[i])
			}
		}
	}
	return
}
func isInArray(s []string, item string) (b bool) {
	b = false
	for i := 0; i < len(s); i++ {
		if s[i] == item {
			b = true
			break
		}
	}
	return
}
