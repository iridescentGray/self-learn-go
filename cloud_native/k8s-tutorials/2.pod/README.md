# Pod

内部管理多个 container

# Pod 公共命令

## 创建 deployment

    kubectl apply -f pod.yaml

## 查看

    # 获取所有pod
    kubectl get pods

    # 查看 pod 详情
    kubectl describe pod <pod name>

    # 跟踪Pod日志 (类似tail -f)
    kubectl logs --follow <PodName>

## 操作 Pod

    #在 Pod 内执行命令
    kubectl exec <PodName> -- ls

    # 打开pod的shell
    kubectl exec -it <PodName> -- bash

    # 转发Pod的端口
    kubectl port-forward <PodName> 4000:80

## 删除

    # 删除单个
    kubectl delete pod <PodName>
    # 删除-根据文件
    kubectl delete -f nginx.yaml
    # 删除全部
    kubectl delete pod --all
