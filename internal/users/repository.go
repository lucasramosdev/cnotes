package users

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository interface {
	GetUserByEmail(ctx context.Context, email *string) (*User, error)
}

type RepositoryPostgres struct {
	Conn *pgxpool.Pool
}

func (r *RepositoryPostgres) GetUserByEmail(ctx context.Context, email *string) (*User, error) {
	var user User
	err := r.Conn.QueryRow(
		ctx,
		`SELECT id, email, password from users where email = $1;`,
		*email,
	).Scan(&user.ID, &user.Email, &user.Password)
	log.Println(err)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
