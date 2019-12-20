# Lab 03: Setup Kubernetes Cluster locally

Trong hướng dẫn này, chúng ta sẽ thiết lập Cụm Kubernetes trên máy tính cục bộ sử dụng công cụ kubeadm.

TODO

1. [x] Tạo 3 máy ảo Virtual box làm 3 nodes: kube-master, kube-node1, kube-node2
2. [x] Cài đặt Docker lên tất cả nodes.
3. [x] Cài đặt kubeadm, kubelet và kubectl lên tất cả nodes.
4. [x] Tạo một control-plane cluster with kubeadm

## 1. Giới thiệu

Kubeadm là một công cụ giúp tự động hóa quá trình cài đặt và triển khai các cụm kubernetes trên môi trường Linux, do chính kubernetes hỗ trợ.

kubeadm hỗ trợ các nền tảng là Ubuntu 16.04+, Debian 9+, CentOS 7, Fedora 25+, HypriotOS v1.0.1+.

## 2. Mô hình

Trong hướng dẫn này, chúng ta sử dụng phiên bản kubeadm 1.6 và chạy 3 máy ảo Ubuntu 16.04 với Virtualbox.

[]()

Cụm cluster có 3 nodes: 1 master và 2 worker nodes.


- kubemaster: 192.168.56.2
- kubenode1: 192.168.56.2
- kubenode2: 192.168.56.2

### Bước 1: Tạo 3 máy ảo Virtual box

1. kube-master
2. kube-node1
3. kube-node2

### Bước 2: Config Unique hostname, MAC address, and product_uuid for every node.

1. Open Virtual Box and create Global IP address: 192.168.56.1/24 and unchecked DHCP Server.

2. Open kube-master machine config a static IP use this command:

```shell
# Find Enthernet name
ifconfig -a

# Config static IP
sudo ifconfig enp0s8 192.168.56.2

# Add configure for enp0s8
sudo nano /etc/network/interfaces

# file: /etc/network/interfaces
# iface enp0s8 inet static
#     address 192.168.56.2
#     netmask 255.255.255.0
```

Update hostname and hosts file

```shell
sudo nano /etc/hostname : kubemaster
sudo nano /etc/hosts : 127.0.0.1    kubemaster
```

## 3. Triển khai

### 3.1 Cài đặt 

Đầu tiên chúng ta phải cài đặt các packages sau đây trên tất cả các node.

- docker
- kubelet
- kubectl
- kubeadm

### Bước 1: Cài đặt Docker CE theo hướng dẫn tại trang chủ Docker

https://docs.docker.com/install/linux/docker-ce/ubuntu/

https://kubernetes.io/docs/setup/production-environment/container-runtimes/

### Bước 2: Cài đặt kubeadm, kubelet và kubectl

```shell
sudo apt-get update && sudo apt-get install -y apt-transport-https curl
curl -s https://packages.cloud.google.com/apt/doc/apt-key.gpg | sudo apt-key add -
cat <<EOF | sudo tee /etc/apt/sources.list.d/kubernetes.list
deb https://apt.kubernetes.io/ kubernetes-xenial main
EOF
sudo apt-get update
sudo apt-get install -y kubelet kubeadm kubectl
sudo apt-mark hold kubelet kubeadm kubectl
```

### 3.2 Khởi tạo master.

On kube-master run this command:

```shell
sudo kubeadm init --pod-network-cidr=10.244.0.0/16 --apiserver-advertise-address=192.168.56.2
```

Installing a pod network add-on

```shell
kubectl apply -f https://raw.githubusercontent.com/coreos/flannel/2140ac876ef134e0ed5af15c65e414cf26827915/Documentation/kube-flannel.yml
```

Check pods

```shell
kubectl get pods --all-namespaces

# NAMESPACE     NAME                                 READY   STATUS    RESTARTS   AGE
# kube-system   coredns-6955765f44-jg64c             1/1     Running   0          10m
# kube-system   coredns-6955765f44-lspqr             1/1     Running   0          10m
# kube-system   etcd-kubemaster                      1/1     Running   0          10m
# kube-system   kube-apiserver-kubemaster            1/1     Running   0          10m
# kube-system   kube-controller-manager-kubemaster   1/1     Running   1          10m
# kube-system   kube-flannel-ds-amd64-lx5jb          1/1     Running   0          9m12s
# kube-system   kube-proxy-6llzp                     1/1     Running   0          10m
# kube-system   kube-scheduler-kubemaster            1/1     Running   1          10m
```

### 3.3 Join in a worker nodes to master

Then you can join any number of worker nodes by running the following on each as root:

```shell
kubeadm join 192.168.56.2:6443 --token 97rh8y.5bka3bvcqaxf5hgi \
    --discovery-token-ca-cert-hash sha256:41c2d74fc1aa657b436c17a606aed775a4118dede868fca0c34f4323fe2fe415
```

Open kube-master and check joined nodes:

```shell
kubectl get nodes

# Result
# NAME         STATUS     ROLES    AGE   VERSION
# kubemaster   Ready      master   18m   v1.17.0
# kubenode1    NotReady   &lt;none&gt;   30s   v1.17.0</pre>
```

### 3.4 Kubeadm reset

To rest kubeadm run this command:

```shell
sudo kubeadm reset -f
```






