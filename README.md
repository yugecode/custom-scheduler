# custom-scheduler
k8s版本为：v1.18.6
```azure
1.拉取k8s源码：
　git clone https://github.com/kubernetes/kubernetes.git
2.切换分支：
　git branch realse-1.18
3.下载k8s依赖：
  go mod tidy
  go mod download
4.替换当前工程的go.mod引用k8s的路径
　　如：k8s.io/api => /home/dev/code/go/src/github.com/kubernetes/staging/src/k8s.io/api
　　后面的/home/dev/code/go/src/github.com/kubernetes就是需要替换的内容，也就是刚才克隆的k8s地址．
```
## 本地调试
```azure
启动需要加参数：
--config=conf/scheduler-config.yaml
--master=http://127.0.0.1:8080 ->api-server的地址
--v=3

配置文件在conf里面
```

## Deploy

```shell
$ kubectl apply -f deploy/custom-scheduler.yaml -n kube-system
```

##　使用调度器
```azure
需要在pod模板指定调度器名称：
schedulerName: custom-scheduler
```
