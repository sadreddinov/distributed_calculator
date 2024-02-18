package models

var Operations *Operation = &Operation{Plus: "10", Minus: "10", Multiply: "10", Divide: "10"}

type Operation struct {
	Plus     string `json: "plus"`
	Minus    string `json: "minus"`
	Multiply string `json: "multiply"`
	Divide   string `json: "divide"`
}
