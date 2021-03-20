package module_mapper

import (
	"database/sql"
)

// return the result
func FindMapper(name string, i string, all bool, order ...string) (interface{}, error) {

	if !all {
		var result User

		selDB, err := UserSysDB.Query("SELECT * FROM userSys WHERE ? = ?", name, i)
		if err != nil {
			return nil, nil
		}
		for selDB.Next() {
			var (
				id       string
				password string
				name     string
				phone    string
				email    string
				admin    int
			)

			err = selDB.Scan(&id, &password, &name, &phone, &email, &admin)
			if err != nil {
				return nil, nil
			}
			result.ID = id
			result.Password = password
			result.Name = name
			result.Phone = phone
			result.Email = email
			result.Admin = admin
		}

		return result, nil
	} else {
		var (
			result []User
			selDB  *sql.Rows
			err    error
		)

		if order != nil {
			selDB, err = UserSysDB.Query("SELECT * FROM userSys ORDER BY ? LIMIT 20, 40", order)
			if err != nil {
				return nil, err

			}
		}
		selDB, err = UserSysDB.Query("SELECT * FROM userSys")
		if err != nil {
			return nil, err
		}

		emp := User{}
		for selDB.Next() {
			var (
				id       string
				password string
				name     string
				phone    string
				email    string
				admin    int
			)
			err = selDB.Scan(&id, &password, &name, &phone, &email, &admin)
			if err != nil {
				return nil, nil
			}
			emp.ID = id
			emp.Password = password
			emp.Name = name
			emp.Phone = phone
			emp.Email = email
			emp.Admin = admin

			result = append(result, emp)
		}

		return result, nil
	}
}
