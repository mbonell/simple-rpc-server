package models

type Request struct {
	Function string
	Args     []string
}

type Response struct {
	Code  int
	Error error
	Msg   []byte
}
