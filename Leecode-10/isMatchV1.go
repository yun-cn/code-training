package main
func isMatch(s string, p string) bool {
	slen, plen := len(s), len(p)
	if slen == 0 {
		if plen%2 != 0 {
			return false
		}
		for j := 1; j < plen; j += 2 {
			if p[j] != '*' {
				return false
			}
		}
		return true
	}
	if plen == 0 {
		return slen == 0
	}
	dp := make([][]bool, slen, slen)
	for i := 0; i < slen; i++ {
		dp[i] = make([]bool, plen, plen)
	}
	// init: dp(0, j)
	dp[0][0] = p[0] == s[0] || p[0] == '.'
	for canPreEmpty, j := true, 1; j < plen; j += 1 {
		if canPreEmpty && j%2 == 1 && p[j] != '*' {
			canPreEmpty = false
		}
		switch p[j] {
		case '.':
			dp[0][j] = canPreEmpty
		case '*':
			if p[j-1] == s[0] || p[j-1] == '.' {
				dp[0][j] = canPreEmpty || (j >= 2 && dp[0][j-2])
			} else {
				dp[0][j] = j-2 >= 0 && dp[0][j-2]
			}
		default:
			dp[0][j] = s[0] == p[j] && canPreEmpty
		}
	}
	// cal: dp(i, j)
	for i := 1; i < slen; i++ {
		for j := 1; j < plen; j++ {
			switch p[j] {
			case '.':
				dp[i][j] = dp[i-1][j-1]
			case '*':
				if p[j-1] == s[i] || p[j-1] == '.' {
					dp[i][j] = (j-2 >= 0 && dp[i][j-2]) || dp[i-1][j]
				} else {
					dp[i][j] = j-2 >= 0 && dp[i][j-2]
				}
			default:
				dp[i][j] = s[i] == p[j] && dp[i-1][j-1]
			}
		}
	}
	return dp[slen-1][plen-1]
}

