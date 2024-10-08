package todo

type TodoList struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"username"`
}
type UserList struct {
	Id     int
	UserId string
	ListId string
}
type TodoItem struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Done        string `json:"done"`
}
type ListItem struct {
	Id     int
	ListId string
	ItemId string
}