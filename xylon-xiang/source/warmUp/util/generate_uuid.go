package util

import (
	"fmt"
	"github.com/satori/go.uuid"
)

func GenerateUUID() (uid string, err error) {

	u1 := uuid.Must(uuid.NewV4(), err)

	uid = fmt.Sprint(u1)

	return uid, err
}
