# Pod

## 创建一个 nginx pod,内部时 nginx-container

    # 创建 pod
    kubectl apply -f nginx.yaml

    # 获取当前pod
    kubectl get pods

    # mapping nginx port to local
    kubectl port-forward nginx-pod 4000:80

## 微调 Pod 内部的 nginx-container

    # 进入 Pod 内容器的 Shell
    kubectl exec -it nginx-pod -- bash

    # 修改nginx 的首页内容
    echo "hello kubernetes by nginx!" > /usr/share/nginx/html/index.html

    # mapping nginx port to local
    kubectl port-forward nginx-pod 4000:80
