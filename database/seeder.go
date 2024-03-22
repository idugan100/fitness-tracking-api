package database

import (
	"database/sql"
	"errors"
	"fitness-tracker-api/testbackend/models"
	"fmt"
)

type Seeder struct {
	DB *sql.DB
}

func NewSeeder(db *sql.DB) Seeder {
	return Seeder{DB: db}
}

func (s Seeder) Clear() error {
	_, err := s.DB.Exec("DELETE FROM Cardio")
	if err != nil {
		return errors.New("error when truncating Cardio table: " + err.Error())
	}

	_, err = s.DB.Exec("DELETE FROM Lifts")
	if err != nil {
		return errors.New("error when truncating Lifts table: " + err.Error())
	}

	_, err = s.DB.Exec("DELETE FROM Workouts")
	if err != nil {
		return errors.New("error when truncating Workout table: " + err.Error())
	}

	_, err = s.DB.Exec("DELETE FROM LiftingLog")
	if err != nil {
		return errors.New("error when truncating LiftingLog table: " + err.Error())
	}

	_, err = s.DB.Exec("DELETE FROM CardioLog")
	if err != nil {
		return errors.New("error when truncating CardioLog table: " + err.Error())
	}
	return nil
}

func (s Seeder) Seed() error {
	cardioList := [4]models.Cardio{
		{Id: 1, Name: "running"},
		{Id: 2, Name: "swimming"},
		{Id: 3, Name: "biking"},
		{Id: 4, Name: "hiking"},
	}

	for _, cardio := range cardioList {
		_, err := s.DB.Exec("INSERT INTO Cardio (id,name) VALUES (?,?)", cardio.Id, cardio.Name)
		if err != nil {
			return fmt.Errorf("error seeding cardio: %w", err)
		}
	}

	t := true
	f := false
	liftList := [4]models.Lift{
		{Id: 1, Name: "deadlift", Compound: &t, Upper: &f, Lower: &t},
		{Id: 2, Name: "squat", Compound: &t, Upper: &f, Lower: &t},
		{Id: 3, Name: "bench press", Compound: &t, Upper: &t, Lower: &f},
		{Id: 4, Name: "bicep curl", Compound: &f, Upper: &t, Lower: &f},
	}

	for _, lift := range liftList {
		_, err := s.DB.Exec("INSERT INTO Lifts (id, name, compound, upper, lower) VALUES (?,?,?,?,?)", lift.Id, lift.Name, lift.Compound, lift.Upper, lift.Lower)
		if err != nil {
			return fmt.Errorf("error seeding lifts: %w", err)
		}
	}

	workoutList := [3]models.Workout{
		{Id: 1, Location: "UFC", Notes: "felt sick this day"},
		{Id: 2, Location: "Around Campus", Notes: "was cold out"},
		{Id: 3, Location: "UFC", Notes: "was sore after this one"},
	}

	for _, workout := range workoutList {
		_, err := s.DB.Exec("INSERT INTO Workouts (id, location, notes) VALUES (?,?,?)", workout.Id, workout.Location, workout.Notes)
		if err != nil {
			return fmt.Errorf("error seeding workouts: %w", err)
		}
	}

	cardioLogList := [5]models.CardioLog{
		{Id: 1, CardioId: 1, WorkoutId: 1, Time: 10, Distance: 1.2},
		{Id: 2, CardioId: 2, WorkoutId: 1, Time: 20, Distance: 2.5},
		{Id: 3, CardioId: 3, WorkoutId: 2, Time: 8, Distance: 1.0},
		{Id: 4, CardioId: 4, WorkoutId: 3, Time: 30, Distance: 3.1},
		{Id: 5, CardioId: 5, WorkoutId: 3, Time: 45, Distance: 4.2},
	}

	query_string := "INSERT INTO CardioLog (id, cardio_id, workout_id, time, distance) VALUES (?,?,?,?,?)"
	for _, cardioLog := range cardioLogList {
		_, err := s.DB.Exec(query_string, cardioLog.Id, cardioLog.CardioId, cardioLog.WorkoutId, cardioLog.Time, cardioLog.Distance)
		if err != nil {
			return fmt.Errorf("error seeding CardioLog: %w", err)
		}
	}

	liftingLogList := [5]models.LiftingLog{
		{Id: 1, LiftId: 1, WorkoutId: 1, Sets: 3, Reps: 3, Weight: 315},
		{Id: 2, LiftId: 2, WorkoutId: 1, Sets: 3, Reps: 5, Weight: 225},
		{Id: 3, LiftId: 3, WorkoutId: 3, Sets: 3, Reps: 8, Weight: 135},
		{Id: 4, LiftId: 3, WorkoutId: 1, Sets: 3, Reps: 3, Weight: 165},
		{Id: 5, LiftId: 4, WorkoutId: 2, Sets: 3, Reps: 12, Weight: 65},
	}

	query_string = "INSERT INTO LiftingLog (id, lift_id, workout_id, sets, reps, weight) VALUES (?,?,?,?,?,?)"
	for _, liftingLog := range liftingLogList {
		_, err := s.DB.Exec(query_string, liftingLog.Id, liftingLog.LiftId, liftingLog.WorkoutId, liftingLog.Sets, liftingLog.Reps, liftingLog.Weight)
		if err != nil {
			return fmt.Errorf("error seeding LiftingLog: %w", err)
		}
	}

	return nil
}
