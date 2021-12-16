package dptr

//给你一个字符串 s，找到 s 中最长的回文子串。
//逐个找，如果 s[i]==s[i+1]或者s[i-1] == s[i+1]
//就开始使用双指针，找出回文的长度

func LongestPalindrome(s string) string {
	if len(s) < 2 {
		return s
	}

	lPtr := 0
	rPtr := 0
	var ret string
	ret = string(s[0])

	for i := 0; i < len(s)-1; i++ {
		if s[i] == s[i+1] {
			lPtr = i
			rPtr = i + 1
			str := getIndexes(s, lPtr, rPtr)
			if len(str) > len(ret) {
				ret = str
			}
		}
		if i > 0 && s[i-1] == s[i+1] {
			lPtr = i - 1
			rPtr = i + 1
			str := getIndexes(s, lPtr, rPtr)
			if len(str) > len(ret) {
				ret = str
			}
		} else {
			continue
		}
	}

	return ret
}

func getIndexes(s string, lPtr, rPtr int) string {
	for lPtr > 0 && rPtr < len(s)-1 {
		lPtr--
		rPtr++

		if s[lPtr] == s[rPtr] {
			continue
		} else {
			lPtr++
			rPtr--
			break
		}
	}
	return s[lPtr : rPtr+1]
}
