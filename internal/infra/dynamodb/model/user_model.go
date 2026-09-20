package dynamodb_model

type UserItem struct {
	PK    string
	SK    string

	Id    string
	Name  string
	Email string
}