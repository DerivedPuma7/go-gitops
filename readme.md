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