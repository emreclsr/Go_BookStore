package author

// Author author
//
// swagger:model Author
type Author struct {
	// The unique id of an author
	// Required: true
	// Minimum: 1
	ID int `json:"id"`

	// The name of an author
	// Required: true
	// Minimum length: 1
	Name string `json:"name"`
}
