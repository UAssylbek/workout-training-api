package controller

type AlterWorkoutReq interface {
    GetID() int        
    GetUserID() int    
    GetName() *string
    GetExercises() *[]Exercise
}

type AlterWorkoutResp interface {
    GetWorkout() Workout
}