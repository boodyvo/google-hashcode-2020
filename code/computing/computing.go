package computing

import (
	"fmt"
	"sort"
)

type ByDays []*Library

func (a ByDays) Len() int           { return len(a) }
func (a ByDays) Less(i, j int) bool { return a[i].T < a[j].T }
func (a ByDays) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

func Computing(B, L, D int, scores []int, libraries []*Library, printf func(f string, a ...interface{})) {
	sort.Sort(ByDays(libraries))
	for i := 0; i < L; i++ {
		fmt.Println(libraries[i].Id, libraries[i].T, libraries[i].M)
	}

	booksRes := make([]int, 0, B)
	librariesRes := make([]LibraryRes, 0)
	curT := 0
	for i := 0; i < L && curT < D; i++ {
		curT += libraries[i].T
		librariesRes = append(librariesRes, LibraryRes{
			Id: libraries[i].Id,
		})
		daysLeft := D - curT

		if daysLeft <= 0 {
			break
		}

		librariesRes[i].Books = libraries[i].Books[:daysLeft]
		booksRes = append(booksRes, librariesRes[i].Books...)
	}

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