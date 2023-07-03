#!/usr/bin/env bash

export KUBERNETES_VERSION=v1.26.0
export CLUSTER_NAME=kubecontest
export CONTROL_PLANE_MACHINE_COUNT=1
export WORKER_MACHINE_COUNT=1

kubectl create configmap antrea-crs-configmap --from-file=antrea.yaml
clusterctl generate cluster kubecontest --from templates/cluster-template.yaml > cluster.yaml
kubectl apply -f cluster.yaml