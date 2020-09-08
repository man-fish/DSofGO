package kmp

func Make(s, p string) (res []int) {
	nxt := make([]int, len(p))
	buildNext(p, nxt)

	var (
		start = 0
		i     = 0
		res   = make([]int, 0)
	)

	for start <= len(s)-len(p) {
		curIdx = start + i

		if i == len(p) {
			res = append(res, start)
			start = curIdx
			i = 0
			continue
		}
		if s[curIdx] == p[i] {
			i++
		} else if i == 0 {
			start++
		} else {
			start = curIdx - nxt[i-1]
			i = nxt[i-1]
		}
	}
	return res

}

func buildNext(p string, nxt []int) {
	var (
		i int = 1
		n int = 0
	)
	nxt[0] = 0
	for i < len(p) {
		if p[i] == p[n] {
			n++
			nxt[i] = n
			i++
		} else if n != 0 {
			n = nxt[n-1]
		} else {
			nxt[i] = n
			i++
		}
	}
}
