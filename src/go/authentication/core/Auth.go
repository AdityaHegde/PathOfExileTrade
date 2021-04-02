package authcore

import accountmodel "github.com/AdityaHegde/PathOfExileTrade/model/account"

type Auth interface {
  Init() error
  Generate(user *accountmodel.User) (string, error)
  Validate(payload string) (string, error)
}
