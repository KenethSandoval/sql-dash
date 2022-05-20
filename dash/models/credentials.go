package models

type Credential struct {
	Name           string
	Host           string
	InsertPriv     string
	SelectPriv     string
	UpdatePriv     string
	DeletePriv     string
	CreatePriv     string
	DropPriv       string
	ReloadPriv     string
	ShutdownPriv   string
	ProcessPriv    string
	FilePriv       string
	GrantPriv      string
	ReferencesPriv string
	IndexPriv      string
	AlterPriv      string
}

func (crendential Credential) FilterValue() string {
	return crendential.Name
}

func (credential Credential) Title() string {
	return credential.Name
}

func (credential Credential) Description() string {
	// TODO: return type db
	//return "mysql"
	return credential.Host
}
