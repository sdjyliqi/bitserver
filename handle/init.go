package handle

import (
	"sync"
	"tgin/conf"
	"tgin/control"
)

var once         sync.Once
var UserCenter    control.UserCenter

func init(){
	once.Do(func() {
		UserCenter = control.CreateUserCenter(&conf.DefaultConfig)
	})
}

