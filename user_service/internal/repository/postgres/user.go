package postgres

import (
	"context"
	"fmt"

	"github.com/qara-qurt/telegram_plus/user_service/internal/model"

	"github.com/jackc/pgx"
	"github.com/jackc/pgx/v5/pgxpool"
)

type User struct {
	db *pgxpool.Pool
}

func NewUser(db *pgxpool.Pool) *User {
	return &User{
		db: db,
	}
}

func (u *User) Create(user model.UserInfo) (string, error) {
	var uuid string

	query := `INSERT INTO users(
				username,
				login,
				birthday_date, 
				email, 
				hash_password)
			VALUES ($1,$2,$3,$4,$5) RETURNING id`

	err := u.db.QueryRow(context.Background(), query,
		user.Username,
		user.Login,
		user.BirthDayDate,
		user.Email,
		user.Password).Scan(&uuid)

	if err != nil {
		return "", err
	}

	return uuid, nil
}

func (u *User) GetUser(uuid string) (model.User, error) {
	var user model.User

	query := `SELECT 
				id,
				username,
				login,
				birthday_date,
				email,
				hash_password,
				status,
				img,
				created_at,
				updated_at 
			FROM users WHERE id = $1`

	err := u.db.QueryRow(context.Background(), query, uuid).Scan(
		&user.ID,
		&user.Username,
		&user.Login,
		&user.BirthDayDate,
		&user.Email,
		&user.HashPassword,
		&user.Status,
		&user.Img,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	if err != nil {
		if err == pgx.ErrNoRows {
			return model.User{}, fmt.Errorf("no user found with ID: %s", uuid)
		}
		return model.User{}, err
	}

	return user, nil
}

func (u *User) GetUsers(limit int32, page int32) ([]model.User, error) {
	var users []model.User

	query := `SELECT
				id,
				username,
				login,
				birthday_date,
				email,
				hash_password,
				status,
				img,
				created_at,
				updated_at 
			FROM users LIMIT $1 OFFSET $2`

	offset := (page - 1) * limit
	rows, err := u.db.Query(context.Background(), query, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var user model.User
		if err := rows.Scan(
			&user.ID,
			&user.Username,
			&user.Login,
			&user.BirthDayDate,
			&user.Email,
			&user.HashPassword,
			&user.Status,
			&user.Img,
			&user.CreatedAt,
			&user.UpdatedAt); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func (u *User) GetUsersByCredentials(text string) ([]model.User, error) {
	var users []model.User

	query := `
        SELECT
            id,
            username,
            login,
            birthday_date,
            email,
            hash_password,
            status,
            img,
            created_at,
            updated_at
        FROM users
        WHERE username ILIKE '%' || $1 || '%' OR login ILIKE '%' || $1 || '%' OR email ILIKE '%' || $1 || '%'`

	arg := "%" + text + "%"
	rows, err := u.db.Query(context.Background(), query, arg)
	if err != nil {
		return nil, fmt.Errorf("error executing query: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var user model.User
		if err := rows.Scan(
			&user.ID,
			&user.Username,
			&user.Login,
			&user.BirthDayDate,
			&user.Email,
			&user.HashPassword,
			&user.Status,
			&user.Img,
			&user.CreatedAt,
			&user.UpdatedAt); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}
