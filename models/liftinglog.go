package models

type LiftingLog struct {
	Id        int `json:"id"`
	LiftId    int `json:"lift_id" binding:"required"`
	Weight    int `json:"weight" binding:"required"`
	Sets      int `json:"sets" biding:"required"`
	Reps      int `json:"reps" biding:"required"`
	WorkoutId int `json:"workout_id" biding:"required"`
}
