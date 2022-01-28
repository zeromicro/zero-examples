## go-zero 使用jwt例子

#### 使用
##### 运行
```shell
make run
```

##### 获取token
```shell
curl --location --request POST '127.0.0.1:8888/user/token'
{"access_token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2Mjc0MzU4NTUsImlhdCI6MTYyNjgzMTA1NX0.gjPdqmGp5waFVK87zgHMMYEorq614oSdNUkjlFlYx94","access_expire":1627435855,"refresh_after":1627133455}
```
##### 请求接口
```shell
curl -w  "\nhttp: %{http_code} \n" --location --request POST '127.0.0.1:8888/user/info' \
--header 'Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2Mjc0MzU4NTUsImlhdCI6MTYyNjgzMTA1NX0.gjPdqmGp5waFVK87zgHMMYEorq614oSdNUkjlFlYx94' \
--header 'Content-Type: application/json' \
--data-raw '{
    "userId": "a"
}'
```

##### Note Before Run 
```
在运行Demo前，需要修改"/etc/jwt-api.yaml"中"AccessSecret"字段；否则访问"/user/info"接口出现 "401:Token is expired "
```
