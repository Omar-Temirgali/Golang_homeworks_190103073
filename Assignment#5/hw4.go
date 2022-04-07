// Homework 4: Concurrency
// Due February 21, 2017 at 11:59pm
package main

import (
	"bufio"
	"io"
	"log"
	"os"
	"strconv"
	"sync"
)

func main() {
	FileSum("input.txt", "output.txt")
}

// Problem 1a: File processing
// You will be provided an input file consisting of integers, one on each line.
// Your task is to read the input file, sum all the integers, and write the
// result to a separate file.

// FileSum sums the integers in input and writes them to an output file.
// The two parameters, input and output, are the filenames of those files.
// You should expect your input to end with a newline, and the output should
// have a newline after the result.
func FileSum(input, output string) {
	sum := 0

	inputFile, err := os.Open(input)
	if err != nil {
		log.Fatal(err)
	}

	defer inputFile.Close()

	scanner := bufio.NewScanner(inputFile)
	for scanner.Scan() {
		intVar, err := strconv.Atoi(scanner.Text())
		sum = sum + intVar
		if err != nil {
			log.Fatal(err)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	outputFile, err := os.Create(output)
	if err != nil {
		log.Fatal(err)
	}
	defer inputFile.Close()

	w := bufio.NewWriter(outputFile)
	w.WriteString(strconv.Itoa(sum) + "\n")

	w.Flush()
}

// Problem 1b: IO processing with interfaces
// You must do the exact same task as above, but instead of being passed 2
// filenames, you are passed 2 interfaces: io.Reader and io.Writer.
// See https://golang.org/pkg/io/ for information about these two interfaces.
// Note that os.Open returns an io.Reader, and os.Create returns an io.Writer.

// IOSum sums the integers in input and writes them to output
// The two parameters, input and output, are interfaces for io.Reader and
// io.Writer. The type signatures for these interfaces is in the Go
// documentation.
// You should expect your input to end with a newline, and the output should
// have a newline after the result.
func IOSum(input io.Reader, output io.Writer) {
	inputFileSc := bufio.NewScanner(input)
	sum := 0
	for inputFileSc.Scan() {
		n, err := strconv.Atoi(inputFileSc.Text())
		if err != nil {
			log.Fatal(err)
		}
		sum += n
	}

	outputFile := bufio.NewWriter(output)
	outputFile.WriteString(strconv.Itoa(sum) + "\n")
	outputFile.Flush()
}

// Problem 2: Concurrent map access
// Maps in Go [are not safe for concurrent use](https://golang.org/doc/faq#atomic_maps).
// For this assignment, you will be building a custom map type that allows for
// concurrent access to the map using mutexes.
// The map is expected to have concurrent readers but only 1 writer can have
// access to the map.

// PennDirectory is a mapping from PennID number to PennKey (12345678 -> adelq).
// You may only add *private* fields to this struct.
// Hint: Use an embedded sync.RWMutex, see lecture 2 for a review on embedding
type PennDirectory struct {
	mx        sync.RWMutex
	directory map[int]string
}

// Add inserts a new student to the Penn Directory.
// Add should obtain a write lock, and should not allow any concurrent reads or
// writes to the map.
// You may NOT write over existing data - simply raise a warning.
func (d *PennDirectory) Add(id int, name string) {
	d.mx.Lock()
	d.directory[id] = name
	d.mx.Unlock()
}

// Get fetches a student from the Penn Directory by their PennID.
// Get should obtain a read lock, and should allow concurrent read access but
// not write access.
func (d *PennDirectory) Get(id int) string {
	d.mx.RLock()
	name := d.directory[id]
	d.mx.RUnlock()
	return name
}

// Remove deletes a student to the Penn Directory.
// Remove should obtain a write lock, and should not allow any concurrent reads
// or writes to the map.
func (d *PennDirectory) Remove(id int) {
	d.mx.Lock()
	delete(d.directory, id)
	d.mx.Unlock()
}
