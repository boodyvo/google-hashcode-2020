package computing

type Library struct {
	N          int
	T          int
	M          int
	Id         int
	Books      []int
	Evaluation float32
}

type LibraryRes struct {
	Id    int
	Books []int
}
