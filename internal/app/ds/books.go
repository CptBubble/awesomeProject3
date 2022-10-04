package ds

type UsersBuy struct {
	ID   uint `gorm:"primarykey"`
	Code uint
	Name string
	Age  int
	Book string
}
