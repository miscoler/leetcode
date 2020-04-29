
const max32 = 2147483647 // 2^31-1
const min32 = -2147483648 // -2^31

func reverse(x int) int {
	if x > max32 || x < min32 {
		return 0
	}
	var sign bool
	if x < 0 {
		sign = true
		x = -x
	}
	res := 0

	for ; x > 0; x = x / 10 {
		digit := x % 10
		res = res * 10 + digit
	}
	if sign {
		res = -res
	}
	return res
}
