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

type controller struct {
    workoutRepo WorkoutRepository    
    authRepo    AuthRepository
}

type WorkoutRepository interface {
    Create(ctx context.Context, workout *Workout) error
    Update(ctx context.Context, workout *Workout) error
    Delete(ctx context.Context, id int) error
    List(ctx context.Context, userID int) ([]Workout, error)
}

type AuthRepository interface {
    CreateUser(ctx context.Context, email, hashedPassword string) error
    GetUser(ctx context.Context, email string) (string, error) // возвращает хешированный пароль
}

func NewController(workoutRepo WorkoutRepository, authRepo AuthRepository) Controller {
    return &controller{
        workoutRepo: workoutRepo,
        authRepo:    authRepo,
    }
}

func (c *controller) ListWorkout(ctx context.Context, req ListWorkoutReq) (ListWorkoutResp, error) {
    return nil, nil
}

func (c *controller) NewWorkout(ctx context.Context, req NewWorkoutReq) (NewWorkoutResp, error) {
    return nil, nil
}

func (c *controller) AlterWorkout(ctx context.Context, req AlterWorkoutReq) (AlterWorkoutResp, error) {
    return nil, nil
}

func (c *controller) RemoveWorkout(ctx context.Context, req RemoveWorkoutReq) (RemoveWorkoutResp, error) {
    return nil, nil
}

func (c *controller) SignUp(ctx context.Context, req SignUpReq) (SignUpResp, error) {
    return nil, nil
}

func (c *controller) SignIn(ctx context.Context, req SignInReq) (SignInResp, error) {
    return nil, nil
}