#!/usr/bin/env sh

set -eux

# Create the namespace
kubectl create namespace $DEPLOY_NAMESPACE || true

# Delete Kubernetes secret if exists
kubectl delete secret makao-service-account --namespace $DEPLOY_NAMESPACE || true

# Create GCP service account file
cat $GOOGLE_APPLICATION_CREDENTIALS >> ./service-account.json

# Recreate service account file as Kubernetes secret
kubectl create secret generic makao-service-account \
    --namespace $DEPLOY_NAMESPACE \
    --from-file=key.json=./service-account.json

helm delete cert-manager --namespace cert-manager || true
kubectl delete namespace cert-manager || true
# kubectl delete -f https://github.com/cert-manager/cert-manager/releases/download/v1.13.2/cert-manager.crds.yaml

helm repo add jetstack https://charts.jetstack.io
helm repo update

# Refer to this: https://cert-manager.io/docs/installation/helm/
kubectl apply -f https://github.com/cert-manager/cert-manager/releases/download/v1.13.2/cert-manager.yaml || true


# Install cert manager Helm chart
# helm install \
#   cert-manager jetstack/cert-manager \
#   --namespace cert-manager \
#   --create-namespace \
#   --version v1.13.2 \
#   --wait \
#   --timeout 300s \
#   --set installCRDs=true

kubectl label crd certificaterequests.cert-manager.io app.kubernetes.io/managed-by=Helm --overwrite
kubectl annotate crd certificaterequests.cert-manager.io meta.helm.sh/release-name=cert-manager meta.helm.sh/release-namespace=cert-manager --overwrite


# Deploying
helm upgrade \
    --install \
    --debug \
    --create-namespace \
    --namespace="${DEPLOY_NAMESPACE}" \
    --set app.name="${DEPLOY_RELEASE_NAME}" \
    --set app.container.env.replicaCount="${REPLICA_COUNT}" \
    --set app.container.env.dockerImage="${DOCKER_IMAGE}" \
    --set app.container.env.googleCloudProjectID="${GOOGLE_CLOUD_PROJECT_ID}" \
    --set app.container.env.db.databaseRegion="${DATABASE_REGION}" \
    --set app.container.env.db.databaseInstance="${DATABASE_INSTANCE}" \
    --set app.container.env.db.databaseInstanceConnectionName="${DATABASE_INSTANCE_CONNECTION_NAME}" \
    --set app.container.env.db.postgresUser="${POSTGRES_USER}" \
    --set app.container.env.db.postgresPassword="${POSTGRES_PASSWORD}" \
    --set app.container.env.db.postgresHost="${POSTGRES_HOST}" \
    --set app.container.env.db.postgresPort="${POSTGRES_PORT}" \
    --set app.container.env.db.postgresDB="${POSTGRES_DB}" \
    --set app.container.env.serviceName="${SERVICE_NAME}" \
    --set app.container.env.db.rootCollectionSuffix="${ROOT_COLLECTION_SUFFIX}" \
    --set app.container.env.debug="${DEBUG}" \
    --set app.container.env.db.repository="${REPOSITORY}" \
    --set app.container.env.aitAPIKey="${AIT_API_KEY}" \
    --set app.container.env.aitUsername="${AIT_USERNAME}" \
    --set app.container.env.aitSenderID="${AIT_SENDER_ID}" \
    --set app.container.env.aitEnvironment="${AIT_ENVIRONMENT}" \
    --set app.container.env.googleProjectNumber="${GOOGLE_PROJECT_NUMBER}" \
    --set app.container.env.firebaseWebAPIKey="${FIREBASE_WEB_API_KEY}" \
    --set app.container.env.oryxProviderChannel="${ORYX_PROVIDER_CHANNEL}" \
    --set app.container.env.jwtSecret="${JWT_SECRET}" \
    --set app.container.env.port="${PORT}" \
    --set app.container.env.defaultResidenceId="${DEFAULT_RESIDENCE_ID}" \
    --set app.container.env.host="${DOMAIN_HOST}"\
    --wait \
    --timeout 300s \
    -f ./charts/values.yaml \
    $DEPLOY_RELEASE_NAME \
    ./charts