# ConfigMap 用途

- ConfigMap 来将你在不同 namespace 环境 (dev/prod)中保存配置
- ConfigMap 中保存的数据不可超 1 MiB。超出此尺寸限制需要考虑挂载存储卷。

# ConfigMap 公共命令

## 创建 ConfigMap 功能

    kubectl apply -f configmap.yaml

## 查看

    # 查看指定namespace的ConfigMap
    kubectl get configmap -n dev

    # 查看全部namespace的ConfigMap
    kubectl get configmap --all-namespaces

## 使用 ConfigMap

- kubectl apply -f deployment.yaml -n dev
- kubectl get pods -n dev

## 删除

    # 删除指定的ConfigMap
    kubectl delete ConfigMap <Name>
