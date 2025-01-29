package controller

type SignUpReq interface {
    GetEmail() string
    GetPassword() string
}

type SignUpResp interface {
    GetError() error
}