# Pod

## 创建一个 nginx pod,内部时 nginx-container

    # create pod
    kubectl apply -f nginx.yaml

    # 获取当前pod
    kubectl get pods
    // nginx-pod         1/1     Running   0           6s

    # mapping nginx port to local
    kubectl port-forward nginx-pod 4000:80

## 微调 Pod 内部的 nginx-container

    # 进入 Pod 内容器的 Shell
    kubectl exec -it nginx-pod -- bash

    # 修改nginx 的首页内容
    echo "hello kubernetes by nginx!" > /usr/share/nginx/html/index.html

    # mapping nginx port to local
    kubectl port-forward nginx-pod 4000:80

## Pod 相关的其他命令

### 日志

    # 跟踪Pod日志 (类似tail -f)
    kubectl logs --follow nginx-pod

### 在 Pod 内执行命令

    kubectl exec nginx-pod -- ls

### 删除 Pod

    kubectl delete pod nginx-pod
    kubectl delete -f nginx.yaml
