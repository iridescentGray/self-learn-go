# Job 的使用

## 普通 Job

    # 启动Job
    kubectl apply -f simple-job.yaml

    # 查看Job
    kubectl get jobs

    # 查看Job创建Pod
    kubectl get pods

    # 查看Job的执行日志
    hello-job-7qg5h

## Corn Job

    kubectl apply -f corn-job.yaml

    kubectl get cronjob

    kubectl get pods
