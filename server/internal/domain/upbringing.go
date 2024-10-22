package domain

import "fmt"

type Upbringing struct {
	Name        string
	Stat        string
	Description string
}

func (u *Upbringing) Value() interface{} {
	return u
}

func (u *Upbringing) String() string {
	return fmt.Sprintf("Name: %s\nStat: %s\nDescription: %s\n", u.Name, u.Stat, u.Description)
}
