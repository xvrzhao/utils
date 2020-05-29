package errors

import (
	"errors"
	"log"
	"testing"
)

func TestWrapper(t *testing.T) {
	SetWrapperOpt(WrapperOpt{
		ProjectDir:    "/Users/xvrzhao/Documents/projects/micro-stacks/",
		ProjectModule: "github.com/micro-stacks/utils/",
	})
	if err := handler(); err != nil {
		log.Println(err)
	}
}

func handler() error {
	m := new(UserModel)
	_, err := m.GetUserByID(1)
	if err != nil {
		return Wrapper(err, "can not get user by id")
	}
	return nil
}

type UserModel struct{}

func (m UserModel) GetUserByID(int) (*UserModel, error) {
	if err := errors.New("can not connect to db"); err != nil {
		return nil, Wrapper(err, "query fail")
	}
	return &m, nil
}
