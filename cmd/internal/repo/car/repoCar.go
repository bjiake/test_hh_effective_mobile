package car

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/jackc/pgconn"
	"hh.ru/cmd/internal/model"
	"hh.ru/cmd/internal/repo"
	"log"
)

type PgSQLCarRepository struct {
	db *sql.DB
}

func NewPgSqlCarRepository(db *sql.DB) *PgSQLCarRepository {
	return &PgSQLCarRepository{
		db: db,
	}
}

type RepositoryCar interface {
	Migrate(ctx context.Context) error
	Create(ctx context.Context, car model.Car) (*model.Car, error)
	All(ctx context.Context) ([]model.Car, error)
	GetByID(ctx context.Context, id int64) (*model.Car, error)
	Update(ctx context.Context, id int64, updatedCar model.Car) (*model.Car, error)
	Delete(ctx context.Context, id int64) error
}

func (r *PgSQLCarRepository) Migrate(ctx context.Context) error {
	carQuery := `
    CREATE TABLE IF NOT EXISTS car(
       	id SERIAL PRIMARY KEY,
		regNum text not NULL,
		mark text not NULL,
		model text not NULL,
		year int not NULL,
		owner int references people(id) not NULL
    );
    `
	_, err := r.db.ExecContext(ctx, carQuery)
	if err != nil {
		message := repo.ErrMigrate.Error() + " car"
		log.Printf("%q: %s\n", message, err.Error())
		return repo.ErrMigrate
	}

	return err
}

func (r *PgSQLCarRepository) Create(ctx context.Context, car model.Car) (*model.Car, error) {
	var id int64
	err := r.db.QueryRowContext(ctx, "INSERT INTO car(regnum, mark, model, year, owner) values($1, $2, $3, $4, $5) RETURNING id", car.RegNum, car.Mark, car.Model, car.Year, car.Owner).Scan(&id)
	if err != nil {
		var pgxError *pgconn.PgError
		if errors.As(err, &pgxError) {
			if pgxError.Code == "23505" {
				return nil, repo.ErrDuplicate
			}
		}
		return nil, err
	}
	car.ID = id

	return &car, nil
}

func (r *PgSQLCarRepository) All(ctx context.Context) ([]model.Car, error) {
	rows, err := r.db.QueryContext(ctx, "SELECT * FROM car")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var all []model.Car
	for rows.Next() {
		var car model.Car
		if err := rows.Scan(&car.ID, &car.RegNum, &car.Mark, &car.Model, &car.Year, &car.Owner); err != nil {
			return nil, err
		}
		all = append(all, car)
	}
	return all, nil
}

func (r *PgSQLCarRepository) GetByID(ctx context.Context, id int64) (*model.Car, error) {
	row := r.db.QueryRowContext(ctx, "SELECT * FROM car WHERE id = $1", id)

	var car model.Car
	if err := row.Scan(&car.ID, &car.RegNum, &car.Mark, &car.Model, &car.Year, &car.Owner); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, repo.ErrNotExist
		}
		return nil, err
	}
	return &car, nil
}

func (r *PgSQLCarRepository) Update(ctx context.Context, id int64, updatedCar model.Car) (*model.Car, error) {
	res, err := r.db.ExecContext(ctx, "UPDATE car SET RegNum = $1, Mark = $2, Model = $3, Year = $4, owner = o.id FROM people o WHERE car.id = $5 AND car.owner = o.id", updatedCar.RegNum, updatedCar.Mark, fmt.Sprint("%v", updatedCar.Year), updatedCar.Owner, id)
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

	return &updatedCar, nil
}

func (r *PgSQLCarRepository) Delete(ctx context.Context, id int64) error {
	res, err := r.db.ExecContext(ctx, "DELETE FROM car WHERE id = $1", id)
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
