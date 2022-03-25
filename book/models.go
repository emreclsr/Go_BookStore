package book

type Book struct {
	ID         int `gorm:"primary_key"`
	Name       string
	PageNumber int
	Stock      int
	Price      float64
	StockCode  int
	ISBN       string
	Author     Author `gorm:"ForeignKey:AuthorID"`
	Deleted    bool
	AuthorID   int
}

type Author struct {
	ID   int
	Name string
}
