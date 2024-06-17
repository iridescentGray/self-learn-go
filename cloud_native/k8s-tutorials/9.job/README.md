# Job 用途

- Job 即一次性任务。无需一直运行
- Job 会创建一/多个 Pod，并将继续重试 Pod 的执行，直到指定数量的 Pod 成功终止
  - 删除 Job 的操作会清除所创建的全部 Pod
  - 挂起 Job 的操作会删除 Job 的所有活跃 Pod，直到 Job 被再次恢复执行

# Helm 公共命令

## 查看

    # 查看现有Job
    kubectl get jobs

## 删除

    # 删除指定的Job
    kubectl delete job
