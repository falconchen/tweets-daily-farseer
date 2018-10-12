package main

import (
	"fmt"
	tf "github.com/falconchen/tweets-daily-farseer"
	"io/ioutil"
	"time"
)

const commentFile = "../data/comments.txt"
const cookieFile = "../data/cookie.txt"

func main() {
	tf.Info.Println("Start working")
	cookieStr, err := ioutil.ReadFile(cookieFile)
	if err != nil {
		tf.Error.Println(err.Error())
		return
	}

	td := tf.New(tf.NewClient(), string(cookieStr))
	for {
		itemRemote, err := td.GetRemote()
		if err != nil {
			tf.Error.Println(err.Error())
		}
		itemLocal, err := td.GetLocal()
		if err != nil {
			tf.Error.Println(err.Error())
			return
		}
		if itemRemote[1] != itemLocal[1] {

			workingUrl, workingId := itemRemote[0], itemRemote[1]
			err := td.UpdateLocal([]byte(workingId))
			if err != nil {
				tf.Error.Println(err.Error())
			}
			tf.Info.Println(fmt.Sprintf("Update: %s => %s ", itemLocal[1], workingId))
			exists, err := td.CheckShafaExists(workingId)
			if err != nil {
				tf.Error.Println(err.Error())
				continue
			}
			if exists {
				tf.Info.Println("Oops, Shafa has gone! ")
				continue
			}
			comment, err := ioutil.ReadFile(commentFile)
			if err != nil {
				tf.Error.Println(err.Error())
				return
			}
			err = td.SendComment(workingId, string(comment))
			if err != nil {
				tf.Error.Println(err.Error())
				//continue
			}
			tf.Info.Println(fmt.Sprintf("Shafa Comment was sent to: %s : %s ", workingUrl, comment))
		}
		time.Sleep(3 * time.Second)
	}
}
