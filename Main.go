// TODO: Threading while creating from a wordlist

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
	n   int    // No. of file
	ext string // Extension
	t   int    // Threads
	v   bool   // Verbosity
	dl  int    // data length (random)
	dc  string // Custom data string
	fl  int    // Length of filename
	pre string // Prefix
	suf string // Suffix
	// w   string // Wordlist with filenames
}

var args arguments
var fileCounter int

func main() {
	// User args
	flag.IntVar(&args.n, "n", 1, "Number of files to create (Default: 1)")
	flag.StringVar(&args.ext, "e", "txt", "Extension to use (Default: txt")
	flag.IntVar(&args.t, "t", 1, "Number of threads (Default: 1)")
	flag.BoolVar(&args.v, "v", false, "Verbosity (Default: Off)")
	flag.IntVar(&args.dl, "dl", 0, "Length of random data to write in each file (Default: 0)")
	flag.StringVar(&args.dc, "dc", "", "Write custom data to each file (Not to be used with -dl)")
	flag.IntVar(&args.fl, "fl", 10, "Length of the filename (Default: 10)")
	flag.StringVar(&args.pre, "prefix", "", "Prefix")
	flag.StringVar(&args.suf, "suffix", "", "Suffix")
	// flag.StringVar(&args.w, "w", "", "Wordlist that contains filenames")
	flag.Parse()

	// Some checks
	flagChecks()

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

func flagChecks() {
	if args.t > args.n {
		log.Fatalln("[-] More no. of threads than files. (t <= n))")
		os.Exit(1)
	}

	if args.dl != 0 && args.dc != "" {
		log.Fatalln("[-] -dc and -dl flags not to be used at the same time.")
		os.Exit(1)
	}

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

func writeData(fd *os.File, data *string) {
	filename := fd.Name()
	_, err := fd.WriteString(*data)
	if err != nil {
		log.Fatalln("[-] Couldn't write to " + filename)
	}
}

// Create the file
func createFile(n int) {
	prefix := args.pre
	suffix := args.suf
	for i := 0; i < n; i++ {

		// a random filename
		filename := prefix + randomName(args.fl) + suffix + "." + args.ext
		fd, err := os.Create(filename)
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}
		defer fd.Close()

		// Write random data in each file
		if args.dl > 0 {
			data := randomName(args.dl)
			writeData(fd, &data)
		}

		// Write user-provided data in each file
		if args.dc != "" {
			writeData(fd, &args.dc)
		}

		// Number of files created
		fileCounter += 1
		if args.v {
			fmt.Println("[+] " + filename + " Created")
		}
	}

}

// func createFileWordlist() {
// 	fd, err := os.Open(args.w)
// 	if err != nil {
// 		log.Fatalln("[-] Unable to open " + args.w)
// 		os.Exit(1)
// 	}
// 	scanner := bufio.NewScanner(fd)
// 	for scanner.Scan() {
// 		filename := scanner.Text()
// 		fd, err := os.Create(filename + "." + args.ext)
// 		if err != nil {
// 			log.Fatalln("[-] Couldnot create " + filename)
// 			os.Exit(1)
// 		}
// 		defer fd.Close()
// 		// Write random data in each file
// 		if args.dl > 0 {
// 			data := randomName(args.dl)
// 			writeData(fd, &data)
// 		}

// 		// Write user-provided data in each file
// 		if args.dc != "" {
// 			writeData(fd, &args.dc)
// 		}
// 	}
// }
