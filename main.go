package main

import (
	"fmt"

	"github.com/marusama/nonsense/deadcode2"
)

func main() {
	f := "/Users/bator.tsyrendylykov/dev/projects/src/hello-go"
	res, err := deadcode2.Deadcode2(f)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	for _, s := range res {
		fmt.Println(s)
	}
}
