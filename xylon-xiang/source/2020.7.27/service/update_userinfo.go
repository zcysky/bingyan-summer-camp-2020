package service

import "2020.7.27/module_mapper"

func UpdateUserInfo(user module_mapper.User) {

	module_mapper.UpdateMapper(user)

}
