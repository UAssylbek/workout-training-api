package controller

type RemoveWorkoutReq interface {
    GetID() string        
    GetUserID() string    
}

type RemoveWorkoutResp interface{}