package leetLongestPalindromicSubstring

// f(n) = max(f(n-1), g(n))
// f(n) - maximum palindrome length on n-step
// g(n) - longest palindrome length on n-step (n is a center of the palindrome)
func longestPalindrome(s string) string {
	n := len(s)
	dp := make([]int, n)
	dp[0] = 1
	pal := s[0:1]
	res := pal
	for i := 1; i < n; i++ {
		pal = centeredPalindrome(s, i)
		dp[i] = max(dp[i-1], len(pal))
		if len(pal) > len(res) {
			res = pal
		}
	}
	return res
}

func centeredPalindrome(s string, i int) string {
	l := i - 1
	r := i + 1
	for l >= 0 && r < len(s) && s[l] == s[r] {
		l--
		r++
	}
	odd := s[l+1 : r]
	l = i - 1
	r = i
	for l >= 0 && r < len(s) && s[l] == s[r] {
		l--
		r++
	}
	even := s[l+1 : r]
	if len(even) > len(odd) {
		return even
	}
	return odd
}
