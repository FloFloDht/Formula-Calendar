package entity

// definition des entités

type Track struct {
	ID       int32
	Name     string
	Location int32
}

type Category struct {
	ID int32
	Name string
}

type GrandPrix struct {
	ID int32
	Name string
	Track Track
	Categories []Category
	Date string
}

