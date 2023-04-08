package model
// Если ты хотела разделить структуры для разной бизнес-логики create/update и тд то нужно и называть структуру по разному,  а не просто User
type User struct {
	UserEntity
	ID string 
}
