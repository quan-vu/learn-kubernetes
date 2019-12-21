# POD Definition

Learn how a pod is definition in Kubernetes. 

![kubernetes-pod-definition](kubernetes-pod-definition.png "Kubernetes pod definition")

**Kubernetes create the pod**

```shell
kubectl create -f pod-definition.yml
```

**Kubernetes get list pods**

```shell
kubectl get pods
```

**Kubernetes get information of a pods**

```shell
kubectl describe pod myapp-pod
```

Easy to write YAML file with "redhat.vscode-yaml" extension which is a YAML Language Support by Red Hat, with built-in Kubernetes and Kedge syntax. 

## Examples

```
apiVersion:
kind:
metadata:


spec:
```
