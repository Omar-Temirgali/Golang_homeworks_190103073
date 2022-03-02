// Homework 0: Hello Go!
package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	// Feel free to use the main function for testing your functions
	fmt.Println("Hello, à¤¦à¥à¤¨à¤¿à¤¯à¤¾!")
	fmt.Println("Сәлем, Әлем!")
	fmt.Println("こんにちは世界！")
	fmt.Println("Hello, World!")

	// Fizzbuzz
	fmt.Println(Fizzbuzz(9), Fizzbuzz(25), Fizzbuzz(75), Fizzbuzz(31))

	//IsPrime
	fmt.Println("15 is prime:", IsPrime(15))
	fmt.Println("0 is prime:", IsPrime(0))
	fmt.Println("7 is prime:", IsPrime(7))
	fmt.Println("941 is prime:", IsPrime(941))
	fmt.Println("-7 is prime:", IsPrime(-7))

	//IsPalindrome
	fmt.Println("rooT is palindrome:", IsPalindrome("rooT"))
	fmt.Println("noon is plaindrome:", IsPalindrome("noon"))
	fmt.Println("wwWwW is palindrome:", IsPalindrome("wwWwW"))
	fmt.Println("1456541 is palindrome:", IsPalindrome("1456541"))
	fmt.Println("Ukraine is palindrome:", IsPalindrome("Ukraine"))
}

// Fizzbuzz is a classic introductory programming problem.
// If n is divisible by 3, return "Fizz"
// If n is divisible by 5, return "Buzz"
// If n is divisible by 3 and 5, return "FizzBuzz"
// Otherwise, return the empty string
func Fizzbuzz(n int) string {
	if n%3 == 0 && n%5 == 0 {
		return "FizzBuzz"
	} else if n%3 == 0 {
		return "Fizz"
	} else if n%5 == 0 {
		return "Buzz"
	} else {
		return ""
	}
}

// IsPrime checks if the number is prime.
// You may use any prime algorithm, but you may NOT use the standard library.
func IsPrime(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

// IsPalindrome checks if the string is a palindrome.
// A palindrome is a string that reads the same backward as forward.
func IsPalindrome(s string) bool {
	s_lower := strings.ToLower(s)

	for i := 0; i < len(s_lower)/2; i++ {
		if s_lower[i] != s_lower[len(s_lower)-i-1] {
			return false
		}
	}
	return true
}
