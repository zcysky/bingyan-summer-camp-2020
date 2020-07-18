package service

import (
	"warmUp/module_mapper"
	"warmUp/util"
)

func RegisterService(registerUser module_mapper.RegisterUser) (done bool, err error) {
	if registerUser.RegisterCode == ""{
		err := module_mapper.GenerateRegisterCode(registerUser.Email)
		if err != nil{
			return false, err
		}
	}
	done, err = module_mapper.AuthRegisterCode(registerUser.RegisterCode)
	if err != nil{
		return false, err
	}

	if !done {
		return done, nil
	}

	uuid, err := util.GenerateUUID()
	if err != nil{
		return done, err
	}

	user := module_mapper.User{
		ID: uuid,
		Password: registerUser.Password,
		Email: registerUser.Email,
		Phone: registerUser.Email,
		Name: registerUser.Name,
	}

	err = module_mapper.InsertMapper(user)

	return done, err

}
