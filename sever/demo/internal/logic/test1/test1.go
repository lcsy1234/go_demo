package test1

import "demo/internal/service"

type sTest1 struct{}

func init() {
	service.RegisterTest1(New())
}

func New() *sTest1 {
	return &sTest1{}
}
