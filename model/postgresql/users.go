package postgresql

import (
	"database/sql"
)

type UserModel struct {
	DB *sql.DB
}
