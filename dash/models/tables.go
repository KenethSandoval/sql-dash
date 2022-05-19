package models

type Tables struct {
	Name  string
	Field string
}

type TableDescribe struct {
	Field   string
	Type    string
	Null    string
	Key     string
	Default string
	Extra   string
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
