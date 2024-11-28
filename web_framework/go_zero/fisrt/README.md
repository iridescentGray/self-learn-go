# Go ZERO Hello Demo



# Create Project Commond

    go install github.com/zeromicro/go-zero/tools/goctl@latest
    goctl env check --install --verbose --force
    goctl api new fisrt


## Start Up
    cd fisrt
    go mod tidy
    go run demo.go

### Test 
    curl --request GET 'http://127.0.0.1:8888/from/me'
