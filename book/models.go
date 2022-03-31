package book

// Book book
//
// swagger:model Book
type Book struct {
	// The unique id of a book
	// Required: true
	// Minimum: 1
	ID int `gorm:"primary_key", json:"id"`

	// The name of a book
	// Required: true
	// Minimum length: 1
	Name string `json:"name"`

	// The page count of a book
	// Minimum: 1
	PageNumber int `json:"pageNumber"`

	// The stock of a book
	// Minimum: 1
	Stock int `json:"stock"`

	// The price of a book
	// Required: true
	// Minimum: 1
	Price float64 `json:"price"`

	// The stock code of book
	StockCode int `json:"stockCode"`

	// The ISBN of a book
	// Example: 978-3-17-148411-2
	ISBN string `json:"ISBN"`

	// The deleted flag of a book
	// Required: true
	Deleted bool `json:"deleted"`

	// The author id of a book
	// Required: true
	// Minimum: 1
	AuthorID int `json:"authorID"`
}
