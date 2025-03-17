package entity

type Track struct {
	ID       int32  `json:"id"`
	Name     string `json:"name"`
	Location string `json:"location"`
}

type Category struct {
	ID   int32  `json:"id"`
	Name string `json:"name"`
}

type GrandPrix struct {
	ID         int32      `json:"id"`
	Name       string     `json:"name"`
	Track      Track      `json:"track"`
	Categories []Category `json:"categories"`
	Date       string     `json:"date"`
}
