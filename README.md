## k3s autoscaler

> 通过判断pod pending状态来弹性扩容添加竞价实例节点

## 安装

```bash
kubectl apply -f https://raw.githubusercontent.com/ysicing/k3s-autoscaler/master/hack/deploy/crd.yaml
kubectl apply -f https://raw.githubusercontent.com/ysicing/k3s-autoscaler/master/hack/deploy/k3s-autoscaler.yaml
```

## Roadmap

- v0
  - 实现扩容