package controller

type Workout struct {
   ID        int    
   UserID    int    
   Name      string    
   Exercises []Exercise
}

type Exercise struct {
   ID     int
   Sets   int
   Reps   int
   Weight float64
}

type workoutReq struct {
   userID    int
   name      string
   exercises []Exercise
}

type alterWorkoutReq struct {
   id        int
   userID    int
   name      *string
   exercises *[]Exercise
}

type removeWorkoutReq struct {
   id     int
   userID int
}

type signUpReq struct {
   email    string
   password string
}

type signInReq struct {
   email    string
   password string
}

type workoutResp struct {
   workout Workout
}

type workoutsResp struct {
   list []Workout
}

type signUpResp struct {
   err error
}

type signInResp struct {
   token string
   err   error
}

func (r *workoutReq) GetUserID() int {
   return r.userID
}

func (r *workoutReq) GetName() string {
   return r.name
}

func (r *workoutReq) GetExercises() []Exercise {
   return r.exercises
}

func (r *alterWorkoutReq) GetID() int {
   return r.id
}

func (r *alterWorkoutReq) GetUserID() int {
   return r.userID
}

func (r *alterWorkoutReq) GetName() *string {
   return r.name
}

func (r *alterWorkoutReq) GetExercises() *[]Exercise {
   return r.exercises
}

func (r *removeWorkoutReq) GetID() int {
   return r.id
}

func (r *removeWorkoutReq) GetUserID() int {
   return r.userID
}

func (r *signUpReq) GetEmail() string {
   return r.email
}

func (r *signUpReq) GetPassword() string {
   return r.password
}

func (r *signInReq) GetEmail() string {
   return r.email
}

func (r *signInReq) GetPassword() string {
   return r.password
}

func (r *workoutResp) GetWorkout() Workout {
   return r.workout
}

func (r *workoutsResp) GetList() []Workout {
   return r.list
}

func (r *signUpResp) GetError() error {
   return r.err
}

func (r *signInResp) GetToken() string {
   return r.token
}

func (r *signInResp) GetError() error {
   return r.err
}


func NewWorkoutRequest(userID int, name string, exercises []Exercise) *workoutReq {
   return &workoutReq{
       userID:    userID,
       name:      name,
       exercises: exercises,
   }
}

func NewAlterWorkoutRequest(id, userID int, name *string, exercises *[]Exercise) *alterWorkoutReq {
   return &alterWorkoutReq{
       id:        id,
       userID:    userID,
       name:      name,
       exercises: exercises,
   }
}

func NewRemoveWorkoutRequest(id, userID int) *removeWorkoutReq {
   return &removeWorkoutReq{
       id:     id,
       userID: userID,
   }
}

func NewSignUpRequest(email, password string) *signUpReq {
   return &signUpReq{
       email:    email,
       password: password,
   }
}

func NewSignInRequest(email, password string) *signInReq {
   return &signInReq{
       email:    email,
       password: password,
   }
}

func NewWorkoutResponse(workout Workout) *workoutResp {
   return &workoutResp{workout: workout}
}

func NewWorkoutsResponse(list []Workout) *workoutsResp {
   return &workoutsResp{list: list}
}

func NewSignUpResponse(err error) *signUpResp {
   return &signUpResp{err: err}
}

func NewSignInResponse(token string, err error) *signInResp {
   return &signInResp{
       token: token,
       err:   err,
   }
}