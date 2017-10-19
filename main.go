package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/marusama/nonsense/buckets"
)

func main() {
	a, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}
	buckets.Print(uint32(a), false)
}
