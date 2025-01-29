package controller

type RemoveWorkoutReq interface {
    GetID() int        
    GetUserID() int    
}

type RemoveWorkoutResp interface{}