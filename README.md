# kubernetes: linkerd grpc load balancing

## generate models

```bash
cd auth
protoc --go_out=. --go-grpc_out=. proto/*.proto
cd ../user
protoc --go_out=. --go-grpc_out=. proto/*.proto
cd ..
```

## run minikube

```bash
minikube start
eval $(minikube docker-env)
```

## build images

```bash
docker buildx build -t app-auth:auth ./auth --platform linux/amd64
docker buildx build -t app-user:user ./user --platform linux/amd64
```

## install linkerd

```bash
curl --proto '=https' --tlsv1.2 -sSfL https://run.linkerd.io/install | sh
linkerd check --pre
linkerd install | kubectl apply -f -
linkerd check
```

## config kubernetes

```bash
cat kubernetes/auth/deployment.yaml | linkerd inject - | kubectl apply -f -
kubectl apply -f kubernetes/user/deployment.yaml
```
