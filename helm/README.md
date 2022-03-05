# Helm charts and configs

## Preparation

I used the following repositories:
* grafana for Grafana (grafana chart)
* prometheus-community for Prometheus (prometheus chart)
* bitnami for Thanos (thanos chart)

Here is how to set them up:

```
helm repo add grafana https://grafana.github.io/helm-charts
helm repo add prometheus-community https://prometheus-community.github.io/helm-charts
helm repo add bitnami https://charts.bitnami.com/bitnami
helm repo update
```

In addition, to help exploring setting custom values, I extracted values this way:

```
helm show values grafana/grafana > values_definitions/grafana-values-all.yml
helm show values prometheus-community/prometheus > values_definitions/prometheus-values-all.yml
helm show values bitnami/thanos > values_definitions/thanos-values-all.yml
# minio is a dependency of thanos
helm show values bitnami/minio > values_definitions/minio-values-all.yml
```

I also wanted to explore charts, so I cloned them locally:

```
git clone git@github.com:grafana/helm-charts.git helm-grafana
git clone git@github.com:prometheus-community/helm-charts.git helm-prometheus-community
git clone git@github.com:bitnami/charts.git helm-bitnami
```


## Deployments

### gotestmetrics

You should adjust values:
* grafana.mycluster.io with your own cluster name

Then run the chart with no specific config:

```
helm install gotestmetrics gotestmetrics
```



### Grafana

You should adjust values:
* grafana.mycluster.io with your own cluster name
* adminPassword with a secure password

Ten it should just work this way:

```
helm install grafana grafana/grafana -f grafana-values.yml
```


### Prometheus

It's a bit more complicated here, as it seems the Helm charts can't provide all I needed for Thanos support.

```
# First of all, create a service account
kubectl apply -f prometheus-serviceaccount.yml

# Then deploy Prometheus
helm install prometheus prometheus-community/prometheus --values prometheus-values.yml

# patch it for full Thanos sidecar setup
kubectl patch statefulset prometheus-server --patch-file prometheus-server-patch-statefulset.yml
kubectl patch service prometheus-server --patch-file prometheus-server-patch-service.yml

# And finall, edit its config to add external_labels required for Thanos
kubectl edit configmap prometheus-server
# add in global section:
#       external_labels:
#         cluster: prometheus
```

### Thanos

Nothing special here, all the specific config is in the values file.

```
helm install thanos bitnami/thanos --values thanos-values.yml
```

## Debugging

For DNS tests, I deployed a `dnsutils` pod:
```
kubectl apply -f dnsutils.yaml
```

Then you can try DNS queries this way:

```
kubectl exec -ti dnsutils -- nslookup _grpc._tcp.prometheus-server.default.svc.eu-central-1.local
```

## Aborted solutions

### kube-prometheus-stack:

I tried to depoly the Prometheus operator, which seems to support Thanos sidecar, but failed (probably because of missing rights for K8S control)

```
helm install prometheus prometheus-community/kube-prometheus-stack --values prometheus-stack-values.yml
```
