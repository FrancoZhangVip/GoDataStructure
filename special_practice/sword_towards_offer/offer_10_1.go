package sword_towards_offer

func Fib(n int) int {
	if n <= 1 {
		return n
	}
	recorder := make([]int, 0)
	recorder = append(recorder, 0, 1)
	for i := 2; i <= n; i++ {
		item := (recorder[i-1] + recorder[i-2]) % 1000000007
		recorder = append(recorder, item)
	}
	return recorder[n]
}
