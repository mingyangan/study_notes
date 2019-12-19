package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	buf := []byte(s)
	var maxLen = 0
	var bufLen = len(buf)
	var nextStart = 0
	var lastStart = 0
	m := make(map[byte]int)
	for i := 0; i < bufLen; i = nextStart {
		for x := lastStart; x < nextStart && x < bufLen; x++ {
			delete(m, buf[x])
		}
		lastStart = nextStart

		for j := i + len(m); j < bufLen; j++ {
			v, ok := m[buf[j]]
			if !ok {
				m[buf[j]] = j
				continue
			}
			nextStart = v + 1
			break
		}
		if maxLen < len(m) {
			maxLen = len(m)
		}
		if maxLen >= bufLen-i {
			break
		}
	}
	return maxLen
}

func main() {
	fmt.Println(lengthOfLongestSubstring("jxdlnaaij"))
}
