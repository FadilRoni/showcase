package query

const (
	AddBook = `INSERT INTO books (title, author, description)
	VALUES ($1, $2, $3)
	Returning *`
)

const (
	GetAllBook = `SELECT * FROM books`
)

const (
	GetBookId = `SELECT * FROM books WHERE id = $1`
)

const (
	UpdateBook = `UPDATE * FROM books WHERE id = $1`
)

const (
	DeleteBook = `DELETE FROM books WHERE id = $1 RETURNING *`
)
