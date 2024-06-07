# readiness

就绪探针 (readiness)

用途:

- 知道容器何时准备好接受请求流量，当一个 Pod 内的所有容器都就绪时，才能认为该 Pod 就绪
- 控制哪个 Pod 作为 Service 的后端，若尚未就绪，会被从负载均衡器中剔除

## 命令

    # 构建readiness 新镜像
    docker build . -t colorfulgray0/hellok8s:noready
    docker push colorfulgray0/hellok8s:noready

    # 部署
    kubectl apply -f deployment.yaml

    # 部署
    kubectl get pods
