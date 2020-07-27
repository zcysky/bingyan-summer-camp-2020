package module_mapper

// delete by id
func Delete(id string) error {

	delForm, err := UserSysDB.Prepare("DELETE FROM userSys WHERE id=?")
	if err != nil {
		return err
	}

	_, _ = delForm.Exec(id)

	return nil
}
