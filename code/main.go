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

func readInput(scanf func(f string, a ...interface{})) (int, int, int, []int, []*comp.Library) {
	var B, L, D int
	scanf("%d %d %d\n", &B, &L, &D)
	scores := make([]int, B)
	for i := 0; i < B; i++ {
		scanf("%d", &scores[i])
	}
	scanf("\n")
	libraries := make([]*comp.Library, L)
	for i := 0; i < L; i++ {
		libraries[i] = &comp.Library{
			Id: i,
		}
		scanf("%d %d %d\n", &libraries[i].N, &libraries[i].T, &libraries[i].M)
		libraries[i].Books = make([]int, libraries[i].N)
		for j := 0; j < libraries[i].N; j++ {
			scanf("%d", &libraries[i].Books[j])
		}
		scanf("\n")
	}

	return B, L, D, scores, libraries
}

func main() {
	var fileName = flag.String("in", "b_read_on.txt", "help message for flagname")
	var fileNameOut = flag.String("out", "b_read_on.out", "help message for flagname")
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
	B, L, D, scores, libraries := readInput(scanf)
	if err != nil {
		log.Fatal("Cannot read file")

		return
	}

	fmt.Println("start computing ...")
	comp.Computing(B, L, D, scores, libraries, printf)
}