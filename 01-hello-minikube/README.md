# Hello Minikube Demo

## Build application

```shell
docker build -t huyquanvu1/hello-minikube .
```

## Run it

```shell
# Start Minukube
minikube start

# Clear hello-minikube-node deployments, services
kubectl delete deployment hello-minikube-node
kubectl delete service hello-minikube-node

# Create deployment and expose service
kubectl create deployment hello-minikube-node --image=huyquanvu1/hello-minikube
kubectl get deployments

# Exposes deployment as a service
kubectl expose deployment hello-minikube-node --type=LoadBalancer --port=8080
kubectl get services

# Local test
minikube service hello-minikube-node
```
