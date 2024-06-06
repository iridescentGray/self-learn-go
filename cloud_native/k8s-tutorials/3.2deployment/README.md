# deployment-2

## 新

## Rolling Update(滚动更新)

    # 构建2新镜像
    docker build . -t colorfulgray0/hellok8s:v2

    # 推送到远程
    docker push colorfulgray0/hellok8s:v2

    # 查看获取deployment创建的pod
    kubectl get pods

    # 转发端口并查看 http://localhost:3000
    kubectl port-forward hellok8s-deployment-68d996b94-8z4nm 3000:3000
