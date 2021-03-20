package service

import (
	"warmUp/module_mapper"
)

func DeleteUserService(host string, id string) (bool, error) {
	tmp, err := module_mapper.FindMapper("id", host, false)
	if err != nil {
		return false, err
	}
	result := tmp.(module_mapper.User)

	if result.Admin == 0 {
		return false, nil
	} else {
		err = module_mapper.Delete(id)
		if err != nil {
			return false, err
		}
		return true, nil
	}
}
