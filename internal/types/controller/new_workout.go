package controller

type NewWorkoutReq interface {
    GetUserID() string
    GetName() string
    GetExercises() []Exercise
}

type NewWorkoutResp interface {
    GetWorkout() Workout
}