# hello_demo

    # 初始化依赖项
    go mod tidy

    #启动应用程序
    go run main.go serve


    # 打包成可执行程序
    CGO_ENABLED=0 go build
    # 运行打包后的程序
    ./myapp serve
