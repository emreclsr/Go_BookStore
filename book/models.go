package book

type Book struct {
	ID         int
	Name       string
	PageNumber int
	Stock      int
	Price      float64
	StockCode  int
	ISBN       string
	Author     Author
	Deleted    bool
}

type Author struct {
	ID   int
	Name string
}
