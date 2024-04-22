package car

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/jackc/pgconn"
	"hh.ru/pkg/db"
	"hh.ru/pkg/domain"
	"hh.ru/pkg/repo/car/interface"
	"log"
)

type carDatabase struct {
	db *sql.DB
}

func NewCarRepository(db *sql.DB) interfaces.CarRepository {
	return &carDatabase{
		db: db,
	}
}

func (r *carDatabase) Migrate(ctx context.Context) error {
	carQuery := `
    CREATE TABLE IF NOT EXISTS car(
       	id SERIAL PRIMARY KEY,
		regNum text not NULL,
		mark text not NULL,
		domain text not NULL,
		year int not NULL,
		owner int references people(id) not NULL
    );
    `
	_, err := r.db.ExecContext(ctx, carQuery)
	if err != nil {
		message := db.ErrMigrate.Error() + " car"
		log.Printf("%q: %s\n", message, err.Error())
		return db.ErrMigrate
	}

	return err
}

func (r *carDatabase) Create(ctx context.Context, car domain.Car) (*domain.Car, error) {
	var id int64
	err := r.db.QueryRowContext(ctx, "INSERT INTO car(regnum, mark, domain, year, owner) values($1, $2, $3, $4, $5) RETURNING id", car.RegNum, car.Mark, car.Model, car.Year, car.Owner).Scan(&id)
	if err != nil {
		var pgxError *pgconn.PgError
		if errors.As(err, &pgxError) {
			if pgxError.Code == "23505" {
				return nil, db.ErrDuplicate
			}
		}
		return nil, err
	}
	car.ID = id

	return &car, nil
}

func (r *carDatabase) All(ctx context.Context) ([]domain.Car, error) {
	rows, err := r.db.QueryContext(ctx, "SELECT * FROM car")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var all []domain.Car
	for rows.Next() {
		var car domain.Car
		if err := rows.Scan(&car.ID, &car.RegNum, &car.Mark, &car.Model, &car.Year, &car.Owner); err != nil {
			return nil, err
		}
		all = append(all, car)
	}
	return all, nil
}

func (r *carDatabase) GetByID(ctx context.Context, id int64) (*domain.Car, error) {
	row := r.db.QueryRowContext(ctx, "SELECT * FROM car WHERE id = $1", id)

	var car domain.Car
	if err := row.Scan(&car.ID, &car.RegNum, &car.Mark, &car.Model, &car.Year, &car.Owner); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, db.ErrNotExist
		}
		return nil, err
	}
	return &car, nil
}

func (r *carDatabase) Update(ctx context.Context, id int64, updatedCar domain.Car) (*domain.Car, error) {
	res, err := r.db.ExecContext(ctx, "UPDATE car SET RegNum = $1, Mark = $2, Model = $3, Year = $4, owner = o.id FROM people o WHERE car.id = $5 AND car.owner = o.id", updatedCar.RegNum, updatedCar.Mark, fmt.Sprint("%v", updatedCar.Year), updatedCar.Owner, id)
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

	return &updatedCar, nil
}

func (r *carDatabase) Delete(ctx context.Context, id int64) error {
	res, err := r.db.ExecContext(ctx, "DELETE FROM car WHERE id = $1", id)
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
