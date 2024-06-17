# 使用 secret

## 创建 secret

    # base64加密,填充到 hellok8s-config.yaml中
    echo "db_password" | base64
    # ZGJfcGFzc3dvcmQK

    echo "ZGJfcGFzc3dvcmQK" | base64 -d
    # db_password

    kubectl apply -f hellok8s-config.yaml

## 构建镜像

    # 构建host 新镜像
    docker build . -t colorfulgray0/hellok8s:secret
    docker push colorfulgray0/hellok8s:secret

## 创建对应 namespace 下的 pod

    kubectl apply -f hellok8s-pod.yaml

## 发送请求测试 secret 效果

    # 测试
    kubectl port-forward hellok8s-pod 3000:3000
    curl http://localhost:3000
