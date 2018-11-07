package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	byt, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, fmt.Sprintf(`{"error":"%v"}`, err))
		os.Exit(1)
	}
	fmt.Fprintf(os.Stdout, fmt.Sprintf(`{"message":"%v"}`, string(byt)))
}
