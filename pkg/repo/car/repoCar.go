package car

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/jackc/pgconn"
	"hh.ru/pkg/api/filter"
	"hh.ru/pkg/db"
	"hh.ru/pkg/domain"
	"hh.ru/pkg/repo/car/interface"
	"log"
	"strings"
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

func (r *carDatabase) Create(ctx context.Context, car domain.Car) (*domain.RequestCar, error) {
	var id int64
	owner, err := r.getOwner(ctx, car.Owner)
	if err != nil {
		return nil, fmt.Errorf("create car: %w", err)
	}

	err = r.db.QueryRowContext(ctx, "INSERT INTO car(regnum, mark, model, year, owner) values($1, $2, $3, $4, $5) RETURNING id", car.RegNum, car.Mark, car.Model, car.Year, car.Owner).Scan(&id)
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

	requestCar := &domain.RequestCar{
		ID:     car.ID,
		RegNum: car.RegNum,
		Mark:   car.Mark,
		Model:  car.Model,
		Year:   car.Year,
		Owner:  owner,
	}

	return requestCar, nil
}

func (r *carDatabase) GetByRegNum(ctx context.Context, regNum string) (*domain.RequestCar, error) {
	row := r.db.QueryRowContext(ctx, "SELECT * FROM car WHERE regnum = $1", regNum)

	var car domain.Car
	if err := row.Scan(&car.ID, &car.RegNum, &car.Mark, &car.Model, &car.Year, &car.Owner); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, db.ErrNotExist
		}
		return nil, err
	}
	owner, err := r.getOwner(ctx, car.Owner)
	if err != nil {
		return nil, fmt.Errorf("getByRegNum: %w", err)
	}

	requestCar := &domain.RequestCar{
		ID:     car.ID,
		RegNum: car.RegNum,
		Mark:   car.Mark,
		Model:  car.Model,
		Year:   car.Year,
		Owner:  owner,
	}
	return requestCar, nil
}

func (r *carDatabase) Get(ctx context.Context, filter *filter.Car) ([]domain.RequestCar, error) {
	query := "SELECT * FROM car"
	var args []interface{}

	if filter != nil {
		var whereClauses []string
		if filter.ID != nil {
			args = append(args, *filter.ID)
			whereClauses = append(whereClauses, fmt.Sprintf("id = $%d", len(args)))
		}
		if filter.RegNum != nil {
			args = append(args, *filter.RegNum)
			whereClauses = append(whereClauses, fmt.Sprintf("regnum = $%d", len(args)))
		}
		if filter.Mark != nil {
			args = append(args, *filter.Mark)
			whereClauses = append(whereClauses, fmt.Sprintf("mark = $%d", len(args)))
		}
		if filter.Model != nil {
			args = append(args, *filter.Model)
			whereClauses = append(whereClauses, fmt.Sprintf("model = $%d", len(args)))
		}
		if filter.Year != nil {
			whereClauses = append(whereClauses, "year = $5")
			args = append(args, *filter.Year)
		}
		if filter.Owner != nil {
			whereClauses = append(whereClauses, "owner = $6")
			args = append(args, *filter.Owner)
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

	var cars []domain.RequestCar
	for rows.Next() {
		var car domain.Car
		if err := rows.Scan(&car.ID, &car.RegNum, &car.Mark, &car.Model, &car.Year, &car.Owner); err != nil {
			return nil, err
		}
		owner, err := r.getOwner(ctx, car.Owner)
		if err != nil {
			return nil, err
		}

		requestCar := domain.RequestCar{
			ID:     car.ID,
			RegNum: car.RegNum,
			Mark:   car.Mark,
			Model:  car.Model,
			Year:   car.Year,
			Owner:  owner,
		}
		cars = append(cars, requestCar)
	}
	return cars, nil
}
func (r *carDatabase) Update(ctx context.Context, id int64, updatedCar domain.Car) (*domain.RequestCar, error) {
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

	owner, err := r.getOwner(ctx, updatedCar.Owner)
	if err != nil {
		return nil, fmt.Errorf("update people: %v", err)
	}

	requestCar := &domain.RequestCar{
		ID:     updatedCar.ID,
		RegNum: updatedCar.RegNum,
		Mark:   updatedCar.Mark,
		Model:  updatedCar.Model,
		Year:   updatedCar.Year,
		Owner:  owner,
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return nil, err
	}

	if rowsAffected == 0 {
		return nil, db.ErrUpdateFailed
	}

	return requestCar, nil
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

func (r *carDatabase) getOwner(ctx context.Context, ownerID int64) (domain.People, error) {
	row := r.db.QueryRowContext(ctx, "SELECT * FROM people WHERE id = $1", ownerID)
	if row == nil {
		return domain.People{}, db.ErrOwnerNotFound
	}

	var owner domain.People
	err := row.Scan(&owner.ID, &owner.Name, &owner.SurName, &owner.Patronymic)
	if err != nil {
		return domain.People{}, db.ErrOwnerNotFound
	}

	return owner, nil
}
