package transport

type Project struct {
	ID   string `db:"id"`
	Name string `db:"name"`
	Code string `db:"code"`
}
