package user

type User struct {
    ID       int64  // Auto-incremented typically by the database
    Username string
    Password string // Remember to hash before storing.
    Email    string
}
