package app

import (
	"context"
	"errors"
	"fmt"
	"hh.ru/cmd/internal/model"
	"hh.ru/cmd/internal/repo"
	"log"
)

func RunRepositoryDemo(ctx context.Context, carRepository repo.PgSQLClassicRepository) error {
	fmt.Println("1. MIGRATE REPOSITORY")

	if err := carRepository.Migrate(ctx); err != nil {
		log.Fatal(err)
	}

	fmt.Println("2. CREATE RECORDS OF REPOSITORY")
	carSample1 := model.Car{
		RegNum: "REGNUM",
		Mark:   "https://gosamples.dev",
		Model:  "Vesta",
		Year:   2002,
		Owner:  1,
	}
	carSample2 := model.Car{
		RegNum: "SAMPLE",
		Mark:   "https://gosamples.dev",
		Model:  "FDS",
		Year:   2005,
		Owner:  2,
	}

	createdGosamples, err := carRepository.Create(ctx, carSample1)
	if errors.Is(err, repo.ErrDuplicate) {
		fmt.Printf("record: %+v already exists\n", carSample1)
	} else if err != nil {
		log.Fatal(err)
	}
	createdGolang, err := carRepository.Create(ctx, carSample2)
	if errors.Is(err, repo.ErrDuplicate) {
		log.Printf("record: %+v already exists\n", carSample2)
	} else if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v\n%+v\n", createdGosamples, createdGolang)

	fmt.Println("3. GET RECORD BY ID")
	gotGosamples, err := carRepository.GetCarByID(ctx, 2)
	if errors.Is(err, repo.ErrNotExist) {
		log.Println("record: GOSAMPLES does not exist in the repository")
	} else if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v\n", gotGosamples)

	fmt.Println("4. UPDATE RECORD")
	createdGosamples.Year = 1
	createdGosamples.Owner = 2
	if _, err := carRepository.Update(ctx, createdGosamples.ID, *createdGosamples); err != nil {
		if errors.Is(err, repo.ErrDuplicate) {
			fmt.Printf("record: %+v already exists\n", createdGosamples)
		} else if errors.Is(err, repo.ErrUpdateFailed) {
			fmt.Printf("update of record: %+v failed", createdGolang)
		} else {
			log.Fatal(err)
		}
	}

	fmt.Println("5. GET ALL")
	all, err := carRepository.GetAll(ctx)
	if err != nil {
		log.Fatal(err)
	}

	for _, website := range all {
		fmt.Printf("%+v\n", website)
	}

	fmt.Println("6. DELETE RECORD")
	if err := carRepository.DeleteCarByID(ctx, createdGolang.ID); err != nil {
		if errors.Is(err, repo.ErrDeleteFailed) {
			fmt.Printf("delete of record: %d failed", createdGolang.ID)
		} else {
			log.Fatal(err)
		}
	}

	fmt.Println("7. GET ALL")
	all, err = carRepository.GetAll(ctx)
	if err != nil {
		log.Fatal(err)
	}
	for _, website := range all {
		fmt.Printf("%+v\n", website)
	}
	return nil
}
