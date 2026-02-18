package storage

import "context"

type Page struct {
	URL      string
	UserName string
}

type Storage interface {
	Save(ctx context.Context, p *Page) error
	PickRandom(ctx context.Context, username string) (*Page, error)
	Remove(ctx context.Context, p *Page) error
	IsExists(ctx context.Context, p *Page) (bool, error)
}