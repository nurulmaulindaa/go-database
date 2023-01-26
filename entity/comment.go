package entity

import "github.com/guregu/null"

type Comment struct {
	Id      int32
	Email   string
	Comment null.String
}
