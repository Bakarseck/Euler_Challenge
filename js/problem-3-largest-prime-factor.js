function isPrime(number) {
  if (number==0 || number==1) {
    return false
  }
  for (let i = 2; i <= number / i ; i++) {
		if (number % i == 0) {
			return false
		}
	}
  return true
}

function FindAllPrime(number) {

	let allPrime = []

	for (let i = 0; i < number; i++) {
		if (isPrime(i)){
			allPrime.push(i)
		}
	}

	return allPrime
}

function largestPrimeFactor(number) {
	let allPrime = FindAllPrime(number)

	let n = allPrime.length

	let largestPrimeFactor = number

	for (let i = n - 1; i >= 0; i--) {
		if (number % allPrime[i] == 0) {
			largestPrimeFactor = allPrime[i]
			break
		}
	}

  console.log(largestPrimeFactor)

	return largestPrimeFactor
}

largestPrimeFactor(2);
