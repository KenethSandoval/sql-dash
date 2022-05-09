package navigation

var Navigation = []string{}

type Model struct {
	CurrentId int
}

func (m *Model) NthTab(nth int) {
	if nth > len(Navigation) {
		nth = len(Navigation)
	} else if nth < 1 {
		nth = 1
	}
}
