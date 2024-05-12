package main

/*
 * @lc app=leetcode.cn id=151 lang=golang
 *
 * [151] 反转字符串中的单词
 */

// @lc code=start
// func reverseWords(s string) string {
// 	list := strings.Split(s, " ")

// 	resBuilder := strings.Builder{}

// 	for i := 0; i < len(list); i++ {
// 		v := list[len(list)-1-i]
// 		if v == " " || v == "" {
// 			continue
// 		} else {
// 			resBuilder.WriteString(strings.TrimSpace(v))
// 			resBuilder.WriteString(" ")
// 		}
// 	}
// 	return strings.TrimSpace(resBuilder.String())
// }

// @lc code=end
