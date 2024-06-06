# deployment

## 定义 deployment.yaml

    metadata.name 是deploy的名字，必须唯一
    replicas 表示pod 副本数量
    selector 表示的是 deployment与pod 的关联方式， matchLabels:hellok8s表示 deployment 会管理 (selector) 所有 labels=hellok8s 的 pod
    template 用于定义 pod 资源

## 创建 deployment

    # 创建deployment
    kubectl apply -f deployment.yaml

    # 获取deployment
    kubectl get deployments

    # 查看获取deployment创建的pod
    kubectl get pods

    # 转发端口并查看 http://localhost:3000
    kubectl port-forward hellok8s-deployment-6599d47755-9j5d9 3000:3000

## 尝试删除由 deployment 创建的 pod，发现删除后会重新创建

    # 打开一个终端观察pod的变化
    kubectl get pods --watch

    # 删除 deployment 创建的 pod
    kubectl delete pod <pod_name>

    # 发现删除后会重新创建
    kubectl get pods
