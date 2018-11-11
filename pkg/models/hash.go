package models

type Hash interface {
	Generate(s string) (string, error)
	Compare(hash string, s string) error
}