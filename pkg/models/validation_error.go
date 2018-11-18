package models

type ValidationErr struct {
	Tag string `json:"tag"`
	Value interface{} `json:"value"`
}