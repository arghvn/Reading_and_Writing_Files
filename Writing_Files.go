package main

// Writing files in Go follows similar patterns to the ones we saw earlier for reading.

import (
	"bufio"
	"fmt"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	d1 := []byte("hello\ngo\n")
	err := os.WriteFile("/tmp/dat1", d1, 0644)
	check(err)

	f, err := os.Create("/tmp/dat2")
	check(err)

	defer f.Close()

	// To start, here’s how to dump a string (or just bytes) into a file.
	// For more granular writes, open a file for writing.
	// It’s idiomatic to defer a Close immediately after opening a file.
	// You can Write byte slices as you’d expect.

	d2 := []byte{115, 111, 109, 101, 10}
	n2, err := f.Write(d2)
	check(err)
	fmt.Printf("wrote %d bytes\n", n2)

	n3, err := f.WriteString("writes\n")
	check(err)
	fmt.Printf("wrote %d bytes\n", n3)

	f.Sync()

	w := bufio.NewWriter(f)
	n4, err := w.WriteString("buffered\n")
	check(err)
	fmt.Printf("wrote %d bytes\n", n4)

	w.Flush()

}

// Issue a Sync to flush writes to stable storage.bufio provides buffered writers
//  in addition to the buffered readers we saw earlier.Use Flush to ensure all buffered operations
//  have been applied to the underlying writer.
// Try running the file-writing code.
// then check the contents of the written files.

// output :
// wrote 5 bytes
// wrote 7 bytes
// wrote 9 bytes

// $ cat /tmp/dat1
// hello
// go

// $ cat /tmp/dat2
// some
// writes
// buffered
