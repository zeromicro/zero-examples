### 1. N/A

1. route definition

- Url: /shorten
- Method: GET
- Request: `shortenReq`
- Response: `shortenResp`

2. request definition



```golang
type ShortenReq struct {
	Url string `form:"url"`
}
```


3. response definition



```golang
type ShortenResp struct {
	Shorten string `json:"shorten"`
}
```

### 2. N/A

1. route definition

- Url: /expand
- Method: GET
- Request: `expandReq`
- Response: `expandResp`

2. request definition



```golang
type ExpandReq struct {
	Shorten string `form:"shorten"`
}
```


3. response definition



```golang
type ExpandResp struct {
	Url string `json:"url"`
}
```

