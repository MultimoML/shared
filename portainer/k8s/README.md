## Deploy

```
helm repo add portainer https://portainer.github.io/k8s/
helm repo update
```


```
kubectl create namespace portainer
```

```
helm upgrade -i -n portainer portainer portainer/portainer --set service.type=ClusterIP ingress.ingressClassName=traefik
```