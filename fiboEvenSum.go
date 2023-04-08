package euler

func FibonacciSum(n int) int {
	
    var sum, prev, current int
    prev, current = 0, 1

    for current <= n {
        if current%2 == 0 {
            sum += current
        }
        prev, current = current, prev+current
    }

    return sum
}