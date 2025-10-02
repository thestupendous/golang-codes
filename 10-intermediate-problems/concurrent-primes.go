/* 
   using Sieve of Eratosthenes
*/ 
package main

import (
	"fmt"
)

func findPrimes(n int, notPrime *[101]bool) {
	var totalIterations int = 0
	for i:=2; i*i<=n; i++ {
		totalIterations += 1
		if !notPrime[i] {
			
			totalIterations -= 1
			for j:=i*i ; j<=n; j += i {
				totalIterations += 1
				notPrime[j] = true
			}
		}
	}

	fmt.Println("found out till ",n,", printing: ")
	for i:=1;i<=n;i++  {
		if !notPrime[i] {
			fmt.Printf("%v, ",i)
		}
	}
	fmt.Println()

	fmt.Println("Checking next (error prone) : for ",n+1,": ",!notPrime[n+1])

	fmt.Println("total iterations : ",totalIterations)
}

func main () {
	var n int
	fmt.Scanln(&n)

	var notPrime [101]bool

	findPrimes(n, &notPrime)


	_ = n
}
