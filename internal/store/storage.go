package store

import (
	"context"
	"database/sql"
	"errors"
	"time"
)

var (
	ErrNotFound          = errors.New("Resource not found")
	ErrConflict          = errors.New("Resource already exists")
	ErrDublicateEmail    = errors.New("Email duplicate")
	ErrDublicateUsername = errors.New("Username duplicate")
	QueryTimeoutDuration = time.Second * 5
)

type Storage struct {
	Post interface {
		Create(context.Context, *Post) error
		GetByID(context.Context, int64) (*Post, error)
		Delete(context.Context, int64) error
		Update(context.Context, *Post) error
		GetUserFeed(ctx context.Context, userID int64, pg PaginatedFeedQuery) ([]PostWithMetaData, error)
	}

	User interface {
		Create(context.Context, *sql.Tx, *User) error
		GetByID(context.Context, int64) (*User, error)
		CreateAndInvite(ctx context.Context, user *User, token string, invitationExp time.Duration) error
	}

	Comments interface {
		Create(context.Context, *Comment) error

		GetByPostID(context.Context, int64) ([]Comment, error)
	}

	Followers interface {
		Follow(ctx context.Context, followerID, userID int64) error
		Unfollow(ctx context.Context, followerID, userID int64) error
	}
}

func NewStorage(db *sql.DB) Storage {
	return Storage{
		Post:      &PostStore{db: db},
		User:      &UserStore{db: db},
		Comments:  &CommentStore{db: db},
		Followers: &FollowerStore{db: db},
	}
}

func withTX(db *sql.DB, ctx context.Context, fn func(*sql.Tx) error) error {
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	if err := fn(tx); err != nil {
		_ = tx.Rollback()
		return err
	}
	return tx.Commit()
}
