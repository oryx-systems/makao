#!/usr/bin/env sh

set -eux

helm upgrade \
    --install \
    --debug \
    --create-namespace \
    --namespace="${DEPLOY_NAMESPACE}" \
    --set app.name="${DEPLOY_RELEASE_NAME}" \
    --set app.container.env.host="${DOMAIN_HOST}"\
    --set app.container.env.network.acme.adminEmail="${ADMIN_EMAIL}"\
    --wait \
    --timeout 300s \
    -f ./tls_deploy/values.yaml \
    $DEPLOY_RELEASE_NAME \
    ./tls_deploy


# Install Nginx Controller
helm repo add ingress-nginx https://kubernetes.github.io/ingress-nginx
helm repo update
helm install ingress-nginx ingress-nginx/ingress-nginx

# Setup cert manager
kubectl create namespace cert-manager || true

kubectl apply --validate=false -f https://github.com/cert-manager/cert-manager/releases/download/v1.13.3/cert-manager.crds.yaml || true

helm repo add jetstack https://charts.jetstack.io
helm repo update

helm install makao-cert-manager-release --namespace cert-manager --version v1.13.3 jetstack/cert-manager --wait --timeout 300s || true