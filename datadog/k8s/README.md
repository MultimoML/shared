## Deploy

```
helm repo add datadog https://helm.datadoghq.com
helm repo update
```

```
kubectl apply -f datadog-secrets.yaml

helm install datadog datadog/datadog -f datadog-agent-chart-values.yaml
```
