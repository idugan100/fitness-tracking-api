package database

import (
	"database/sql"
	"fmt"
)

type Seeder struct {
	DB *sql.DB
}

func NewSeeder(db *sql.DB) Seeder {
	return Seeder{DB: db}
}

func (s Seeder) Clear() error {
	fmt.Printf("clear running")
	return nil
	// _, err := s.DB.Exec("DELETE FROM Cardio")
	// if err != nil {
	// 	return errors.New("error when truncating Cardio table: " + err.Error())
	// }

	// _, err = s.DB.Exec("DELETE FROM Lifts")
	// if err != nil {
	// 	return errors.New("error when truncating Lifts table: " + err.Error())
	// }

	// _, err = s.DB.Exec("DELETE FROM Workouts")
	// if err != nil {
	// 	return errors.New("error when truncating Workout table: " + err.Error())
	// }

	// _, err = s.DB.Exec("DELETE FROM LiftingLog")
	// if err != nil {
	// 	return errors.New("error when truncating LiftingLog table: " + err.Error())
	// }

	// _, err = s.DB.Exec("DELETE FROM CardioLog")
	// if err != nil {
	// 	return errors.New("error when truncating CardioLog table: " + err.Error())
	// }

	// return nil
}

func (s Seeder) Seed() error {
	// cardioList := [4]models.Cardio{
	// 	models.Cardio{Id: 1, Name: "running"},
	// 	models.Cardio{Id: 2, Name: "swimming"},
	// 	models.Cardio{Id: 3, Name: "biking"},
	// 	models.Cardio{Id: 4, Name: "hiking"},
	// }

	// for cardio := range cardioList {
	// 	s.DB.Exec("INSERT INTO")
	// }
	// //seed all tables
	fmt.Printf("seeder running")
	return nil
}
