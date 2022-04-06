package simple

func (*Ref) ClimbStairs() {

}

func climbStairs(n int) int {

	p, q, r := 0, 0, 1

	for i := 0; i < n; i++ {
		p = q
		q = r
		r = p + q
	}

	return r
}
