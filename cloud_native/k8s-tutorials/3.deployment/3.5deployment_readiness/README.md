# readiness

就绪探针 (readiness)

用途:

- 就绪探针控制哪个 Pod 能够作为 Service 的后端，若未就绪，会被从负载均衡器中剔除,不接受请求

## 命令

    # 构建readiness 新镜像
    docker build . -t colorfulgray0/hellok8s:noready
    docker push colorfulgray0/hellok8s:noready

    # 部署
    kubectl apply -f deployment.yaml

    # 部署
    kubectl get pods
