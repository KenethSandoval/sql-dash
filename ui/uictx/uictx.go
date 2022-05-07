package uictx

type Ctx struct {
	Screen  [2]int
	Content [2]int
	Loading bool
}

func New() Ctx {
	return Ctx{
		Screen:  [2]int{0, 0},
		Content: [2]int{0, 0},
		Loading: false,
	}
}
