package people

import (
	"context"
	"database/sql"
	"errors"
	"github.com/jackc/pgconn"
	"hh.ru/cmd/internal/model"
	"hh.ru/cmd/internal/repo"
	"log"
)

type PgSQLPeopleRepository struct {
	db *sql.DB
}

func NewPgSqlPeopleRepository(db *sql.DB) *PgSQLPeopleRepository {
	return &PgSQLPeopleRepository{
		db: db,
	}
}

type RepositoryPeople interface {
	Migrate(ctx context.Context) error
	Create(ctx context.Context, people model.People) (*model.People, error)
	All(ctx context.Context) ([]model.People, error)
	GetByID(ctx context.Context, id int64) (*model.People, error)
	Update(ctx context.Context, id int64, updatedPeo model.People) (*model.People, error)
	Delete(ctx context.Context, id int64) error
}

func (r *PgSQLPeopleRepository) Migrate(ctx context.Context) error {
	peopleQuery := `
    CREATE TABLE IF NOT EXISTS people(
		id SERIAL PRIMARY KEY,
		name text not NULL,
		SurName text not NULL,
		patronymic text not NULL
    );
    `

	_, err := r.db.ExecContext(ctx, peopleQuery)
	if err != nil {
		message := repo.ErrMigrate.Error() + " people"
		log.Printf("%q: %s\n", message, err.Error())
		return repo.ErrMigrate
	}

	return err
}

func (r *PgSQLPeopleRepository) Create(ctx context.Context, people model.People) (*model.People, error) {
	var id int64
	err := r.db.QueryRowContext(ctx, "INSERT INTO people(name, SurName, patronymic) values($1, $2, $3) RETURNING id", people.Name, people.SurName, people.Patronymic).Scan(&id)
	if err != nil {
		var pgxError *pgconn.PgError
		if errors.As(err, &pgxError) {
			if pgxError.Code == "23505" {
				return nil, repo.ErrDuplicate
			}
		}
		return nil, err
	}
	people.ID = id

	return &people, nil
}

func (r *PgSQLPeopleRepository) All(ctx context.Context) ([]model.People, error) {
	rows, err := r.db.QueryContext(ctx, "SELECT * FROM people")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var all []model.People
	for rows.Next() {
		var people model.People
		if err := rows.Scan(&people.ID, &people.Name, &people.SurName, &people.Patronymic); err != nil {
			return nil, err
		}
		all = append(all, people)
	}
	return all, nil
}

func (r *PgSQLPeopleRepository) GetByID(ctx context.Context, id int64) (*model.People, error) {
	row := r.db.QueryRowContext(ctx, "SELECT * FROM people WHERE id = $1", id)

	var people model.People
	if err := row.Scan(&people.ID, &people.Name, &people.SurName, &people.Patronymic); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, repo.ErrNotExist
		}
		return nil, err
	}
	return &people, nil
}

func (r *PgSQLPeopleRepository) Update(ctx context.Context, id int64, updatedpeople model.People) (*model.People, error) {
	res, err := r.db.ExecContext(ctx, "UPDATE people SET Name = $1, SurName = $2, Patronymic = $3 FROM people o WHERE people.id = $4", updatedpeople.Name, updatedpeople.SurName, updatedpeople.Patronymic, id)
	if err != nil {
		var pgxError *pgconn.PgError
		if errors.As(err, &pgxError) {
			if pgxError.Code == "23505" {
				return nil, repo.ErrDuplicate
			}
		}
		return nil, err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return nil, err
	}

	if rowsAffected == 0 {
		return nil, repo.ErrUpdateFailed
	}

	return &updatedpeople, nil
}

func (r *PgSQLPeopleRepository) Delete(ctx context.Context, id int64) error {
	res, err := r.db.ExecContext(ctx, "DELETE FROM people WHERE id = $1", id)
	if err != nil {
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return repo.ErrDeleteFailed
	}

	return err
}
