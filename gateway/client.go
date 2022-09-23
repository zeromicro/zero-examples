package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/zeromicro/go-zero/rest/httpc"
)

var domain = flag.String("domain", "http://localhost:8888", "the domain to request")

func main() {
	flag.Parse()

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()
	for t := range ticker.C {
		resp, err := httpc.Do(context.Background(), http.MethodGet, *domain+"/pingHello/"+t.Format("2006-01-02 15:04:05"), nil)
		if err != nil {
			fmt.Println(err)
			return
		}
		io.Copy(os.Stdout, resp.Body)
	}

}
