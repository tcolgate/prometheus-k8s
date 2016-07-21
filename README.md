# Kubernetes configs to setup prometheus

Initially aimed at minikube, but should be adaptable
for real clusters.

- Adds node_exporter DaemonSet to each node
- Creates a persistant volume and claim for prometheus data
- Creates a prometheus instance
- Creates an alertmanager instance.
- Prometheus will monitor all elements of the kube cluster (including
  services)
- A grafana deployment (no state is kept yet, and no default
  data source setup)
- A templated grafana dashboard (not loaded by default yet)

TODO:

- Get grafana state on the PV
- Prometheus rules in a config map
- Reload prometheus config on rule update
- HA Prom (should just mean changing the rep count)
- HA Alertmanager
- Alertmanager config

License:

These configs are released under the Apache 2.0 license. All images
downloaded are subject to their individual licenses.
