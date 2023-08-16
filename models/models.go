package models

type Request struct {
	ID     uint32 `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Date   string `json:"date"`
	Link   string `json:"link"`
}

type Music struct {
	ID     *uint32 `json:"id"`
	Title  *string `json:"title"`
	Author *string `json:"author"`
	Date   *string `json:"date"`
	Link   *string `json:"link"`
}

type Auth struct {
	Username *string `json:"username"`
	Password *string `json:"password"`
}
