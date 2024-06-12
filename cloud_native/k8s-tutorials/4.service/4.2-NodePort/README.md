# 创建 service nodePort

将 hellok8s:host 的 3000 端口映射到 k8s Node 机器的 30000 端口
可以通过访问 http://node1-ip:30000 或者 http://node2-ip:30000 访问到服务。

## 创建 service

    # deployment和Docker镜像,服用4.1节的

    # 创建service,通过 selector: app: hellok8s 与deployment关联
    kubectl delete  service  service-hellok8s-clusterip
    kubectl apply -f service-hellok8s-nodeport.yaml

    # 因为是通过minikube部署,所以我们就用minikube的ip
    minikube ip
    # 若minikube运行在Docker Desktop时,则无法根据ip上述命令获取其到的ip,因为 docker 部分网络限制导致无法用ip直连
    minikube service service-hellok8s-nodeport --url

    # 发送请求
    curl http://192.168.49.2:30000
