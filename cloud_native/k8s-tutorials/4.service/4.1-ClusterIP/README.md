# 创建 service cluster ip

## 先 deployment

    # 构建host 新镜像
    docker build . -t colorfulgray0/hellok8s:host
    docker push colorfulgray0/hellok8s:host

    # 创建deployment
    kubectl apply -f deployment.yaml

## 创建 service

    # 创建service,通过 selector: app: hellok8s 与deployment关联
    kubectl apply -f service-hellok8s-clusterip.yaml

    # 查看service,我能能看到service的ip
    # 可以通过在集群其它应用中访问 service-hellok8s-clusterip 的 IP 来访问 hellok8s:host 服务
    kubectl get service

    # 查看endpoints,我们能看到3个pod各有不同的ip
    kubectl get endpoints

    # 查看 pod ,我们能看到它有了ip
    kubectl get pod -o wide

## 创建一个 Nginx,对外暴露

    # 启动nginx
    kubectl apply -f nginx.yaml
    # 进入nginx的shell
    kubectl exec -it nginx -- bash
    # 在 nginx的shell中请求service,每次返回的 hellok8s:host hostname 不一样,说明有自动负载均衡
    curl <service_ip>:3000
