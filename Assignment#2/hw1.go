// Homework 1: Finger Exercises
// Due January 31, 2017 at 11:59pm
package main

import (
	"fmt"
	"regexp"
	"sort"
	"strings"
)

func main() {
	fmt.Println(ParsePhone("123-456-7890"))
	fmt.Println(ParsePhone("12 7  456-  7  8  9 0"))
	fmt.Println(ParsePhone("12 456wd w-78 90"))
	fmt.Println(ParsePhone("1 2 3 4 5 6 7 8 9 0"))
	fmt.Println(ParsePhone("78 698 07"))

	fmt.Println(ParsePhoneWithRegex("123-456-7890"))
	fmt.Println(ParsePhoneWithRegex("123-d asa78d    90"))
	fmt.Println(ParsePhoneWithRegex("123-75 7+ =-7 890"))

	fmt.Println(Anagram("moon", "noom"))
	fmt.Println(Anagram("pool", "loop"))
	fmt.Println(Anagram("frds", "srfd"))
	fmt.Println(Anagram("dsalae", "dsavae"))
	fmt.Println(Anagram("өрт", "төр"))

	fmt.Println(FindEvens([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}))

	fmt.Println(SliceProduct([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}))
	fmt.Println(SliceProduct([]int{2, -2, 5, 6, -6}))

	fmt.Println(Unique([]int{1, 1, 2, 4, 9, 7, 5, 7, 6, 9, 4, 7, 8}))

	fmt.Println(InvertMap(map[string]int{
		"henry": 4,
		"jim":   5,
	}))

	fmt.Println(TopCharacters("mississippi", 2))
	fmt.Println(TopCharacters("қарақалпақ", 2))
}

// ParsePhone parses a string of numbers into the format (123) 456-7890.
// This function should handle any number of extraneous spaces and dashes.
// All inputs will have 10 numbers and maybe extra spaces and dashes.
// For example, ParsePhone("123-456-7890") => "(123) 456-7890"
//              ParsePhone("1 2 3 4 5 6 7 8 9 0") => "(123) 456-7890"
func ParsePhone(phone string) string {
	phone_clear := strings.Replace(strings.Replace(phone, " ", "", -1), "-", "", -1) //first replace empty spaces, then dashes

	if len(phone_clear) == 10 {
		return "(" + phone_clear[:3] + ") " + phone_clear[3:6] + "-" + phone_clear[6:]
	}
	return "Invalid format or not enough numbers"
}

func ParsePhoneWithRegex(phone string) string {
	regex := regexp.MustCompile("[0-9]+") //using regex, filters only numbers

	if phone_clear := strings.Join(regex.FindAllString(phone, -1), ""); len(phone_clear) == 10 { //Join string because it was []string
		return "(" + phone_clear[:3] + ") " + phone_clear[3:6] + "-" + phone_clear[6:]
	}
	return "Invalid format or not enough numbers"
}

// Anagram tests whether the two strings are anagrams of each other.
// This function is NOT case sensitive and should handle UTF-8
func Anagram(s1, s2 string) bool {
	s1 = strings.ToLower(s1)
	s2 = strings.ToLower(s2)

	str1 := strings.Split(s1, "")
	str2 := strings.Split(s2, "")

	sort.Strings(str1)
	sort.Strings(str2)

	if len(str1) != len(str2) {
		return false
	}

	for i, j := range str1 {
		if j != str2[i] {
			return false
		}
	}

	return true
}

// FindEvens filters out all odd numbers from input slice.
// Result should retain the same ordering as the input.
func FindEvens(e []int) []int {
	var evens []int
	for _, e_num := range e {
		if e_num%2 == 0 {
			evens = append(evens, e_num)
		}
	}
	return evens
}

// SliceProduct returns the product of all elements in the slice.
// For example, SliceProduct([]int{1, 2, 3}) => 6
func SliceProduct(e []int) int {
	var sum int
	for _, n := range e {
		sum = sum + n
	}
	return sum
}

// Unique finds all distinct elements in the input array.
// Result should retain the same ordering as the input.
func Unique(e []int) []int {
	var unique []int
	m := make(map[int]bool)
	for _, n := range e {
		if _, v := m[n]; !v {
			m[n] = true
			unique = append(unique, n)
		}
	}
	return unique
}

// InvertMap inverts a mapping of strings to ints into a mapping of ints to strings.
// Each value should become a key, and the original key will become the corresponding value.
// For this function, you can assume each value is unique.
func InvertMap(kv map[string]int) map[int]string {
	invertedMap := make(map[int]string)

	for key, value := range kv {
		invertedMap[value] = key
	}
	return invertedMap
}

// TopCharacters finds characters that appear more than k times in the string.
// The result is the set of characters along with their occurrences.
// This function MUST handle UTF-8 characters.
func TopCharacters(s string, k int) map[rune]int {
	topCharacters := make(map[rune]int)

	for _, l := range s {
		if c := strings.Count(s, string(l)); c > k {
			topCharacters[l] = c
		}
	}
	return topCharacters
}
