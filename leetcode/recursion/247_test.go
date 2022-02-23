package recursion

//中心对称数  找到所有长度为n的中心对称数
/*
	n=1 0 1 8
	n=2 11 69 88 96
*/

func helper(n, m int) []string {
	if n < 0 || m < 0 || n < m {
		panic("input error")
	}

	rs := []string{}

	if n == 0 {
		return rs
	}
	if n == 1 {
		rs := append(rs, "0", "1", "8")
		return rs
	}

	n2 := helper(n-2, m)
	for i := 0; i < len(n2); i++ {
		a := n2[i]
		if n != m {
			rs = append(rs, "0"+a+"0")
		}
		rs = append(rs, "1"+a+"1")
		rs = append(rs, "8"+a+"8")
		rs = append(rs, "6"+a+"9")
		rs = append(rs, "9"+a+"6")
	}
	return rs
}
