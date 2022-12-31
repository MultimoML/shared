## Deploy

```
helm repo add traefik https://helm.traefik.io/traefik
helm repo update
```

```
kubectl create namespace traefik
```

```
kubectl apply -f traefik-config.yaml
```

```
helm install traefik traefik/traefik --namespace=traefik --values=traefik-chart-values.yaml --set providers.kubernetescrd.allowCrossNamespace=true 
```