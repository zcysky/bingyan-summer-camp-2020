package service

import (
	"2020.7.27/middlewr"
	"2020.7.27/module_mapper"
	"2020.7.27/util"
)

func RegisterService(registerUser *module_mapper.RegisterUser) (done bool, uid string, jwt string, err error) {
	if registerUser.RegisterCode == "" {
		err := module_mapper.GenerateRegisterCode(registerUser.Email)
		if err != nil {
			return false, "", "", err
		}

		return true, "", "", nil
	}
	done, err = module_mapper.AuthRegisterCode(registerUser.RegisterCode)
	if err != nil {
		return false, "", "", err
	}

	if !done {
		return done, "", "", nil
	}

	uuid, err := util.GenerateUUID()
	if err != nil {
		return done, "", "", err
	}

	user := module_mapper.User{
		ID:       uuid,
		Password: registerUser.Password,
		Email:    registerUser.Email,
		Phone:    registerUser.Email,
		Name:     registerUser.Name,
	}

	err = module_mapper.InsertMapper(user)

	token, err := middlewr.CreateJwtToken(uuid)
	if err != nil {
		return done, uuid, "", err
	}
	return done, uuid, token, err

}
