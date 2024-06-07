# deployment force update

    强制更新是更deployment的Pod的一个不优雅的方法,会直接替换掉运行中的Pod

## 命令

    # 构建v2新镜像
    docker build . -t colorfulgray0/hellok8s:v2
    docker push colorfulgray0/hellok8s:v2

    # 部署
    kubectl apply -f deployment.yaml

    # 查看获取deployment创建的pod
    kubectl get pods

    # 转发端口并查看 http://localhost:3000
    kubectl port-forward hellok8s-deployment-68d996b94-8z4nm 3000:3000
