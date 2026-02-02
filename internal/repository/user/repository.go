package user

import (
	"context"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/mrlexus21/auth/internal/model"
	"github.com/mrlexus21/auth/internal/repository"
	"github.com/mrlexus21/auth/internal/repository/user/converter"
	modelRepo "github.com/mrlexus21/auth/internal/repository/user/model"
)

const (
	tableName             = "users"
	idColumn              = "id"
	nameColumn            = "name"
	emailColumn           = "email"
	passwordColumn        = "password"
	passwordConfirmColumn = "password_confirm"
	roleColumn            = "role"
	createdAtColumn       = "created_at"
	updatedAtColumn       = "updated_at"
)

type repo struct {
	db *pgxpool.Pool
}

func NewRepository(db *pgxpool.Pool) repository.UserRepository {
	return &repo{db: db}
}

func (r *repo) Get(ctx context.Context, id int64) (*model.User, error) {
	query, args, err := sq.Select(idColumn, nameColumn, emailColumn, roleColumn, createdAtColumn, updatedAtColumn).
		From(tableName).
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{idColumn: id}).
		ToSql()

	if err != nil {
		return nil, err
	}

	var user modelRepo.User
	if err := r.db.QueryRow(ctx, query, args...).Scan(&user.ID, &user.Info.Name, &user.Info.Email, &user.Info.Role, &user.CreatedAt, &user.UpdatedAt); err != nil {
		return nil, err
	}

	return converter.ToUserFromRepo(&user), nil
}

func (r *repo) Create(ctx context.Context, user *model.CreateUser) (int64, error) {
	query, args, err := sq.Insert(tableName).
		PlaceholderFormat(sq.Dollar).
		Columns(nameColumn, emailColumn, passwordColumn, passwordConfirmColumn, roleColumn).
		Values(user.Info.Name, user.Info.Email, user.Credentials.Password, user.Credentials.PasswordConfirm, user.Info.Role).
		Suffix("RETURNING id").
		ToSql()
	if err != nil {
		return 0, err
	}

	var userID int64
	if err := r.db.QueryRow(ctx, query, args...).Scan(&userID); err != nil {
		return 0, err
	}

	return userID, nil
}

func (r *repo) Update(ctx context.Context, info *model.UpdateUserInfo) error {
	builderUpdate := sq.Update(tableName).
		PlaceholderFormat(sq.Dollar).
		Set(updatedAtColumn, time.Now())

	if name := info.Name; name != nil && *name != "" {
		builderUpdate = builderUpdate.Set(nameColumn, *name)
	}
	if email := info.Email; email != nil && *email != "" {
		builderUpdate = builderUpdate.Set(emailColumn, *email)
	}

	builderUpdate = builderUpdate.Where(sq.Eq{idColumn: info.ID})

	query, args, err := builderUpdate.ToSql()
	if err != nil {
		return err
	}

	_, err = r.db.Exec(ctx, query, args...)

	return err
}

func (r *repo) Delete(ctx context.Context, id int64) error {
	query, args, err := sq.Delete(tableName).
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{idColumn: id}).
		ToSql()
	if err != nil {
		return err
	}

	_, err = r.db.Exec(ctx, query, args...)

	return err
}
