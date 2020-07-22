package service

import "warmUp/module_mapper"

func UpdateUserInfo(user module_mapper.User) {

	module_mapper.UpdateMapper(user)

}
