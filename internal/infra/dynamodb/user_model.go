package dynamodb

type UserItem struct {
	PK    string
	SK    string

	Id    string
	Name  string
	Email string
}