package postgresql

import (
	// "errors"
	"database/sql"
	// "tadeobennett/celebrateease/pkg/models"
)

type UserModel struct {
	DB *sql.DB
}