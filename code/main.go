package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"

	comp "./computing"
)



var results int64

func setResults(result int64) {
	results = result
}

func readInput(scanf func(f string, a ...interface{})) (int, int, []int, error) {
	var n, m int
	scanf("%d %d\n", &m, &n)
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		scanf("%d", &arr[i])
	}

	return n, m, arr, nil
}

func main() {
	var fileName = flag.String("in", "a_example.in", "help message for flagname")
	var fileNameOut = flag.String("out", "a_example.out", "help message for flagname")
	flag.Parse()

	fmt.Println(*fileName)
	fmt.Println(*fileNameOut)

	//fileName := "e_also_big.in"
	file, err := os.Open(*fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	//fileNameOut := "e_also_big.out"
	fileOut, err := os.Create(*fileNameOut)
	if err != nil {
		fileOut, err = os.Open(*fileNameOut)
		if err != nil {
			log.Fatal(err)
		}
	}

	defer fileOut.Close()

	reader := bufio.NewReader(file)
	writer := bufio.NewWriter(fileOut)
	scanf := func(f string, a ...interface{}) { _, _ = fmt.Fscanf(reader, f, a...) }
	printf := func(f string, a ...interface{}) { _, _ = fmt.Fprintf(writer, f, a...) }

	defer writer.Flush()

	results = 0
	n, m, types, err := readInput(scanf)
	if err != nil {
		log.Fatal("Cannot read file")

		return
	}

	fmt.Println("start computing ...")
	comp.Computing(m, n, types, printf)
}