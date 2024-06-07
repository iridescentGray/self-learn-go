# liveness

    存活探针

用途:

- 探活发现应用无回应时,重启容器

## 命令

    # 构建liveness新镜像
    docker build . -t colorfulgray0/hellok8s:liveness
    docker push colorfulgray0/hellok8s:liveness

    # 部署
    kubectl apply -f deployment.yaml

    # 发现 pod 没有READY
    kubectl get pods

    # 20秒后  describe Pod ,发现 pod 没有READY
    kubectl describe pod <Pod>
