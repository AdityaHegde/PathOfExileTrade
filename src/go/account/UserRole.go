package account

// UserRole is exported
type UserRole int

const (
  Admin UserRole = iota
  Owner = 5
  Moderator = 15
  Provider = 25
  Consumer = 50
)
