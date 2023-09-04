package users

import (
	"fmt"
	"time"

	"github.com/ShineWaine/cursogo/modelos"
)

func AltaUsuario() {
	u := new(modelos.User)
	u.AddUser(10, "Pablo", time.Now(), true)
	fmt.Println(u)
}
