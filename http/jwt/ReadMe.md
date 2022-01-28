## Using JWT in go-zero

##### Run
```shell
make run
```

##### Get token
```shell
curl --location --request POST '127.0.0.1:8888/user/token'
{"access_token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2Mjc0MzU4NTUsImlhdCI6MTYyNjgzMTA1NX0.gjPdqmGp5waFVK87zgHMMYEorq614oSdNUkjlFlYx94","access_expire":1627435855,"refresh_after":1627133455}
```
##### Request the API
```shell
curl -w  "\nhttp: %{http_code} \n" --location --request POST '127.0.0.1:8888/user/info' \
--header 'Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2Mjc0MzU4NTUsImlhdCI6MTYyNjgzMTA1NX0.gjPdqmGp5waFVK87zgHMMYEorq614oSdNUkjlFlYx94' \
--header 'Content-Type: application/json' \
--data-raw '{
    "userId": "a"
}'
```

##### Notice
Before running this demo, you need to modify `AccessSecret` in `/etc/jwt-api.yaml`, otherwise you'll get `401: Token is expired` on accessing `/user/info`
