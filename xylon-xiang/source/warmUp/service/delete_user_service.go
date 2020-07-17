package service

import (
	"warmUp/module_mapper"
)

func DeleteUserService(id string) (bool, error) {
	tmp, err := module_mapper.FindMapper("id", id, false)
	if err != nil {
		return false, err
	}
	result := tmp.(module_mapper.User)

	if result.Admin == 0 {
		return false, nil
	} else {
		err = module_mapper.Delete(result.ID)
		if err != nil {
			return false, err
		}
		return true, nil
	}
}
