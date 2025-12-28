package transport

type Project struct {
	ID   int    `db:"id"`
	Name string `db:"name"`
	Code string `db:"code"`
}
