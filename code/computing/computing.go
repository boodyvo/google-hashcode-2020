package computing

import (
	"fmt"
	"sort"
)

var scores []int

type ByCustomeScore []*Library
func (a ByCustomeScore) Len() int           { return len(a) }
func (a ByCustomeScore) Less(i, j int) bool { return a[i].Evaluation > a[j].Evaluation }
func (a ByCustomeScore) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

type ByBookScore    []int
func (a ByBookScore) Len() int           { return len(a) }
func (a ByBookScore) Less(i, j int) bool { return scores[a[i]] > scores[a[j]] }
func (a ByBookScore) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

func (library *Library) SetLibraryEvaluation(used *[]bool, daysLeft int) {
	res := 0
	daysLeft -= library.T
	it := 0
	n := len(library.Books)
	for i := 0; it < n && i < daysLeft; i++ {
		for j := 0; it < n && j < library.T; j++ {
			for it < n && (*used)[library.Books[it]] {
				it++
			}
			if it < n {
				res += scores[library.Books[it]]
			}
		}
	}

	library.Evaluation = float32(res) / float32(library.T)
}

func (library *Library) SortBooksByScore() {
	sort.Sort(ByBookScore(library.Books))
}

func Computing(B, L, D int, ss []int, libraries []*Library, printf func(f string, a ...interface{})) {
	scores = ss
	used := make([]bool, B)

	//books := make([]bool, B)
	s := 0
	librariesRes := make([]LibraryRes, 0)
	curT := 0

	for i := 0; i < L; i++ {
		libraries[i].SortBooksByScore()
		daysLeft := D - curT
		libraries[i].SetLibraryEvaluation(&used, daysLeft)
	}

	sort.Sort(ByCustomeScore(libraries))
	librariesUsed := make([]bool, L)

	for i := 0; i < L && curT < D; i++ {
		for j := 0; j < L; j++ {
			libraries[j].SetLibraryEvaluation(&used, D - curT)
		}
		sort.Sort(ByCustomeScore(libraries))
		number := 0
		for number < L && librariesUsed[libraries[number].Id] {
			number++
		}
		if number >= L {
			break
		}

		curLibrary := libraries[number].Id
		librariesUsed[curLibrary] = true

		curT += libraries[curLibrary].T
		daysLeft := D - curT

		if daysLeft <= 0 {
			break
		}
		librariesRes = append(librariesRes, LibraryRes{
			Id: libraries[curLibrary].Id,
		})
		daysLeft *= libraries[curLibrary].M

		it := 0
		n := len(libraries[curLibrary].Books)
		for il := 0; it < n && il < daysLeft; il++ {
			for jl := 0; it < n && jl < libraries[curLibrary].T; jl++ {
				for it < n && used[libraries[curLibrary].Books[it]] {
					it++
				}
				if it < n {
					librariesRes[i].Books = append(librariesRes[i].Books, libraries[curLibrary].Books[it])
					used[libraries[curLibrary].Books[it]] = true
					s += scores[libraries[curLibrary].Books[it]]
					// scores[libraries[curLibrary].Books[it]]
				}
			}
		}

		//librariesRes[i].Books = libraries[i].Books[:daysLeft]
		//for _, bb := range librariesRes[i].Books {
		//	if !books[bb] {
		//		s += scores[bb]
		//	}
		//	books[bb] = true
		//}
		//booksRes = append(booksRes, librariesRes[i].Books...)
	}

	fmt.Println(s)
	printf("%d\n", len(librariesRes))
	for _, libr := range librariesRes {
		printf("%d %d\n", libr.Id, len(libr.Books))
		for j := 0; j < len(libr.Books); j++ {
			printf("%d", libr.Books[j])
			if j == len(libr.Books) - 1 {
				printf("\n")
			} else {
				printf(" ")
			}
		}
	}
}


//librariesResBest := make([]LibraryRes, 0)
//bestScore := 0


//for it := 0; it < 100; it++ {
//	librariesRes := make([]LibraryRes, 0)
//	books := make([]bool, B)
//	used := make([]bool, L)
//	curT := 0
//	s := 0
//	i := 0
//	for curT < D {
//		number := rand.Intn(L)
//		for used[number] {
//			number = rand.Intn(L)
//		}
//		used[number] = true
//
//		curT += libraries[number].T
//		librariesRes = append(librariesRes, LibraryRes{
//			Id: libraries[number].Id,
//		})
//		daysLeft := D - curT
//
//		if daysLeft <= 0 {
//			break
//		}
//		if daysLeft > len(libraries[number].Books) {
//			daysLeft = len(libraries[number].Books)
//		}
//
//		librariesRes[i].Books = libraries[number].Books[:daysLeft]
//		for _, bb := range librariesRes[i].Books {
//			if !books[bb] {
//				s += scores[bb]
//			}
//			books[bb] = true
//		}
//		i++
//		//booksRes = append(booksRes, librariesRes[number].Books...)
//	}
//	if s > bestScore {
//		bestScore = s
//		librariesResBest = librariesRes
//	}
//}
//
//fmt.Print("score", bestScore)
//
//printf("%d\n", len(librariesResBest) - 1)
//for _, libr := range librariesResBest {
//	printf("%d %d\n", libr.Id, len(libr.Books))
//	for j := 0; j < len(libr.Books); j++ {
//		printf("%d", libr.Books[j])
//		if j == len(libr.Books) - 1 {
//			printf("\n")
//		} else {
//			printf(" ")
//		}
//	}
//}