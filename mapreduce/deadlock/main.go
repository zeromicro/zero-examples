package main

import (
	"log"
	"strconv"

	"github.com/zeromicro/go-zero/core/mr"
)

type User struct {
	Uid  int
	Name string
}

func main() {
	uids := []int{111, 222, 333}
	res, err := mr.MapReduce(func(source chan<- int) {
		for _, uid := range uids {
			source <- uid
		}
	}, func(item int, writer mr.Writer[*User], cancel func(error)) {
		uid := item
		user := &User{
			Uid:  uid,
			Name: strconv.Itoa(uid),
		}
		writer.Write(user)
	}, func(pipe <-chan *User, writer mr.Writer[any], cancel func(error)) {
		// missing writer.Write(...), should not panic
	})
	if err != nil {
		log.Print(err)
		return
	}
	log.Print(len(res.([]*User)))
}
