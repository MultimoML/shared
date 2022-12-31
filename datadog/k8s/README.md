## Deploy

```
helm repo add datadog https://helm.datadoghq.com
helm repo update
```

```
helm install datadog -f datadog-agent-chart-values.yaml  --set datadog.apiKey=<DATADOG_API_KEY> datadog/datadog
```

