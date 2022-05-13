package models

type Credential struct {
	Name string
	Host string
}

func (crendential Credential) FilterValue() string {
	return crendential.Name
}

func (credential Credential) Title() string {
	return credential.Name
}

func (credential Credential) Description() string {
	return "mysql"
}
