# kubernetes: linkerd grpc load balancing

## ArgoCD

```
helm template argo/chart --set name=auth --set path=auth > apps/auth.yaml
helm template argo/chart --set name=auth --set path=kubernetes/auth > apps/auth.yaml
helm template argo/chart --set name=user --set path=kubernetes/user > apps/user.yaml
```

## NEW ISTIO

```bash
kubectl apply -f kubernetes/auth/deployment.yaml
kubectl apply -f kubernetes/auth/istio.yaml
kubectl apply -f kubernetes/user/deployment.yaml
kubectl apply -f kubernetes/user/istio.yaml
```

## OLD

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
