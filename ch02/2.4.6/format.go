package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	fmt.Fprintf(os.Stdout, "write with os.Stdout as %v", time.Now())
}
