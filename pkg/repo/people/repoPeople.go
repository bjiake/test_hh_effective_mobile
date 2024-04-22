package people

import (
	"context"
	"database/sql"
	"errors"
	"github.com/jackc/pgconn"
	"hh.ru/pkg/db"
	"hh.ru/pkg/domain"
	"hh.ru/pkg/repo/people/interface"
	"log"
)

type peopleDatabase struct {
	db *sql.DB
}

func NewPeopleRepository(db *sql.DB) interfaces.PeopleRepository {
	return &peopleDatabase{
		db: db,
	}
}

func (r *peopleDatabase) Migrate(ctx context.Context) error {
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
		message := db.ErrMigrate.Error() + " people"
		log.Printf("%q: %s\n", message, err.Error())
		return db.ErrMigrate
	}

	return err
}

func (r *peopleDatabase) Create(ctx context.Context, people domain.People) (*domain.People, error) {
	var id int64
	err := r.db.QueryRowContext(ctx, "INSERT INTO people(name, SurName, patronymic) values($1, $2, $3) RETURNING id", people.Name, people.SurName, people.Patronymic).Scan(&id)
	if err != nil {
		var pgxError *pgconn.PgError
		if errors.As(err, &pgxError) {
			if pgxError.Code == "23505" {
				return nil, db.ErrDuplicate
			}
		}
		return nil, err
	}
	people.ID = id

	return &people, nil
}

func (r *peopleDatabase) All(ctx context.Context) ([]domain.People, error) {
	rows, err := r.db.QueryContext(ctx, "SELECT * FROM people")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var all []domain.People
	for rows.Next() {
		var people domain.People
		if err := rows.Scan(&people.ID, &people.Name, &people.SurName, &people.Patronymic); err != nil {
			return nil, err
		}
		all = append(all, people)
	}
	return all, nil
}

func (r *peopleDatabase) GetByID(ctx context.Context, id int64) (*domain.People, error) {
	row := r.db.QueryRowContext(ctx, "SELECT * FROM people WHERE id = $1", id)

	var people domain.People
	if err := row.Scan(&people.ID, &people.Name, &people.SurName, &people.Patronymic); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, db.ErrNotExist
		}
		return nil, err
	}
	return &people, nil
}

func (r *peopleDatabase) Update(ctx context.Context, id int64, updatedpeople domain.People) (*domain.People, error) {
	res, err := r.db.ExecContext(ctx, "UPDATE people SET Name = $1, SurName = $2, Patronymic = $3 FROM people o WHERE people.id = $4", updatedpeople.Name, updatedpeople.SurName, updatedpeople.Patronymic, id)
	if err != nil {
		var pgxError *pgconn.PgError
		if errors.As(err, &pgxError) {
			if pgxError.Code == "23505" {
				return nil, db.ErrDuplicate
			}
		}
		return nil, err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return nil, err
	}

	if rowsAffected == 0 {
		return nil, db.ErrUpdateFailed
	}

	return &updatedpeople, nil
}

func (r *peopleDatabase) Delete(ctx context.Context, id int64) error {
	res, err := r.db.ExecContext(ctx, "DELETE FROM people WHERE id = $1", id)
	if err != nil {
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return db.ErrDeleteFailed
	}

	return err
}
