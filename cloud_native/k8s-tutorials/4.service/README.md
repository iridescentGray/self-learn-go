# 用途

- Service 为 pod 提供一个稳定的 Endpoint,负责接收请求并将它们传递给它后面的所有 pod
- Service 的 Pod 集合发生更改，请求会导向最新的 pod

## ServiceTypes

Service 类型，默认是 ClusterIP。Type 的值包括如下：

- ClusterIP：通过集群的内部 IP 暴露服务，选择该值时服务只能够在集群内部访问。 这是默认选项
- NodePort：通过每个节点上的 IP 和静态端口（NodePort）暴露服务。 NodePort 服务会路由到自动创建的 ClusterIP 服务。 通过请求 <节点 IP>:<节点端口>，你可以从集群的外部访问一个 NodePort 服务
- LoadBalancer：使用云提供商的负载均衡器向外部暴露服务。 外部负载均衡器可以将流量路由到自动创建的 NodePort 服务和 ClusterIP 服务上
- ExternalName：通过返回 CNAME 和对应值，可以将服务映射到 externalName 字段的内容（例如，foo.bar.example.com）。 无需创建任何类型代理

# Service 公共命令

## 创建 service

    kubectl apply -f service.yaml

## 查看

    # 查看现有Service
     kubectl get service

## 删除

    # 删除指定的service
    kubectl delete service <Name>
    # 删除全部
    kubectl delete service --all
