package models

type Tables struct {
	Name string
	Size string
}

func (table Tables) FilterValue() string {
	return table.Name
}

func (table Tables) Title() string {
	return table.Name
}

func (table Tables) Description() string {
	// TODO: return type db
	return "mysql table size"
}
