package main

import (
	"context"
	"database/sql"
	"github.com/brianvoe/gofakeit"
	"github.com/jackc/pgx/v4"
	"log"
	"time"
)

const (
	dbDSN = "host=localhost port=54321 dbname=auth user=auth-user password=auth-password sslmode=disable"
)

func main() {
	ctx := context.Background()

	// Создаем соединение с базой данных
	con, err := pgx.Connect(ctx, dbDSN)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	defer func(con *pgx.Conn, ctx context.Context) {
		_ = con.Close(ctx)
	}(con, ctx)

	password := gofakeit.Password(true, true, true, true, false, 12)
	passwordConfirm := password

	// Делаем запрос на вставку записи в таблицу user
	res, err := con.Exec(ctx, "INSERT INTO user_ent (name, email, password, password_confirm, role) VALUES ($1, $2, $3, $4, $5)", gofakeit.Name(), gofakeit.Email(), password, passwordConfirm, gofakeit.Number(0, 1))
	if err != nil {
		log.Fatalf("failed to insert user: %v", err)
	}

	log.Printf("inserted %d rows", res.RowsAffected())

	// Делаем запрос на выборку записей из таблицы user
	rows, err := con.Query(ctx, "SELECT id, name, email, role, created_at, updated_at FROM user_ent")
	if err != nil {
		log.Fatalf("failed to select users: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var name, email string
		var role int
		var createdAt time.Time
		var updatedAt sql.NullTime

		err = rows.Scan(&id, &name, &email, &role, &createdAt, &updatedAt)
		if err != nil {
			log.Fatalf("failed to scan note: %v", err)
		}

		log.Printf("id: %d, name: %s, email: %s, role: %d, created_at: %v, updated_at: %v\n", id, name, email, role, createdAt, updatedAt)
	}
}
