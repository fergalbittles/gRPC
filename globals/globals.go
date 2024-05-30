package globals

import (
	"sync"
	"sync/atomic"
)

// User will store user data
type User struct {
	ID        int64
	FirstName string
	LastName  string
	UserName  string
	Password  string
}

var (
	mu    sync.RWMutex
	id    atomic.Int64
	users []User = make([]User, 0)
)

func useNextID() int64 {
	return id.Add(1)
}

func AppendUsers(user ...User) {
	mu.Lock()
	defer mu.Unlock()
	for _, u := range users {
		if u.ID == 0 {
			u.ID = useNextID()
		}
	}
	users = append(users, user...)
}

func ListUsers() []User {
	mu.RLock()
	defer mu.RUnlock()
	uu := make([]User, len(users))
	copy(uu, users)
	return uu
}
