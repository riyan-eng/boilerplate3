package datastruct

type Task struct {
	ID     string `db:"id"`
	Name   string `db:"name"`
	Detail string `db:"detail"`
}
