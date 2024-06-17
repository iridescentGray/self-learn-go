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

## 创建对应 namespace 下的 pod

    kubectl apply -f hellok8s-pod.yaml -n dev
    kubectl apply -f hellok8s-pod.yaml -n test

## 发送请求测试 config 效果

    # 测试Dev
    kubectl port-forward hellok8s-pod 3000:3000 -n dev
    curl http://localhost:3000

    # 测试TEST
    kubectl port-forward hellok8s-pod 3000:3000 -n test
    curl http://localhost:3000
