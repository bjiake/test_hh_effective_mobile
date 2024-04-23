package people

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/jackc/pgconn"
	"hh.ru/pkg/api/filter"
	"hh.ru/pkg/db"
	"hh.ru/pkg/domain"
	"hh.ru/pkg/repo/people/interface"
	"log"
	"strings"
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

func (r *peopleDatabase) Get(ctx context.Context, filter *filter.People) ([]domain.People, error) {
	query := "SELECT * FROM people"
	var args []interface{}

	if filter != nil {
		var whereClauses []string
		if filter.ID != nil {
			args = append(args, *filter.ID)
			whereClauses = append(whereClauses, fmt.Sprintf("id = $%d", len(args)))
		}
		if filter.Name != nil {
			args = append(args, *filter.Name)
			whereClauses = append(whereClauses, fmt.Sprintf("name = $%d", len(args)))
		}
		if filter.SurName != nil {
			args = append(args, *filter.SurName)
			whereClauses = append(whereClauses, fmt.Sprintf("surname = $%d", len(args)))
		}
		if filter.Patronymic != nil {
			args = append(args, *filter.Patronymic)
			whereClauses = append(whereClauses, fmt.Sprintf("patronymic = $%d", len(args)))
		}

		if len(whereClauses) > 0 {
			query += " WHERE " + strings.Join(whereClauses, " AND ")
		}
	}

	rows, err := r.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var peoples []domain.People
	for rows.Next() {
		var people domain.People
		if err := rows.Scan(&people.ID, &people.Name, &people.SurName, &people.Patronymic); err != nil {
			return nil, err
		}
		peoples = append(peoples, people)
	}
	return peoples, nil
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
