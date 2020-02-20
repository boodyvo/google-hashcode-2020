package computing

import (
	"fmt"
	"math"
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
	n := len(library.Books)
	daysLeft -= library.T
	it := 0
	for i := 0; it < n && i < daysLeft; i++ {
		for j := 0; it < n && j < library.M; j++ {
			for it < n && (*used)[library.Books[it]] {
				it++
			}
			if it < n {
				res += scores[library.Books[it]]
			}
		}
	}
	//for i := 0; i < n; i++ {
	//	if !(*used)[library.Books[i]] {
	//		res += scores[library.Books[i]]
	//	}
	//}

	library.Evaluation = float32(res)  / float32(library.T)
}

func BackEvaluation(library *Library, used *[]bool, daysLeft int) {
	res := 0
	n := len(library.Books)
	daysLeft -= library.T
	it := 0
	for i := 0; it < n && i < daysLeft; i++ {
		for j := 0; it < n && j < library.M; j++ {
			for it < n && (*used)[library.Books[it]] {
				it++
			}
			if it < n {
				res += scores[library.Books[it]]
			}
		}
	}
	//for i := 0; i < n; i++ {
	//	if !(*used)[library.Books[i]] {
	//		res += scores[library.Books[i]]
	//	}
	//}

	library.Evaluation = float32(res)  / float32(library.T)
}

func Back3Evaluation(library *Library, used *[]bool, daysLeft int) {
	res := 0
	n := len(library.Books)
	daysLeft -= library.T
	it := 0
	for i := 0; it < n && i < daysLeft; i++ {
		for j := 0; it < n && j < library.M; j++ {
			for it < n && (*used)[library.Books[it]] {
				it++
			}
			if it < n {
				res += scores[library.Books[it]]
			}
		}
	}
	//for i := 0; i < n; i++ {
	//	if !(*used)[library.Books[i]] {
	//		res += scores[library.Books[i]]
	//	}
	//}

	library.Evaluation = float32(res)  / float32(library.T * library.T)
}

func BackEmptyEvaluation(library *Library, used *[]bool, daysLeft int) {
	res := 0
	n := len(library.Books)
	daysLeft -= library.T
	it := 0
	for i := 0; it < n && i < daysLeft; i++ {
		for j := 0; it < n && j < library.M; j++ {
			for it < n && (*used)[library.Books[it]] {
				it++
			}
			if it < n {
				res += scores[library.Books[it]]
			}
		}
	}
	//for i := 0; i < n; i++ {
	//	if !(*used)[library.Books[i]] {
	//		res += scores[library.Books[i]]
	//	}
	//}

	library.Evaluation = float32(res)  / float32(library.T * library.T * library.T)
}

func ForwardEvaluation(library *Library, used *[]bool, daysLeft int) {
	res := 0
	n := len(library.Books)

	for i := 0; i < n; i++ {
		if !(*used)[library.Books[i]] {
			res += scores[library.Books[i]]
		}
	}

	library.Evaluation = float32(res) / float32(library.T)
}

func Forward3Evaluation(library *Library, used *[]bool, daysLeft int) {
	res := 0
	n := len(library.Books)

	for i := 0; i < n; i++ {
		if !(*used)[library.Books[i]] {
			res += scores[library.Books[i]]
		}
	}

	library.Evaluation = float32(res) / float32(library.T * library.T)
}

func StraightEvaluation(library *Library, used *[]bool, daysLeft int) {
	res := 0
	n := len(library.Books)

	for i := 0; i < n; i++ {
		res += scores[library.Books[i]]
	}

	library.Evaluation = float32(res) / float32(library.T)
}

func Straight3Evaluation(library *Library, used *[]bool, daysLeft int) {
	res := 0
	n := len(library.Books)

	for i := 0; i < n; i++ {
		res += scores[library.Books[i]]
	}

	library.Evaluation = float32(res) / float32(library.T * library.T * library.T)
}

func StraightEmptyEvaluation(library *Library, used *[]bool, daysLeft int) {
	res := 0
	n := len(library.Books)

	for i := 0; i < n; i++ {
		res += scores[library.Books[i]]
	}

	library.Evaluation = float32(res)
}

func ForwardEmptyEvaluation(library *Library, used *[]bool, daysLeft int) {
	res := 0
	n := len(library.Books)

	for i := 0; i < n; i++ {
		if !(*used)[library.Books[i]] {
			res += scores[library.Books[i]]
		}
	}

	library.Evaluation = float32(res)
}

func RandomEvaluation(library *Library, used *[]bool, daysLeft int, tt int) {
	res := 0
	n := len(library.Books)

	for i := 0; i < n; i++ {
		if !(*used)[library.Books[i]] {
			res += scores[library.Books[i]]
		}
	}

	newtt := math.Pow(float64(library.T), float64(tt) / float64(10))
	library.Evaluation = float32(res) / float32(newtt)
}

func GetRandomEvaluation(tt int) func(library *Library, used *[]bool, daysLeft int) {
	return func (library *Library, used *[]bool, daysLeft int) {
		res := 0
		n := len(library.Books)
		daysLeft -= library.T
		it := 0
		for i := 0; it < n && i < daysLeft; i++ {
			for j := 0; it < n && j < library.M; j++ {
				for it < n && (*used)[library.Books[it]] {
					it++
				}
				if it < n {
					res += scores[library.Books[it]]
				}
			}
		}

		newtt := math.Pow(float64(library.T), float64(tt) / float64(10))
		library.Evaluation = float32(res) / float32(newtt)
	}
}

func CreateEvaluationFunction(x int) func(library *Library, used *[]bool, daysLeft int) {
	switch x {
	case 1:
		return BackEvaluation
	case 2:
		return Back3Evaluation
	case 3:
		return BackEmptyEvaluation
	case 4:
		return ForwardEvaluation
	case 5:
		return Forward3Evaluation
	case 6:
		return ForwardEmptyEvaluation
	case 7:
		return StraightEvaluation
	case 8:
		return Straight3Evaluation
	case 9:
		return StraightEmptyEvaluation
	default:
		return StraightEmptyEvaluation
	}
}

func (library *Library) SortBooksByScore() {
	sort.Sort(ByBookScore(library.Books))
}

func Computing(B, L, D int, ss []int, libraries []*Library, printf func(f string, a ...interface{})) {
	bestRes := make([]LibraryRes, 0)
	bestS := 0
	for x := 0; x < 30; x += 3 {
		curEval := GetRandomEvaluation(x)
		scores = ss
		used := make([]bool, B)

		//books := make([]bool, B)
		s := 0
		librariesRes := make([]LibraryRes, 0)
		curT := 0

		for i := 0; i < L; i++ {
			libraries[i].SortBooksByScore()
			daysLeft := D - curT

			curEval(libraries[i], &used, daysLeft)
			//libraries[i].SetLibraryEvaluation(&used, daysLeft)
		}

		sort.Sort(ByCustomeScore(libraries))
		librariesUsed := make([]bool, L)

		for i := 0; i < L && curT < D; i++ {
			//fmt.Println("iteration", i, curT)
			//uu := make([]bool, B)
			if i % 100 == 0 {
				for j := 0; j < L; j++ {
					//libraries[j].SetLibraryEvaluation(&used, D-curT)
					curEval(libraries[i], &used, D-curT)
				}
			}
			sort.Sort(ByCustomeScore(libraries))
			number := 0
			for number < L && librariesUsed[libraries[number].Id] {
				number++
			}
			if number >= L {
				i--
				continue
			}

			librariesUsed[libraries[number].Id] = true

			curT += libraries[number].T
			daysLeft := D - curT

			if daysLeft <= 0 {
				break
			}
			librariesRes = append(librariesRes, LibraryRes{
				Id: libraries[number].Id,
			})

			it := 0
			n := len(libraries[number].Books)
			for il := 0; it < n && il < daysLeft; il++ {
				for jl := 0; it < n && jl < libraries[number].M; jl++ {
					for it < n && used[libraries[number].Books[it]] {
						it++
					}
					if it < n {
						librariesRes[i].Books = append(librariesRes[i].Books, libraries[number].Books[it])
						used[libraries[number].Books[it]] = true
						// scores[libraries[curLibrary].Books[it]]
					}
				}
			}
		}
		used = make([]bool, B)
		for _, libr := range librariesRes {
			for j := 0; j < len(libr.Books); j++ {
				if !used[libr.Books[j]] {
					s += scores[libr.Books[j]]
					used[libr.Books[j]] = true
				}
			}
		}
		if s > bestS {
			bestS = s
			bestRes = librariesRes
		}
		fmt.Println(s)
	}

	//used := make([]bool, B)
	printf("%d\n", len(bestRes))
	fmt.Println(bestS)
	for _, libr := range bestRes {
		printf("%d %d\n", libr.Id, len(libr.Books))
		for j := 0; j < len(libr.Books); j++ {
			printf("%d", libr.Books[j])
			if j == len(libr.Books) - 1 {
				printf("\n")
			} else {
				printf(" ")
			}
			//if !used[libr.Books[j]] {
			//	s += scores[libr.Books[j]]
			//	used[libr.Books[j]] = true
			//}
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