package author

type authorService struct {
	repo AuthorRepository
}

func NewAuthorService(repo AuthorRepository) AuthorService {
	return authorService{repo: repo}
}

type AuthorService interface {
	GetAuthor(id int) (Author, error)
	GetAuthors() ([]Author, error)
	CreateAuthor(author Author) (int, error)
	UpdateAuthor(id int, author Author) error
	DeleteAuthor(id int) error
}

// Compile time proof of interface implementation
var _ AuthorService = authorService{}

func (s authorService) GetAuthor(id int) (Author, error) {
	return s.repo.Get(id)
}

func (s authorService) GetAuthors() ([]Author, error) {
	return s.repo.GetAll()
}

func (s authorService) CreateAuthor(author Author) (int, error) {
	return s.repo.Create(author)
}

func (s authorService) UpdateAuthor(id int, author Author) error {
	return s.repo.Update(id, author)
}

func (s authorService) DeleteAuthor(id int) error {
	return s.repo.Delete(id)
}
