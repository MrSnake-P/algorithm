package string

/*
例1：
输入：["h","e","l","l","o"]
输出：["o","l","l","e","h"]
例2：
输入：["H","a","n","n","a","h"]
输出：["h","a","n","n","a","H"]
*/
/*
使用双指针：
1.分别指向头和尾
2.然后进行交换直至left=right
*/

func reverseString(s []byte) {
	left := 0
	right := len(s) - 1
	for left < right {
		s[left], s[right] = s[right], s[left]
		left++
		right--
	}
}
