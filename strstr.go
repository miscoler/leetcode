func strStr(haystack string, needle string) (res int) {
	lenNeedle := len(needle)
	if lenNeedle == 0 {
		return 0
	}
	res = -1
	for i := 0; i < len(haystack) - lenNeedle + 1; i ++ {
		j:= 0
		for ; j < lenNeedle && haystack[i + j] == needle[j]; j++ {
		}
		if j > lenNeedle - 1 {
			return i
		}
	}
	return
}
