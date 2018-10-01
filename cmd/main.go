package main

import (
	"fmt"
	tf "github.com/falconchen/tweets-daily-farseer"
	"github.com/gpmgo/gopm/modules/log"
	"time"
)

func main() {
	td := tf.New(tf.NewClient())
	for {
		itemRemote, err := td.GetRemote()
		if err != nil {
			log.Error(err.Error())
		}
		itemLocal, err := td.GetLocal()
		if err != nil {
			log.Error(err.Error())
		}
		if itemRemote[1] != itemLocal[1] {
			td.UpdateLocal([]byte(itemRemote[1]))
			err := td.SendComment(itemRemote[1], "小小编国庆快乐！ 我来抢沙发了:) ")
			if err != nil {
				log.Error(err.Error())
			}
			log.Info(fmt.Sprint("update: %s change to %s \n", itemLocal[1], itemRemote[1]))
		}
		time.Sleep(time.Second)
	}
}
