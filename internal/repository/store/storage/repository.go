package storage

import (
	"gorm.io/gorm"
	"time"
)

type (
	Repository struct {
		db *gorm.DB

		options Options
	}

	Options struct {
		Timeout      time.Duration
		DefaultLimit uint64
	}
)

func New(db *gorm.DB, options Options) *Repository {
	repo := &Repository{
		db: db,
	}

	repo.SetOptions(options)

	return repo
}

func (r *Repository) SetOptions(options Options) {
	if options.Timeout == 0 {
		options.Timeout = time.Second * 30
	}

	if options.DefaultLimit == 0 {
		options.DefaultLimit = 15
	}

	if r.options != options {
		r.options = options
	}
}
