package service

import "warmUp/module_mapper"

// bool auth whether host is a admin
func GetUserInfoService(isAll bool, hostId string, id string) (bool, interface{}, error) {
	hostResult, err := module_mapper.FindMapper("id", hostId, false)
	if err != nil{
		return false, nil, err
	}
	admin := hostResult.(module_mapper.User).Admin
	if admin == 0 {
		return false, nil, nil
	} else {
		if !isAll{
			idResult, err := module_mapper.FindMapper("id", id, false)
			if err != nil{
				return true, nil, err
			}
			return true, idResult, nil
		}

		idResult, err := module_mapper.FindMapper("id", "", true)
		if err != nil{
			return true, nil, err
		}

		return true, idResult, nil
	}

}
