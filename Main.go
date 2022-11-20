// Create a bunch of files

// Number of files - Integer *
// Extension - string *
// Random Extension - Wordlist
// Threading - Int
// Write data - bool
// Random Data - bool
// CustomData - stirng

package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"sync"
	"time"
)

const (
	chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789_"
)

type arguments struct {
	n   int
	ext string
	t   int
	v   bool
}

var args arguments
var fileCounter int

func main() {
	// User args
	flag.IntVar(&args.n, "n", 1, "Number of files to create (Default: 0)")
	flag.StringVar(&args.ext, "e", "txt", "Extension to use")
	flag.IntVar(&args.t, "t", 1, "Number of threads(Default: 1)")
	flag.BoolVar(&args.v, "v", false, "Verbosity (Default: Off)")
	flag.Parse()

	if args.t > args.n {
		log.Fatalln("No. of Threads should be less than or equal to number of files. (t <= n))")
		return
	}

	// Seed to generate a random number
	rand.Seed(time.Now().UnixNano())

	filesAndThreads := filesInThreads()

	var wg sync.WaitGroup
	wg.Add(args.t)

	for _, filesInAThread := range filesAndThreads {
		files := filesInAThread

		go func() {
			createFile(files)
			wg.Done()
		}()
	}
	wg.Wait()

}

func filesInThreads() []int {
	filesPerThread := []int{}

	fileNum := args.n / args.t

	// Number of files to create in each thread
	for i := 0; i < args.t; i++ {
		filesPerThread = append(filesPerThread, fileNum)
	}

	// The numbers are not always distributed equally
	// The following will handle that case
	// If the given number is greater than the number of files that will be created
	if args.n > fileNum*len(filesPerThread) {
		remaining := args.n - fileNum*len(filesPerThread)
		for i := 0; i < remaining; i++ {
			filesPerThread[i]++
		}
	}
	return filesPerThread
}

// Returns Random String of given length
func randomName(length int) string {
	// holder for random string
	s := ""

	for i := 0; i < length; i++ {
		randomChar := string(chars[rand.Intn(len(chars))])
		s += randomChar
	}
	return s
}

// Create the file
func createFile(n int) {

	for i := 0; i < n; i++ {

		// a random filename
		filename := randomName(6) + "." + args.ext

		fd, err := os.Create(filename)
		if err != nil {
			log.Fatal(err)
		}
		defer fd.Close()

		// Number of files created
		fileCounter += 1
		if args.v {
			fmt.Println("[+] " + filename + " Created")

		}
	}
}
