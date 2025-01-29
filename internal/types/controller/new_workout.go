package controller

type NewWorkoutReq interface {
    GetUserID() int
    GetName() string
    GetExercises() []Exercise
}

type NewWorkoutResp interface {
    GetWorkout() Workout
}
