package service

import "fmt"

type testService struct{}

type TestService interface {
	HelloWorld() string
}

func NewTestService() TestService {
	return &testService{}
}

func (ts *testService) HelloWorld() string {
	fmt.Println("Hello World")

	return "hello world"
}
