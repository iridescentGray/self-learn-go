# Pod

内部管理 Pod,为 Pod 创建多个 replicas

# deployment 公共命令

## 创建 deployment

    kubectl apply -f deployment.yaml

## 查看

    # 查看现有deployment
    kubectl get deployment

    # 查看指定 deployment 的历史版本
    kubectl rollout history deployment <name>

## 回滚

    # 回滚到指定deployment版本
    kubectl rollout undo deployment/hellok8s-deployment --to-revision=2

## 删除

    # 删除指定的deployment
    kubectl delete deployment <Name>
    # 删除全部
    kubectl delete deployment --all
