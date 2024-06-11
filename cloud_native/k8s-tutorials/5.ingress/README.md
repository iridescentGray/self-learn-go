# Ingress 用途

- Ingress 为 Service 提供外部可访问的 URL、负载均衡流量、 SSL/TLS，以及基于名称的虚拟托管。
- Ingress 可以“简单理解”为服务的网关 Gateway，它是所有流量的入口、路由
- 必须拥有一个 Ingress 控制器,仅创建 Ingress 资源本身没有任何效果
- minikube 默认使用的是 nginx-ingress 控制器 ，目前 minikube 也支持 Kong-Ingress 控制器

# Ingress 公共命令

## 启用 Ingress 功能

    minikube addons enable ingress

## 查看

    # 查看现有Ingress
    kubectl get ingress

## 删除

    # 删除指定的Ingress
    kubectl delete ingress <Name>
