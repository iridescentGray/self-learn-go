# deployment rolling update

    滚动更新是保证新版本的Pod还没有ready之前，先不删除旧版本的 pod

## 命令

    # 构建v3新镜像
    docker build . -t colorfulgray0/hellok8s:v3
    docker push colorfulgray0/hellok8s:v3


    # 观察
    kubectl get pods --watch

    # 部署
    kubectl apply -f deployment.yaml

    # 转发端口并查看 http://localhost:3000
    kubectl port-forward hellok8s-deployment-68d996b94-8z4nm 3000:3000
