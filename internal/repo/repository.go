package repo

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/Alina9496/documents/internal/domain"
	"github.com/Alina9496/documents/internal/service/dto"
	"github.com/Alina9496/tool/pkg/logger"
	"github.com/Alina9496/tool/pkg/postgres"
	"github.com/Masterminds/squirrel"
	"github.com/google/uuid"

	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
)

type Repository struct {
	pg *postgres.Postgres
	l  *logger.Logger
}

func New(pg *postgres.Postgres, l *logger.Logger) *Repository {
	return &Repository{
		pg: pg,
		l:  l,
	}
}

type db interface {
	Exec(ctx context.Context, sql string, arguments ...interface{}) (commandTag pgconn.CommandTag, err error)
	Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row
}

func (r *Repository) conn(ctx context.Context) db {
	if tx, ok := ctx.Value(tansactionKey).(pgx.Tx); ok {
		return tx
	}
	return r.pg.Pool
}

func (r *Repository) ExecTx(ctx context.Context, fn func(ctx context.Context) error) error {
	if _, ok := ctx.Value(tansactionKey).(pgx.Tx); ok {
		return fn(ctx)
	}

	tx, err := r.pg.Pool.Begin(ctx)
	if err != nil {
		return err
	}

	ctx = context.WithValue(ctx, tansactionKey, tx)

	defer func() {
		if p := recover(); p != nil {
			if errRollback := tx.Rollback(ctx); errRollback != nil {
				r.l.Error("rollback err %s", errRollback)
			}
			err = fmt.Errorf("panic :%s", p)
			return
		}
		if errCommit := tx.Commit(ctx); errCommit != nil {
			r.l.Error("commit err %s", errCommit)
		}
	}()
	return fn(ctx)
}

func (r *Repository) Registration(ctx context.Context, user *domain.User) error {
	query, args, err := r.pg.Builder.
		Insert(tableUser).
		Columns(
			"login",
			"password",
			"created_at",
		).
		Values(
			user.Login,
			user.Password,
			time.Now(),
		).
		Suffix(suffixReturningID).
		ToSql()
	if err != nil {
		return fmt.Errorf("error build query: %w", err)
	}

	_, err = r.conn(ctx).Exec(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("error registration user: %w", err)
	}

	return nil
}

func (r *Repository) CheckUser(ctx context.Context, user *domain.User) (uuid.UUID, error) {
	query, args, err := r.pg.Builder.
		Select("id").
		From(tableUser).
		Where(squirrel.Eq{"login": user.Login}).
		Where(squirrel.Eq{"password": user.Password}).
		ToSql()
	if err != nil {
		return uuid.Nil, fmt.Errorf("error build query: %w", err)
	}

	var id uuid.UUID
	err = r.conn(ctx).QueryRow(ctx, query, args...).Scan(&id)
	if err != nil {
		return uuid.Nil, fmt.Errorf("error check user: %w", err)
	}

	return id, nil
}

func (r *Repository) Authentication(ctx context.Context, user *domain.User) error {
	query, args, err := r.pg.Builder.
		Insert(tableToken).
		Columns(
			"user_id",
			"token",
			"created_at",
		).
		Values(
			user.ID,
			user.Token,
			time.Now(),
		).
		Suffix(suffixReturningID).
		ToSql()
	if err != nil {
		return fmt.Errorf("error build query: %w", err)
	}

	var id uuid.UUID
	err = r.conn(ctx).QueryRow(ctx, query, args...).Scan(&id)
	if err != nil {
		return fmt.Errorf("error authentication user: %w", err)
	}

	return nil
}

func (r *Repository) LogOut(ctx context.Context, token string) error {
	query, args, err := r.pg.Builder.
		Delete(tableToken).
		Where(squirrel.Eq{"token": token}).
		ToSql()
	if err != nil {
		return fmt.Errorf("error build query: %w", err)
	}

	commandTag, err := r.conn(ctx).Exec(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("error delete token: %w", err)
	}
	if commandTag.RowsAffected() == 0 {
		return ErrTokenNotFound
	}
	return nil
}

func (r *Repository) GetUserID(ctx context.Context, token string) (uuid.UUID, error) {
	query, args, err := r.pg.Builder.
		Select("user_id").
		From(tableToken).
		Where(squirrel.Eq{"token": token}).
		ToSql()
	if err != nil {
		return uuid.Nil, fmt.Errorf("error build query: %w", err)
	}

	var userID uuid.UUID
	err = r.conn(ctx).QueryRow(ctx, query, args...).Scan(&userID)
	if err != nil {
		return uuid.Nil, fmt.Errorf("error check user: %w", err)
	}

	return userID, nil
}

func (r *Repository) Save(ctx context.Context, document *domain.Document) (uuid.UUID, error) {
	sql, args, err := r.pg.Builder.Insert(tableDocument).SetMap(map[string]any{
		"name":       document.Name,
		"file":       document.Content,
		"mime":       document.Mime,
		"is_public":  document.Public,
		"user_id":    document.UserID,
		"created_at": time.Now(),
	}).Suffix(suffixReturningID).ToSql()
	if err != nil {
		return uuid.Nil, fmt.Errorf("error build query: %w", err)
	}

	var id uuid.UUID
	err = r.conn(ctx).QueryRow(ctx, sql, args...).Scan(&id)
	if err != nil {
		return uuid.Nil, fmt.Errorf("error save document: %w", err)
	}

	return id, nil
}

func (r *Repository) AddGrant(ctx context.Context, grant *domain.Grant) error {
	sql, args, err := r.pg.Builder.Insert(tableGrant).
		SetMap(map[string]any{
			"user_id":          grant.UserID,
			"document_id":      grant.DocumentID,
			"grant_user_login": grant.GrantUserLogin,
			"created_at":       time.Now(),
		}).
		ToSql()
	if err != nil {
		return fmt.Errorf("error build query: %w", err)
	}

	_, err = r.conn(ctx).Exec(ctx, sql, args...)
	if err != nil {
		return fmt.Errorf("error add grant: %w", err)
	}

	return nil
}

func (r *Repository) GetDocument(ctx context.Context, id uuid.UUID) (*domain.Document, error) {
	sql, args, err := r.pg.Builder.Select(
		"file",
		"mime",
		"is_public",
		"user_id",
	).From(tableDocument).Where(
		squirrel.Eq{"id": id},
	).
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("error build query: %w", err)
	}

	var document domain.Document

	err = r.conn(ctx).QueryRow(ctx, sql, args...).Scan(&document.Content, &document.Mime, &document.Public, &document.UserID)
	if err != nil {
		return nil, fmt.Errorf("error add grant: %w", err)
	}

	return &document, nil
}

func (r *Repository) CheckGrant(ctx context.Context, documentID uuid.UUID, login string) (bool, error) {
	sql, args, err := r.pg.Builder.Select("1").From(tableGrant).
		Where(squirrel.Eq{"grant_user_login": login}).
		Where(squirrel.Eq{"document_id": documentID}).
		ToSql()
	if err != nil {
		return false, fmt.Errorf("error build query: %w", err)
	}

	var exist bool
	err = r.conn(ctx).QueryRow(ctx, sql, args...).Scan(&exist)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return false, nil
		}
		return false, fmt.Errorf("error build query: %w", err)
	}

	return exist, nil
}

func (r *Repository) GetUser(ctx context.Context, id uuid.UUID) (*domain.User, error) {
	sql, args, err := r.pg.Builder.Select(
		"login",
	).From(tableUser).
		Where(squirrel.Eq{"id": id}).
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("error build query: %w", err)
	}

	var user domain.User
	err = r.conn(ctx).QueryRow(ctx, sql, args...).Scan(&user.Login)
	if err != nil {
		return nil, fmt.Errorf("error exec query: %w", err)
	}

	return &user, nil
}

func (r *Repository) GetDocuments(ctx context.Context, filter *dto.GetDocuments) ([]domain.Document, error) {
	where := make(squirrel.Or, 0, 3)

	where = append(where, squirrel.Eq{filter.Key: filter.Value})
	where = append(where, squirrel.Eq{"d.user_id": filter.UserID})
	where = append(where, squirrel.Eq{"g.grant_user_login": filter.Login})

	query, args, err := r.pg.Builder.Select(
		"d.id",
		"d.name",
		"d.mime",
		"d.is_public",
		"d.created_at",
		"array_agg(DISTINCT g.grant_user_login) AS grant_user_logins",
	).From("public.document AS d").
		Join("public.grants AS g ON d.user_id = g.user_id").
		Where(where).
		GroupBy("d.id", "d.name", "d.mime", "d.is_public").
		Limit(uint64(filter.Limit)).
		ToSql()

	if err != nil {
		return nil, fmt.Errorf("error building query: %w", err)
	}

	documents := make([]domain.Document, 0, filter.Limit)
	rows, err := r.conn(ctx).Query(ctx, query, args...)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var doc domain.Document
		var grantLogins sql.NullString

		err := rows.Scan(&doc.ID, &doc.Name, &doc.Mime, &doc.Public, &doc.CreatedAt, &grantLogins)
		if err != nil {
			return nil, err
		}

		if grantLogins.Valid {
			doc.Grant = strings.Split(strings.Trim(grantLogins.String, "{}"), ",")
		}

		documents = append(documents, doc)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return documents, nil
}

func (r *Repository) DeleteDocument(ctx context.Context, id, userID uuid.UUID) (uuid.UUID, error) {
	sql, args, err := r.pg.Builder.Delete(tableDocument).
		Where(squirrel.Eq{"id": id}).
		Where(squirrel.Eq{"user_id": userID}).
		ToSql()
	if err != nil {
		return uuid.Nil, fmt.Errorf("error building query: %w", err)
	}
	_, err = r.conn(ctx).Exec(ctx, sql, args...)
	if err != nil {
		return uuid.Nil, err
	}

	return id, nil
}
