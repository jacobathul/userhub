package user

import "errors"

var users = []User{} // Temporary in-memory storage

type UserService struct{}

func (s *UserService) CreateUser(user User) (*User, error) {
    // For this example, we'll just append to our in-memory slice.
    // In real-world scenarios, this is where you'd interact with a database.

    // Basic validation example
    if user.Username == "" || user.Password == "" {
        return nil, errors.New("invalid input")
    }

    user.ID = int64(len(users) + 1) // Mock ID assignment
    users = append(users, user)
    return &user, nil
}