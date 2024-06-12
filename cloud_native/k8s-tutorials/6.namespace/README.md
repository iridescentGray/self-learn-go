# Namespace 用途

- k8s 能不能在不同环境 dev test uat prod 中区分资源，让不同环境的资源独立互相不影响，dev 环境给开发使用，test 环境给 QA 使用
- Kubernetes 的 Namespace 提供一种机制，将同一集群中的资源划分为相互隔离的组。 同一名字空间内的资源名称要唯一，跨名字空间时没有这个要求
- 默认使用的 namespace 是 default
-

# Namespace 公共命令

## 创建 Namespace 功能

    kubectl apply -f namespaces.yaml

## 查看

    # 查看现有Namespace
    kubectl get namespaces

## 使用 namespace/ 在 namespace 下创建资源/获取资源

- 在其他命令后面加上 -n namespace 例如:
  - kubectl apply -f deployment.yaml -n dev
  - kubectl get pods -n dev

## 删除

    # 删除指定的Namespace
    kubectl delete namespace <Name>
