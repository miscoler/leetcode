func plusOne(digits []int) (res []int) {
    var overflow bool = true
    n := len(digits)
    for i := n - 1; i >= 0 && overflow; i-- {
        digits[i] = (digits[i] + 1) % 10
        if digits[i] > 0 {
            overflow = false
        }
    }
    if overflow {
        digits = append([]int{1}, digits...)
    }
    return digits
}
