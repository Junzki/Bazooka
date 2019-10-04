package controllers

import (
	"bazooka/internal/pkg/models"
	"errors"
)

type UserController struct {
	saver IUserSaver
}

func (c UserController) Validate(u *models.User) error {
	if nil == u || !u.IsValid() {
		return errors.New("invalid user")
	}

	return nil
}

func (c UserController) GetUser(u *models.User) error {
	var err error = nil

	if err = c.Validate(u); nil != err {
		return err
	}

	return c.saver.Load(u)
}

func (c UserController) CreateUser(u *models.User) error {
	var err error = nil
	err = c.Validate(u)
	if nil != err {
		return err
	}

	return c.saver.Save(u)
}

func (c UserController) UpdateUser(u *models.User) error {
	var err error = nil

	err = c.Validate(u)
	if nil != err {
		return err
	}

	existed := &models.User{
		Uid: u.Uid,
	}
	err = c.saver.Load(existed)
	if nil != err {
		return errors.New("user does not exist")
	}

	u.ID = existed.ID
	return c.saver.Save(u)
}

var userController *UserController

func GetUserController() *UserController {
	if nil != userController {
		return userController
	}

	userController = &UserController{
		saver: &UserSaver{},
	}
	return userController
}
