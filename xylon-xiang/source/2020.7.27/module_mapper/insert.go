package module_mapper

import "database/sql"

func InsertMapper(user User) error {
	_, err := UserSysDB.Query("SELECT * FROM uesrSys LIMIT 1")
	if err == sql.ErrNoRows {
		_, err = UserSysDB.Exec("CREATE TABLE example (id integer, password varchar," +
			" name varchar, phone varchar, email varchar, admin integer)")
	}

	insForm, err := UserSysDB.Prepare("INSERT INTO userSys (id, name, password, email, phone) VALUES (?,?,?,?,?)")
	if err != nil {
		return err
	}

	_, _ = insForm.Exec(user.ID, user.Name, user.Password, user.Email, user.Phone)

	return nil
}
