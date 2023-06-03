package datastruct

type Task struct {
	ID     string `db:"id" json:"id"`
	Name   string `db:"name" json:"name"`
	Detail string `db:"detail" json:"detail"`
	Total  uint32 `db:"total" json:"-"`
}
