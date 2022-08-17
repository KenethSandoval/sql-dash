package uictx

import "tuidb/dash"

type Ctx struct {
	Screen  [2]int
	Content [2]int
	Client  *dash.Dash
	Loading bool
}

func New(client *dash.Dash) Ctx {
	return Ctx{
		Screen:  [2]int{0, 0},
		Content: [2]int{0, 0},
		Client:  client,
		Loading: false,
	}
}
