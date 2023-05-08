# go-price-fetcher

A micro service witten in Golang with docker and kubernetes

## Start the application

```
go run main.go
// or
make run
```

## Start the application with docker

```
docker build -t go-app-price-fetcher:latest .
docker image ls
docker run -d -p 3000:3000 --name priceFetcher go-app-price-fetcher:latest
docker ps
docker logs -f priceFetcher

// remove a running Docker container
docker rm -f priceFetcher

// remove a image
docker rmi go-app-price-fetcher:latest
docker rmi {IMAGE_ID}
```

## Start the application with docker-compose

```
docker-compose build
docker-compose up -d
docker ps
 
docker-compose down
```

## Start the application with kubernetes (minikube)

```
# switch docker environment to use the Docker daemon inside Minikube first
eval $(minikube docker-env)

# build image
docker-compose build
docker image ls

# apply Kubernetes resource configurations 
kubectl apply -f kubernetes/

# some commands for verifing the deployments
kubectl get nodes,pod,svc
kubectl describe pod {pod name}
kubectl describe nodes {node name}


# check the logs
kubectl logs -f {pod name}

# test the service within a Kubernetes cluster
kubectl run -it curl --rm --image=nginx:alpine sh
curl web-service:3000/health
exit

# access a specific service with minikube
minikube service list
minikube service {name}

# access aspecific service without minikube
# check the service name
kubectl get pod,svc
NAME                                  READY   STATUS    RESTARTS   AGE
pod/web-deployment-69b8b6bfbf-5zpfp   1/1     Running   0          84m
pod/web-deployment-69b8b6bfbf-h7mpv   1/1     Running   0          84m
pod/web-deployment-69b8b6bfbf-hfzkg   1/1     Running   0          84m

NAME                  TYPE        CLUSTER-IP       EXTERNAL-IP   PORT(S)          AGE
service/kubernetes    ClusterIP   10.96.0.1        <none>        443/TCP          2d10h
service/web-service   NodePort    10.102.163.218   <none>        3000:30295/TCP   2d10h
# 3000 is the target port: This is the port number on which your application inside the pod is listening.
# 30295 is the node port: NodePort is a port number that is exposed on all the nodes of the Kubernetes cluster. 
# It allows external traffic to reach the service. 
# In this case, the node port is 30295, which means that any traffic coming to the nodes on port 30295 will be directed to the service.

# use the kubectl port-forward command to forward a local port to the service's port
kubectl port-forward service/web-service 3000:3000
Forwarding from 127.0.0.1:3000 -> 3000
Forwarding from [::1]:3000 -> 3000
Handling connection for 3000
# keep the terminal running
# we can access the service by localhost:3000

# If the service has an external IP address (not <none>), you can directly access it using the external IP and port combination.
http://<external-ip>:<port>

# delete all pods
kubectl delete deploy --all
```