package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/zeromicro/go-zero/rest/httpc"
	"github.com/zeromicro/zero-examples/http/httpc/types"
)

var domain = flag.String("domain", "http://localhost:3333", "the domain to request")

func main() {
	flag.Parse()

	req := types.Request{
		Node:   "foo",
		ID:     1024,
		Header: "foo-header",
		Body:   "hello, world",
	}
	resp, err := httpc.Do(context.Background(), http.MethodPost, *domain+"/nodes/:node", req)
	if err != nil {
		fmt.Println(err)
		return
	}

	io.Copy(os.Stdout, resp.Body)
}
