package accountmodel

// UserRole is exported
type UserRole int

const (
  Admin UserRole = iota
  Provider
  Consumer
)
