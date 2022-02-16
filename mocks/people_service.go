package mocks

import (
	"github.com/marklude/go-sovtech/graph/model"
	"github.com/stretchr/testify/mock"
)

type MockedPeopleResolver struct {
	mock.Mock
}

func (r *MockedPeopleResolver) Authentication(input model.NewUser) *model.Token {
	args := r.Called(input)
	return args.Get(0).(*model.Token)
}

func (r *MockedPeopleResolver) Peoples(first *int) ([]*model.People, error) {
	args := r.Called(first)
	return args.Get(0).([]*model.People), nil
}

func (r *MockedPeopleResolver) PeopleByName(name string) (*model.People, error) {
	args := r.Called(name)
	return args.Get(0).(*model.People), nil
}
