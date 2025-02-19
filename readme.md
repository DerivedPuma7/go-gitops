# GitOps

## Docker
### build image
```console
docker build -t derivedpuma7/gitops:version .
```

### push to docker hub
```console
docker push derivedpuma7/gitops:version
```

### run container
```console
docker run --rm -p 8080:8080 derivedpuma7/gitops:version
```

## Kind
### create kind cluster
```console
kind create cluster --name=gitops
```

### change context
```console
kubectl cluster-info --context kind-gitops
```

## ArgoCD
### Apply argo config to k8s pods
```console
kubectl apply -n argocd -f https://raw.githubusercontent.com/argoproj/argo-cd/stable/manifests/install.yaml
```

### Get argocd first login password
```console
argocd admin initial-password -n argocd
```

### Argocd port-forwarding
```console
kubectl port-forward svc/argocd-server -n argocd 8080:443
```