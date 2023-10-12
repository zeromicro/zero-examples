package main

import (
	"log"
	"time"

	"github.com/zeromicro/go-zero/core/mr"
	"github.com/zeromicro/go-zero/core/timex"
)

type user struct{}

func (u *user) User(uid int64) (interface{}, error) {
	time.Sleep(time.Millisecond * 30)
	return nil, nil
}

type store struct{}

func (s *store) Store(pid int64) (interface{}, error) {
	time.Sleep(time.Millisecond * 50)
	return nil, nil
}

type order struct{}

func (o *order) Order(pid int64) (interface{}, error) {
	time.Sleep(time.Millisecond * 40)
	return nil, nil
}

var (
	userRpc  user
	storeRpc store
	orderRpc order
)

func main() {
	start := timex.Now()
	_, err := productDetail(123, 345)
	if err != nil {
		log.Printf("product detail error: %v", err)
		return
	}
	log.Printf("productDetail time: %v", timex.Since(start))

	// the data processing
	res, err := checkLegal([]int64{1, 2, 3})
	if err != nil {
		log.Printf("check error: %v", err)
		return
	}
	log.Printf("check res: %v", res)
}

type ProductDetail struct {
	User  interface{}
	Store interface{}
	Order interface{}
}

func productDetail(uid, pid int64) (*ProductDetail, error) {
	var pd ProductDetail
	err := mr.Finish(func() (err error) {
		pd.User, err = userRpc.User(uid)
		return
	}, func() (err error) {
		pd.Store, err = storeRpc.Store(pid)
		return
	}, func() (err error) {
		pd.Order, err = orderRpc.Order(pid)
		return
	})
	if err != nil {
		return nil, err
	}

	return &pd, nil
}

func checkLegal(uids []int64) ([]int64, error) {
	r, err := mr.MapReduce(func(source chan<- int64) {
		for _, uid := range uids {
			source <- uid
		}
	}, func(item int64, writer mr.Writer[int64], cancel func(error)) {
		uid := item
		ok, err := check(uid)
		if err != nil {
			cancel(err)
		}
		if ok {
			writer.Write(uid)
		}
	}, func(pipe <-chan int64, writer mr.Writer[[]int64], cancel func(error)) {
		var uids []int64
		for p := range pipe {
			uids = append(uids, p)
		}
		writer.Write(uids)
	})
	if err != nil {
		return nil, err
	}

	return r, nil
}

func check(uid int64) (bool, error) {
	// do something check user legal
	time.Sleep(time.Millisecond * 20)
	return true, nil
}
