package euler

func LargestPrimeFactor(number int) int {
	allPrime := FindAllPrime(number)
	n := len(allPrime)
	largestPrimeFactor := number
	for i := n - 1; i >= 0; i-- {
		if number % allPrime[i] == 0{
			largestPrimeFactor = allPrime[i]
			break
		}
	}
	return largestPrimeFactor
}

func isPrime(number int) bool {
	if number == 0 || number == 1 {
		return false
	}
	for i := 2; i <= number / i ; i++ {
		if number % i == 0 {
			return false
		}
	}
	return true
}

func FindAllPrime(number int) []int {
	allPrime := []int{}
	for i := 0; i < number; i++ {
		if isPrime(i) {
			allPrime = append(allPrime, i)
		}
	}
	return allPrime
}
