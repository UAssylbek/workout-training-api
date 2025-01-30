package controller

import "context"

type Controller interface {
    SignUp(context.Context, SignUpReq) (SignUpResp, error)
    SignIn(context.Context, SignInReq) (SignInResp, error)
    ListWorkout(context.Context, ListWorkoutReq) (ListWorkoutResp, error)
    NewWorkout(context.Context, NewWorkoutReq) (NewWorkoutResp, error)
    AlterWorkout(context.Context, AlterWorkoutReq) (AlterWorkoutResp, error)
    RemoveWorkout(context.Context, RemoveWorkoutReq) (RemoveWorkoutResp, error)
}

type WorkoutRepository interface {
    Create(ctx context.Context, workout *Workout) error
    Update(ctx context.Context, workout *Workout) error
    Delete(ctx context.Context, id string) error
    List(ctx context.Context, userID string) ([]Workout, error)
}

type AuthRepository interface {
    CreateUser(ctx context.Context, email, hashedPassword string) error
    GetUser(ctx context.Context, email string) (string, error) 
}