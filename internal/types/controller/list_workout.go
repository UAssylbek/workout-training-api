package controller

type ListWorkoutReq interface {
    GetUserID() string    
}

type ListWorkoutResp interface {
    GetList() []Workout
}