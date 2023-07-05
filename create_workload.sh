#!/usr/bin/env bash

export KUBERNETES_VERSION=v1.26.3
export CLUSTER_NAME=kubecontest
export CONTROL_PLANE_MACHINE_COUNT=1
export WORKER_MACHINE_COUNT=1
export CNI_NAME=antrea

kubectl create configmap ${CNI_NAME}-crs-configmap --from-file=${CNI_NAME}.yaml
clusterctl generate cluster kubecontest --from templates/cluster-template.yaml > cluster.yaml
kubectl apply -f cluster.yaml