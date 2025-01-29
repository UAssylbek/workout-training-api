package controller

type ListWorkoutReq interface {
    GetUserID() int    
}

type ListWorkoutResp interface {
    GetList() []Workout
}
