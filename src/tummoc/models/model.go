package models

import (
	"fmt"

	"github.com/beego/beego/v2/client/orm"
)

func main() {
	name := "default"
	force := true
	verbose := true

	err := orm.RunSyncdb(name, force, verbose)
	if err != nil {
		fmt.Println(err)
	}
}
