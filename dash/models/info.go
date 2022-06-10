package models

type Info struct {
	Version  string
	UserConn string
}

func (info Info) FilterValue() string {
	return info.Version
}

func (info Info) Title() string {
	return info.UserConn
}

func (info Info) Description() string {
	return info.Version
}
