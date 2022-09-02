 #!/bin/bash
date
{
    cmd="go run hello.go"
    eval ${cmd}
} &
sleep 2
{
    cmd="go run gateway.go"
    eval ${cmd}
} &
sleep 2
{
    cmd="go run client.go"
     eval ${cmd}
} &
  
 wait
 date 

