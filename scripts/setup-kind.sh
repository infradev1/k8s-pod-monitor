#!/bin/bash
set -e

CLUSTER_NAME="pod-monitor"

echo "Creating KinD cluster..."
kind create cluster --name "$CLUSTER_NAME" --config manifests/kind-cluster.yaml

echo "Waiting for nodes to be ready..."
kubectl wait --for=condition=Ready nodes --all --timeout=60s

echo "Deploying test pod with restart loop..."
kubectl apply -f manifests/crashloop-pod.yaml

echo "Done. Cluster '$CLUSTER_NAME' is ready."