package globals

// User will store user data
type User struct {
	ID        int
	FirstName string
	LastName  string
	UserName  string
	Password  string
}

var AllUsers []User
var nextID = 1

func UseNextID() int {
	id := nextID
	nextID++
	return id
}
