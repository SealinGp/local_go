package main

var keyChar = []byte("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

//根据所给整型返回一个字符
//n = 1
//n = 5
func genKey(n int) string  {
	if n == 0 {
		return string(keyChar[0])
	}
	//l = 62
	l := len(keyChar)

	// i = 20
	s := make([]byte,20)
	i := len(s)

	//n = 1,i = 20
	//n = 5,i = 20
	for n > 0 && i >= 0 {

		//i=19
		i--

		//j = 1 % 62 = 1
		//j = 5 / 62 = 5
		j := n % l

		//n = 1-1 / 62 = 0
		//n = 5 - 5 / 62 = 0
		n = (n - j) / l

		//s[19] = keyChar[1] = 1
		//s[19] = keyChar[5] = 4
		s[i] = keyChar[j]
	}
	return string(s[i:])
}