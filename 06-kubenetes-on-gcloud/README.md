# Setting up HTTP Load Balancing with Ingress

TODO

[x] 1. Setting up HTTP Load Balancing with Ingress

https://cloud.google.com/kubernetes-engine/docs/tutorials/http-balancer

[x] 2. Configuring Domain Names with Static IP Addresses

https://cloud.google.com/kubernetes-engine/docs/tutorials/configuring-domain-name-static-ip

https://cloud.google.com/kubernetes-engine/docs/tutorials/http-balancer

## Step 1: Deploy a web application
Create a Deployment using the sample web application container image that listens on a HTTP server on port 8080:

To create the Deployment, download web-deployment.yaml, then apply the resource to the cluster:

```shell
kubectl apply -f web-deployment.yaml
```

## Step 2: Expose your Deployment as a Service internally
Create a Service resource to make the web deployment reachable within your container cluster.

To create the Deployment, download web-service.yaml, then apply the resource to the cluster:

```shell
kubectl apply -f web-service.yaml
```

When you create a Service of type NodePort with this command, GKE makes your Service available on a randomly- selected high port number (e.g. 32640) on all the nodes in your cluster.

Verify the Service was created and a node port was allocated:

```shell
kubectl get service web
```

Output:

```
NAME      TYPE       CLUSTER-IP      EXTERNAL-IP   PORT(S)          AGE
web       NodePort   10.35.245.219   <none>        8080:32640/TCP   5m
```

In the sample output above, the node port for the web Service is 32640. Also, note that there is no external IP allocated for this Service. Since the GKE nodes are not externally accessible by default, creating this Service does not make your application accessible from the Internet.

To make your HTTP(S) web server application publicly accessible, you need to create an Ingress resource.

## Step 3: Create an Ingress resource
Ingress is a Kubernetes resource that encapsulates a collection of rules and configuration for routing external HTTP(S) traffic to internal services.

On GKE, Ingress is implemented using Cloud Load Balancing. When you create an Ingress in your cluster, GKE creates an HTTP(S) load balancer and configures it to route traffic to your application.

While the Kubernetes Ingress is a beta resource, meaning how you describe the Ingress object is subject to change, the Cloud Load Balancers that GKE provisions to implement the Ingress are production-ready.

The following config file defines an Ingress resource that directs traffic to your web Service:

```shell
apiVersion: extensions/v1beta1
kind: Ingress
metadata:
  name: basic-ingress
spec:
  backend:
    serviceName: web
    servicePort: 8080
```

To deploy this Ingress resource, download basic-ingress.yaml and run:

```shell
kubectl apply -f basic-ingress.yaml
```

Once you deploy this manifest, Kubernetes creates an Ingress resource on your cluster. The Ingress controller running in your cluster is responsible for creating an HTTP(S) Load Balancer to route all external HTTP traffic (on port 80) to the web NodePort Service you exposed.

## Step 4: Visit your application
Find out the external IP address of the load balancer serving your application by running:

```shell
kubectl get ingress basic-ingress
```

Output:

```
NAME            HOSTS     ADDRESS         PORTS     AGE
basic-ingress   *         203.0.113.12    80        2m
```

Note: It may take a few minutes for GKE to allocate an external IP address and set up forwarding rules until the load balancer is ready to serve your application. In the meanwhile, you may get errors such as HTTP 404 or HTTP 500 until the load balancer configuration is propagated across the globe.

Point your browser to the external IP address of your application and see a plain text HTTP response like the following:

```
Hello, world!
Version: 1.0.0
Hostname: web-6498765b79-fq5q5
```

You can visit Load Balancing on Cloud Console and inspect the networking resources created by the Ingress controller.