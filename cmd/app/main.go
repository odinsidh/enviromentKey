package main

import (
	"enviromentkey/internal/extractor"
	"fmt"
	"log"
)

func main() {
	r, err := extractor.GetExtractor("default")
	if err != nil {
		log.Fatal(err)
	}

	dtoLoc, err := r.Handle("request.txt")
	if err != nil {
		log.Fatal(err)
	}

	for index, value := range dtoLoc {
		fmt.Println(index, value)
	}

}
