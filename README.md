####
    go get github.com/freeddser/fileServer/client
    go get github.com/freeddser/fileServer/server
    
#### Server
    ./server.go --port 5050
  
#### Client
    ./client.go --server http://127.0.0.1:5050 --filename a.txt

#### use curl upload file
    curl  -F "file=@/var/log/syslog" "localhost:5050/upload?name=abc"
