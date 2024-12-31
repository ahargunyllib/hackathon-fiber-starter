package fixture

import (
	"time"

	"github.com/ahargunyllib/hackathon-fiber-starter/domain/entity"
	"github.com/google/uuid"
)

var Rows = []string{
	"id", "name", "email", "password", "created_at", "updated_at", "deleted_at",
}

// * Make sure its the same order as the Struct
var (
	ActiveUser1 = entity.User{
		ID:        uuid.New(),
		Name:      "activeUser1",
		Email:     "activeUser1@gmail.com",
		Password:  "password",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		DeletedAt: nil,
	}
	ActiveUser2 = entity.User{
		ID:        uuid.New(),
		Name:      "activeUser2",
		Email:     "activeUser2@gmail.com",
		Password:  "password",
		CreatedAt: time.Now().Add(1 * time.Hour),
		UpdatedAt: time.Now().Add(1 * time.Hour),
		DeletedAt: nil,
	}
	InactiveUser1 = entity.User{
		ID:        uuid.New(),
		Name:      "inactiveUser1",
		Email:     "inactiveUser1@gmail.com",
		Password:  "password",
		CreatedAt: time.Now().Add(2 * time.Hour),
		UpdatedAt: time.Now().Add(2 * time.Hour),
		DeletedAt: func(t time.Time) *time.Time { return &t }(time.Now().Add(2 * time.Hour)),
	}
	InactiveUser2 = entity.User{
		ID:        uuid.New(),
		Name:      "inactiveUser2",
		Email:     "inactiveUser2@gmail.com",
		Password:  "password",
		CreatedAt: time.Now().Add(3 * time.Hour),
		UpdatedAt: time.Now().Add(3 * time.Hour),
		DeletedAt: func(t time.Time) *time.Time { return &t }(time.Now().Add(3 * time.Hour)),
	}
)
