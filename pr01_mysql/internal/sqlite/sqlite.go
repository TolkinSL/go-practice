package sqlite

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"

	"github.com/TolkinSL/go-practice/pr01_mysql/internal/storage"
)

type SQLite struct {
	db *sql.DB
}

func New(patch string) (*SQLite, error) {
	db, err := sql.Open("sqlite3", patch)
	if err != nil {
		return nil, fmt.Errorf("can't open db file: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("can't connect to database: %w", err)
	}

	return &SQLite{db: db}, nil
}

// Save Сохраняет страницу в хранилище
func (s *SQLite) Save(ctx context.Context, p *storage.Page) error {
	q := `INSERT INTO pages (url, user_name) VALUES (?, ?)`

	if _, err := s.db.ExecContext(ctx, q, p.URL, p.UserName); err != nil {
		return fmt.Errorf("can't save page: %w", err)
	}

	return nil
}

// IsExist Проверяет присутствует ли страница в хранилище.
func (s *SQLite) IsExists(ctx context.Context, p *storage.Page) (bool, error) {
	q := `SELECT COUNT(*) FROM pages WHERE url = ? AND user_name = ?`

	var count int
	if err := s.db.QueryRowContext(ctx, q, p.URL, p.UserName).Scan(&count); err != nil {
		return false, fmt.Errorf("can't check page exist: %w", err)
	}

	return count > 0, nil
}

func (s *SQLite) Init(ctx context.Context) error {
	q := `CREATE TABLE IF NOT EXISTS pages (url TEXT, user_name TEXT)`

	_, err := s.db.ExecContext(ctx, q)
	if err != nil {
		return fmt.Errorf("can't create table: %w", err)
	}

	return nil
}
