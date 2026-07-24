# Kubernetes deployment

These manifests target the local Kubernetes cluster provided by Docker Desktop
and follow the same layout as `Bank-exchange-rate-service`.

## Prerequisites

- The main backend deployment must create the shared `bank` namespace.
- Build the application image as `bank-repository-service:local`.
- The main backend deployment must provide Kafka through a Service named
  `kafka` on port `9092`.
- Enable Kubernetes Metrics Server for the HPA.
- Replace the local passwords in `01-config.yaml` outside development.

## Deploy

```powershell
kubectl apply -f docker/kubernetes
kubectl get pods -n bank
kubectl get services -n bank
kubectl get hpa -n bank
```

The gRPC endpoint used by the backend is:

```text
bank-repository.bank.svc.cluster.local:50052
```

PostgreSQL and Redis are internal-only services. The repository deployment
connects to the main backend Kafka at:

```text
kafka.bank.svc.cluster.local:9092
```

On startup, every repository replica idempotently ensures the configured
`KAFKA_TOPICS` exist in that shared broker.
