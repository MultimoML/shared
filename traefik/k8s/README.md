## Deploy

```
helm repo add traefik https://helm.traefik.io/traefik
helm repo update
```

```
kubectl cluster-info --kubeconfig=/etc/multimo/prod-kubeconfig.yaml

chmod 700 /etc/multimo/prod-kubeconfig.yaml
export KUBECONFIG=/etc/multimo/prod-kubeconfig.yaml
```

```
kubectl apply -f traefik-secrets.yaml
```

```
helm install traefik traefik/traefik --values=traefik-chart-values.yaml

helm upgrade traefik traefik/traefik --values=traefik-chart-values.yaml
```

```
kubectl apply -f traefik-dashboard-auth.yaml
kubectl apply -f traefik-dashboard-ingressroute.yaml
```

```
helm uninstall traefik
```

Made with https://traefik.io/blog/install-and-configure-traefik-with-helm/ and
https://github.com/traefik/traefik-helm-chart/blob/master/traefik/values.yaml.