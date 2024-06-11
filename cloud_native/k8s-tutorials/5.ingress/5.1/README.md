# 创建 ingress

## 先 deployment

    # 启动业务
    kubectl apply -f hellok8s.yaml
    kubectl apply -f nginx.yaml

    # 启动ingress
    kubectl apply -f ingress.yaml

    # 查看ingress
    kubectl get ingress

    # 如果是docker-desktop 需要手动暴露服务
    minikube service ingress-nginx-controller -n ingress-nginx --url

    #测试ingress的效果
    curl http://<ip>/hello
    curl http://<ip>/
