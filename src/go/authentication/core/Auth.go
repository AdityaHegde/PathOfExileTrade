package authcore

import (
  "github.com/AdityaHegde/PathOfExileTrade/account"
)

// Auth is exported
type Auth interface {
  Init() error
  Generate(user *account.User) (string, error)
  Validate(payload string) (string, error)
}
