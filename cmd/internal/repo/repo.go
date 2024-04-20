package repo

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5/pgconn"
	"hh.ru/cmd/internal/model"
	"log"
)

// DB represents a Database instance

var (
	ErrMigrate      = errors.New("migration failed")
	ErrDuplicate    = errors.New("record already exists")
	ErrNotExist     = errors.New("row does not exist")
	ErrUpdateFailed = errors.New("update failed")
	ErrDeleteFailed = errors.New("delete failed")
)

type PgSQLClassicRepository struct {
	db *sql.DB
}

func NewPgSqlClassicRepository(db *sql.DB) *PgSQLClassicRepository {
	return &PgSQLClassicRepository{
		db: db,
	}
}

type Repository interface {
	Migrate(ctx context.Context) error
	Create(ctx context.Context, car model.Car) (*model.Car, error)
	GetAll(ctx context.Context) ([]model.Car, error)
	GetCarByID(ctx context.Context, id int64) (*model.Car, error)
	Update(ctx context.Context, id int64, updatedCar model.Car) (*model.Car, error)
	DeleteCarByID(ctx context.Context, id int64) error
}

func (r *PgSQLClassicRepository) Migrate(ctx context.Context) error {
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

	peopleQuery := `
    CREATE TABLE IF NOT EXISTS people(
		id SERIAL PRIMARY KEY,
		name text not NULL,
		surName text not NULL,
		patronymic text not NULL
    );
    `
	_, err := r.db.ExecContext(ctx, carQuery)
	if err != nil {
		message := ErrMigrate.Error() + " car"
		log.Printf("%q: %s\n", message, err.Error())
		return ErrMigrate
	}

	_, err = r.db.ExecContext(ctx, peopleQuery)
	if err != nil {
		message := ErrMigrate.Error() + " people"
		log.Printf("%q: %s\n", message, err.Error())
		return ErrMigrate
	}

	return err
}

func (r *PgSQLClassicRepository) Create(ctx context.Context, car model.Car) (*model.Car, error) {
	var id int64
	err := r.db.QueryRowContext(ctx, "INSERT INTO car(regnum, mark, model, year, owner) values($1, $2, $3, $4, $5) RETURNING id", car.RegNum, car.Mark, car.Model, car.Year, car.Owner).Scan(&id)
	if err != nil {
		var pgxError *pgconn.PgError
		if errors.As(err, &pgxError) {
			if pgxError.Code == "23505" {
				return nil, ErrDuplicate
			}
		}
		return nil, err
	}
	car.ID = id

	return &car, nil
}

func (r *PgSQLClassicRepository) GetAll(ctx context.Context) ([]model.Car, error) {
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

func (r *PgSQLClassicRepository) GetCarByID(ctx context.Context, id int64) (*model.Car, error) {
	row := r.db.QueryRowContext(ctx, "SELECT * FROM car WHERE id = $1", id)

	var car model.Car
	if err := row.Scan(&car.ID, &car.RegNum, &car.Mark, &car.Model, &car.Year, &car.Owner); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNotExist
		}
		return nil, err
	}
	return &car, nil
}

func (r *PgSQLClassicRepository) Update(ctx context.Context, id int64, updatedCar model.Car) (*model.Car, error) {
	res, err := r.db.ExecContext(ctx, "UPDATE car SET RegNum = $1, Mark = $2, Model = $3, Year = $4, owner = $5 WHERE id = $4", updatedCar.RegNum, updatedCar.Mark, fmt.Sprint("%v", updatedCar.Year), updatedCar.Owner, id)
	if err != nil {
		var pgxError *pgconn.PgError
		if errors.As(err, &pgxError) {
			if pgxError.Code == "23505" {
				return nil, ErrDuplicate
			}
		}
		return nil, err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return nil, err
	}

	if rowsAffected == 0 {
		return nil, ErrUpdateFailed
	}

	return &updatedCar, nil
}

func (r *PgSQLClassicRepository) DeleteCarByID(ctx context.Context, id int64) error {
	res, err := r.db.ExecContext(ctx, "DELETE FROM car WHERE id = $1", id)
	if err != nil {
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return ErrDeleteFailed
	}

	return err
}

//func ConnectToDB() {
//	err := godotenv.Load()
//	if err != nil {
//		log.Fatal("Error loading env file \n", err)
//	}
//
//	dsn := fmt.Sprintf("host=localhost user=%s password=%s dbname=%s port=%s sslmode=disable",
//		os.Getenv("PSQL_USER"), os.Getenv("PSQL_PASS"), os.Getenv("PSQL_DBNAME"), os.Getenv("PSQL_PORT"))
//
//	log.Print("Connecting to PostgreSQL DB...")
//	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
//		Logger: logger.Default.LogMode(logger.Info),
//	})
//
//	if err != nil {
//		log.Fatal("Failed to connect to database. \n", err)
//		os.Exit(2)
//
//	}
//	log.Println("connected")
//}
