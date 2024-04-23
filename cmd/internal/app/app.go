//Тут я тестил репозитории к бд

package app

//import (
//	"context"
//	"errors"
//	"fmt"
//	"hh.ru/cmd/internal/repo"
//	"hh.ru/cmd/internal/repo/car"
//	"hh.ru/cmd/internal/repo/people"
//	"hh.ru/pkg/domain"
//	"log"
//)
//
//func RunCarRepositoryDemo(ctx context.Context, carRepository d.carDatabase) error {
//	fmt.Println("1. MIGRATE REPOSITORY")
//
//	if err := carRepository.Migrate(ctx); err != nil {
//		log.Fatal(err)
//	}
//
//	fmt.Println("2. CREATE RECORDS OF REPOSITORY")
//	carSample1 := domain.Car{
//		RegNum: "Я111XX243",
//		Mark:   "https://gosamples.dev",
//		Model:  "Vesta",
//		Year:   2002,
//		Owner:  1,
//	}
//	if err := carSample1.Validate(); err != nil {
//		log.Println(err.Error())
//	}
//
//	carSample2 := domain.Car{
//		RegNum: "SAMPLE",
//		Mark:   "https://gosamples.dev",
//		Model:  "FDS",
//		Year:   2005,
//		Owner:  2,
//	}
//
//	createdGosamples, err := carRepository.Create(ctx, carSample1)
//	if errors.Is(err, repo.ErrDuplicate) {
//		fmt.Printf("record: %+v already exists\n", carSample1)
//	} else if err != nil {
//		log.Fatal(err)
//	}
//	createdGolang, err := carRepository.Create(ctx, carSample2)
//	if errors.Is(err, repo.ErrDuplicate) {
//		log.Printf("record: %+v already exists\n", carSample2)
//	} else if err != nil {
//		log.Fatal(err)
//	}
//
//	fmt.Printf("%+v\n%+v\n", createdGosamples, createdGolang)
//
//	fmt.Println("3. GET RECORD BY ID")
//	gotGosamples, err := carRepository.GetByID(ctx, 1)
//	if errors.Is(err, repo.ErrNotExist) {
//		log.Println("record: GOSAMPLES does not exist in the repository")
//	} else if err != nil {
//		log.Fatal(err)
//	}
//
//	fmt.Printf("%+v\n", gotGosamples)
//
//	fmt.Println("4. UPDATE RECORD")
//	createdGosamples.Year = 1
//	if _, err := carRepository.Update(ctx, createdGosamples.ID, *createdGosamples); err != nil {
//		if errors.Is(err, repo.ErrDuplicate) {
//			fmt.Printf("record: %+v already exists\n", createdGosamples)
//		} else if errors.Is(err, repo.ErrUpdateFailed) {
//			fmt.Printf("update of record: %+v failed", createdGolang)
//		} else {
//			log.Fatal(err)
//		}
//	}
//
//	fmt.Println("5. GET ALL")
//	all, err := carRepository.All(ctx)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	for _, website := range all {
//		fmt.Printf("%+v\n", website)
//	}
//
//	fmt.Println("6. DELETE RECORD")
//	if err := carRepository.Delete(ctx, createdGolang.ID); err != nil {
//		if errors.Is(err, repo.ErrDeleteFailed) {
//			fmt.Printf("delete of record: %d failed", createdGolang.ID)
//		} else {
//			log.Fatal(err)
//		}
//	}
//
//	fmt.Println("7. GET ALL")
//	all, err = carRepository.All(ctx)
//	if err != nil {
//		log.Fatal(err)
//	}
//	for _, website := range all {
//		fmt.Printf("%+v\n", website)
//	}
//	return nil
//}
//
//func RunPeopleRepositoryDemo(ctx context.Context, peopleRepository people.RepositoryPeople) error {
//	fmt.Println("\n\n1. MIGRATE PEOPLE REPOSITORY")
//
//	if err := peopleRepository.Migrate(ctx); err != nil {
//		log.Fatal(err)
//	}
//
//	fmt.Println("2. CREATE RECORDS OF REPOSITORY")
//	carSample1 := domain.People{
//		Name:       "Boba",
//		SurName:    "Saymon",
//		Patronymic: "Joroto",
//	}
//	carSample2 := domain.People{
//		Name:       "SSS",
//		SurName:    "FFF",
//		Patronymic: "DDD",
//	}
//
//	createdGosamples, err := peopleRepository.Create(ctx, carSample1)
//	if errors.Is(err, repo.ErrDuplicate) {
//		fmt.Printf("record: %+v already exists\n", carSample1)
//	} else if err != nil {
//		log.Fatal(err)
//	}
//	createdGolang, err := peopleRepository.Create(ctx, carSample2)
//	if errors.Is(err, repo.ErrDuplicate) {
//		log.Printf("record: %+v already exists\n", carSample2)
//	} else if err != nil {
//		log.Fatal(err)
//	}
//
//	fmt.Printf("%+v\n%+v\n", createdGosamples, createdGolang)
//
//	fmt.Println("3. GET RECORD BY ID")
//	gotGosamples, err := peopleRepository.GetByID(ctx, 2)
//	if errors.Is(err, repo.ErrNotExist) {
//		log.Println("record: GOSAMPLES does not exist in the repository")
//	} else if err != nil {
//		log.Fatal(err)
//	}
//
//	fmt.Printf("%+v\n", gotGosamples)
//
//	fmt.Println("4. UPDATE RECORD")
//	createdGosamples.Name = "Dad"
//	if _, err := peopleRepository.Update(ctx, createdGosamples.ID, *createdGosamples); err != nil {
//		if errors.Is(err, repo.ErrDuplicate) {
//			fmt.Printf("record: %+v already exists\n", createdGosamples)
//		} else if errors.Is(err, repo.ErrUpdateFailed) {
//			fmt.Printf("update of record: %+v failed", createdGolang)
//		} else {
//			log.Fatal(err)
//		}
//	}
//
//	fmt.Println("5. GET ALL")
//	all, err := peopleRepository.All(ctx)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	for _, website := range all {
//		fmt.Printf("%+v\n", website)
//	}
//
//	fmt.Println("6. DELETE RECORD")
//	if err := peopleRepository.Delete(ctx, createdGolang.ID); err != nil {
//		if errors.Is(err, repo.ErrDeleteFailed) {
//			fmt.Printf("delete of record: %d failed", createdGolang.ID)
//		} else {
//			log.Fatal(err)
//		}
//	}
//
//	fmt.Println("7. GET ALL")
//	all, err = peopleRepository.All(ctx)
//	if err != nil {
//		log.Fatal(err)
//	}
//	for _, website := range all {
//		fmt.Printf("%+v\n", website)
//	}
//	return nil
//}
