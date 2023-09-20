package entities

type Policy struct {
	Version   string
	Id        string
	Statement []Statement
}

type Statement struct {
	Sid       string
	Effect    string
	Principal interface{}
	Action    string
	Resource  string
}
