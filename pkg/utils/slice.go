package utils

func Filter(source []int64, filter func(int64) bool) []int64 {
	result := make([]int64, 0, len(source))
	for _, v := range source {
		a := v
		if filter(a) {
			result = append(result, a)
		}
	}
	return result
}

// StringSliceDiff s1-s2
func StringSliceDiff(s1, s2 []string) []string {
	if len(s1) == 0 {
		return nil
	}
	if len(s2) == 0 {
		return s1
	}

	tmpMap := make(map[string]int8, len(s1)+len(s2))
	diff := make([]string, 0, len(s1))
	for _, v := range s1 {
		tmpMap[v] |= 1
	}
	for _, v := range s2 {
		tmpMap[v] |= 2
	}
	for k, v := range tmpMap {
		if v == 1 {
			diff = append(diff, k)
		}
	}
	return diff
}

//多个切片求交集
func Intersect(lists [][]string) []string {
	var inter []string
	mp := make(map[string]int)
	l := len(lists)

	// 特判 只传了0个或者1个切片的情况
	if l == 0 {
		return make([]string, 0)
	}
	if l == 1 {
		for _, s := range lists[0] {
			if _, ok := mp[s]; !ok {
				mp[s] = 1
				inter = append(inter, s)
			}
		}
		return inter
	}

	// 一般情况
	// 先使用第一个切片构建map的键值对
	for _, s := range lists[0] {
		if _, ok := mp[s]; !ok {
			mp[s] = 1
		}
	}

	// 除去第一个和最后一个之外的list
	for _, list := range lists[1 : l-1] {
		for _, s := range list {
			if _, ok := mp[s]; ok {
				// 计数+1
				mp[s]++
			}
		}
	}

	for _, s := range lists[l-1] {
		if _, ok := mp[s]; ok {
			if mp[s] == l-1 {
				inter = append(inter, s)
			}
		}
	}
	return inter
}
