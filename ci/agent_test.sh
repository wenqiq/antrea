#!/bin/bash

pod_counts=(5 10 20 40 80 100)

namespace="test-pod-namespace"

kubectl create namespace $namespace

for count in "${pod_counts[@]}"; do
    echo "Testing with $count Pods..."

    cat <<EOF | kubectl apply -f -
apiVersion: apps/v1
kind: Deployment
metadata:
  name: test-pod-deployment-$count
  namespace: $namespace
spec:
  replicas: $count
  selector:
    matchLabels:
      app: test-pod
  template:
    metadata:
      labels:
        app: test-pod
    spec:
      containers:
      - name: test-pod-container
        image: nginx:latest
        resources:
          requests:
            cpu: "100m"
            memory: "128Mi"
          limits:
            cpu: "200m"
            memory: "256Mi"
EOF

    echo "Waiting for Pods to be running..."
    kubectl wait --for=condition=available --timeout=180s deployment/test-pod-deployment-$count -n $namespace

    echo "Monitoring node resources..."
    echo "CPU and Memory usage for $count Pods:" >> pod-resources-results.txt
    date >> pod_mem.txt
    echo "CPU and Memory usage for $count Pods:" >> pod_mem.txt
    kubectl top pods -n kube-system | grep antrea-agent >> pod_mem.txt

    echo "Deleting Pods for $count..."
    kubectl delete deployment/test-pod-deployment-$count -n $namespace

    echo "Waiting for node resources to stabilize..."
    sleep 60
done

kubectl delete namespace $namespace