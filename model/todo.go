package model

type Todo struct {
	Id int `db:"id"`
	Text string `db:"text"`
	Checked bool `db:"checked"`
}