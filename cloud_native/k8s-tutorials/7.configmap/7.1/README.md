# 使用 config map

## namespace

    假设以下两个namespace已经存在
    namespace/dev
    namespace/test

## 构建镜像

    # 构建host 新镜像
    docker build . -t colorfulgray0/hellok8s:configmap
    docker push colorfulgray0/hellok8s:configmap

## 创建 ConfigMap

    kubectl apply -f hellok8s-config-dev.yaml -n dev
    kubectl apply -f hellok8s-config-test.yaml -n test
