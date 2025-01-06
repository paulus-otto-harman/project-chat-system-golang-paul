package seeder

import (
	"homework/domain"
	"homework/helper"
)

func User() []domain.User {
	return []domain.User{
		{
			Name:     "User Satu",
			Email:    "user1@mail.com",
			Password: helper.HashPassword("user1"),
		},
		{
			Name:     "User Dua",
			Email:    "user2@mail.com",
			Password: helper.HashPassword("user2"),
		},
		{
			Name:     "User Tiga",
			Email:    "user3@mail.com",
			Password: helper.HashPassword("user3"),
		},
		{
			Name:     "User Empat",
			Email:    "user4@mail.com",
			Password: helper.HashPassword("user4"),
		},
		{
			Name:     "User Lima",
			Email:    "user5@mail.com",
			Password: helper.HashPassword("user5"),
		},
		{
			Name:     "User Enam",
			Email:    "user6@mail.com",
			Password: helper.HashPassword("user6"),
		},
	}
}
