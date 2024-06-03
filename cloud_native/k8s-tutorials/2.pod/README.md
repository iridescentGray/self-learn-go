# Pod

## 创建一个 nginx pod

    # create pod
    kubectl apply -f nginx.yaml

    # 获取当前pod
    kubectl get pods
    // nginx-pod         1/1     Running   0           6s

    # mapping nginx port to local
    kubectl port-forward nginx-pod 4000:80

## 微调

    # 进入 Pod 内容器的 Shell
    kubectl exec -it nginx-pod -- bash

    # 修改nginx 的首页内容
    echo "hello kubernetes by nginx!" > /usr/share/nginx/html/index.html

    # mapping nginx port to local
    kubectl port-forward nginx-pod 4000:80
