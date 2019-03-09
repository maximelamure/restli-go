package client

type Order struct {
	ID     string
	Number int
	Label  string
}

type User struct {
	ID     string
	Number int
	Label  string
}

type Result struct {
	resps map[string]Response
}
