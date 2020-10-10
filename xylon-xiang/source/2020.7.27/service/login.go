package service

import (
	"2020.7.27/middlewr"
	"2020.7.27/module_mapper"
)

func LoginService(id string, password string) (isLog bool, token string, err error) {

	result, err := module_mapper.FindMapper("id", id, false)
	if err != nil {
		return false, "", err
	}

	if password != result.(module_mapper.User).Password {
		return false, "", nil
	}

	token, err = middlewr.CreateJwtToken(result.(module_mapper.User).ID)
	if err != nil {
		return true, "", err
	}

	return true, token, nil
}
