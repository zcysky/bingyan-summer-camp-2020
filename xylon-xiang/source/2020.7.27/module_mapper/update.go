package module_mapper

import (
	"log"
)

func UpdateMapper(user User) {

	insForm, err := UserSysDB.Prepare("UPDATE userSys SET name=?, password=?,phone=?, email=? WHERE id=?")
	if err != nil {
		log.Fatal(err)
	}

	_, _ = insForm.Exec(user.Name, user.Password, user.Phone, user.Email, user.ID)

}
