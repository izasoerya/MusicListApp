package models

type Request struct {
	ID     uint32 `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Date   string `json:"date"`
	Link   string `json:"link"`
	CategoryID *string   // Foreign key column

}

type Music struct {
	ID     *uint32 `json:"id"`
	Title  *string `json:"title"`
	Author *string `json:"author"`
	Date   *string `json:"date"`
	Link   *string `json:"link"`
	CategoryID *string   // Foreign key column
}

type Auth struct {
	ID       *uint64 `gorm:"primary_key;autoIncrement"`
	Username *string `json:"username"`
	Password *string `json:"password"`
}
