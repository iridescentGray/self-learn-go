# Secret 用途

- Secret 来存储加密信息，虽然在资源文件的编码上，只是通过 Base64 的方式简单编码
- 生产中更推荐专业的 AWS KMS 服务进行密钥管理

# Secret 公共命令

## 创建 Secret 功能

    kubectl apply -f secret.yaml

## 查看

    # 查看指定namespace的Secret
    kubectl get secret

    # 查看全部namespace的Secret
    kubectl get secret --all-namespaces

    # 查看详情
    kubectl describe secret <map_name>

## 使用 Secret

    # Secret的内容会添加到Env中,并且转换为base64解码后的形式
    os.Getenv("Secret中的内容")

## 删除

    # 删除指定的Secret
    kubectl delete secret <Name>
