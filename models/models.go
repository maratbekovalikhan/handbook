package models

type Entry struct {
	ID       int
	Title    string
	Content  string
	Category string
}

type Category struct {
	ID   int
	Name string
}
