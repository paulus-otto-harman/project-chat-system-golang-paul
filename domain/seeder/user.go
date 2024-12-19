package seeder

import (
	"homework/domain"
	"homework/helper"
)

func User() []domain.User {
	return []domain.User{
		{
			Email:    "super@mail.com",
			Password: helper.HashPassword("super"),
			Role:     domain.SuperAdmin,
			Profile: domain.Profile{
				FullName: "Super Admin",
				Phone:    "00",
				Salary:   150,
			},
		},
		{
			Email:    "admin@mail.com",
			Password: helper.HashPassword("admin"),
			Role:     domain.Admin,
			Profile: domain.Profile{
				FullName: "Admin Satu",
				Phone:    "01",
				Salary:   100,
			},
			Permissions: []domain.Permission{
				{ID: 1, Name: "Dashboard"},
				{ID: 2, Name: "Reports"},
				{ID: 6, Name: "Settings"},
			},
		},
		{
			Email:    "staff@mail.com",
			Password: helper.HashPassword("staff"),
			Role:     domain.Staff,
			Profile: domain.Profile{
				FullName: "Staff Satu",
				Phone:    "02",
				Salary:   50,
			},
			Permissions: []domain.Permission{
				{ID: 4, Name: "Orders"},
				{ID: 5, Name: "Customers"},
			},
		},
	}
}
