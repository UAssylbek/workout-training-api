package controller

type AlterWorkoutReq interface {
    GetID() string        
    GetUserID() string    
    GetName() *string
    GetExercises() *[]Exercise
}

type AlterWorkoutResp interface {
    GetWorkout() Workout
}