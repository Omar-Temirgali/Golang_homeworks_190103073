// Homework 3: Interfaces
// Due February 14, 2017 at 11:59pm
package main

import (
	"fmt"
	"sort"
	"strings"
	"sync"
)

func main() {
	people := []*Person{
		NewPerson("Vlad", "Mathews"),
		NewPerson("Willie", "Stevenson"),
		NewPerson("Ahyan", "Hackett"),
		NewPerson("Tayyib", "Alston"),
		NewPerson("Samiha", "Rivas"),
		NewPerson("Kailum", "Bravo"),
		NewPerson("Sureira", "Alston"),
		NewPerson("Monica ", "Winter"),
		NewPerson("Jacque", "Frederick"),
		NewPerson("Kaja", "Page"),
		NewPerson("Sureira", "Alston"),
	}

	fmt.Println(people)
	sort.Sort(PersonSlice(people))
	for _, person := range people {
		fmt.Println(*person)
	}

	fmt.Println(IsPalindrome(sort.IntSlice{1, 2, 1}))
	fmt.Println(IsPalindrome(sort.IntSlice{1, 4, 2, 1}))
	fmt.Println(IsPalindrome(sort.StringSlice{"DWD"}))
	fmt.Println(IsPalindrome(sort.StringSlice{"DWD", "MLM"}))
	fmt.Println(IsPalindrome(sort.StringSlice{"DWD", "MLM", "DWD"}))

	slice := []int{24, 45, 75, 4, 35, 7}
	fmt.Println(Fold(slice, 2, func(i1, i2 int) int { return 1 }))
}

// Problem 1: Sorting Names
// Sorting in Go is done through interfaces!
// To sort a collection (such as a slice), the type must satisfy sort.Interface,
// which requires 3 methods: Len() int, Less(i, j int) bool, and Swap(i, j int).
// To actually sort a slice, you need to first implement all 3 methods on a
// custom type, and then call sort.Sort on your the PersonSlice type.
// See the Go documentation: https://golang.org/pkg/sort/ for full details.

// Person stores a simple profile. These should be sorted by alphabetical order
// by last name, followed by the first name, followed by the ID. You can assume
// the ID will be unique, but the names need not be unique.
// Sorting should be case-sensitive and UTF-8 aware.
type Person struct {
	ID        int
	FirstName string
	LastName  string
}

func (p Person) String() string {
	return fmt.Sprintf("%s %s %d", p.FirstName, p.LastName, p.ID)
}

type autoIncrement struct {
	sync.Mutex
	id int
}

func (a *autoIncrement) ID() (id int) {
	a.Lock()
	defer a.Unlock()

	id = a.id + 1
	a.id++
	return
}

var ai autoIncrement

type PersonSlice []*Person

func (a PersonSlice) Len() int      { return len(a) }
func (a PersonSlice) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a PersonSlice) Less(i, j int) bool {
	fName1 := strings.ToLower(a[i].FirstName)
	fName2 := strings.ToLower(a[j].FirstName)
	lName1 := strings.ToLower(a[i].LastName)
	lName2 := strings.ToLower(a[j].LastName)
	id1 := a[i].ID
	id2 := a[i].ID

	if lName1 == lName2 && fName1 == fName2 {
		return id1 < id2
	} else if lName1 == lName2 {
		return fName1 < fName2
	} else {
		return lName1 < lName2
	}
}

// NewPerson is a constructor for Person. ID should be assigned automatically in
// sequential order, starting at 1 for the first Person created.
func NewPerson(first, last string) *Person {
	return &Person{FirstName: first, LastName: last, ID: ai.ID()}
}

// Problem 2: IsPalindrome Redux
// Using a function that simply requires sort.Interface, you should be able to
// check if a sequence is a palindrome. You may use, adapt, or modify your code
// from HW0. Note that the input does not need to be a string: any type which
// satisfies sort.Interface can (and will) be used to test. This means that the
// only functionality you should use should come from the sort.Interface methods
// Ex: [1, 2, 1] => true

// IsPalindrome checks if the string is a palindrome.
// A palindrome is a string that reads the same backward as forward.
func IsPalindrome(s sort.Interface) bool {
	for i, j := 0, s.Len()-1; j > i; {
		if !s.Less(i, j) && !s.Less(j, i) {
			i++
			j--
		} else {
			return false
		}
	}
	return true
}

// Problem 3: Functional Programming
// Write a function Fold which applies a function repeatedly on a slice,
// producing a single value via repeated application of an input function.
// The behavior of Fold should be as follows:
//   - When s is empty, return v (default value)
//   - When s has 1 value (x0), apply f once: f(v, x0)
//   - When s has 2 values (x0, x1), apply f twice, from left to right: f(f(v, x0), x1)
//   - Continue this pattern recursively to obtain the final result.

// Fold applies a left to right application of f on s starting with v.
// Note the argument signature of f - func(int, int) int.
// This means f is a function which has 2 int arguments and returns an int.
func Fold(s []int, v int, f func(int, int) int) int {
	if len(s) == 0 {
		return v
	} else if len(s) == 1 {
		return f(v, s[0])
	} else if len(s) == 2 {
		return f(f(v, s[0]), s[1])
	}
	return f(v, s[v])
}
