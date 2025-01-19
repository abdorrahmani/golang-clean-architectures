package post

type Post struct {
	ID    uint   `gorm:"primaryKey" json:"id"`
	Title string `gorm:"size:255,not null" json:"title"`
	Body  string `gorm:"type:text;not null" json:"body"`
}
