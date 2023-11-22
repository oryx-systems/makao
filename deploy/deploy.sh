#!/usr/bin/env sh

set -eux

# Create the namespace
kubectl create namespace $DEPLOY_NAMESPACE || true

helm upgrade \
    --install \
    --debug \
    --create-namespace \
    --namespace="${DEPLOY_NAMESPACE}" \
    --set app.name="${DEPLOY_RELEASE_NAME}" \
    --set app.container.env.replicaCount="${REPLICA_COUNT}" \
    --set app.container.env.dockerImage="${DOCKER_IMAGE}" \
    --set app.container.env.googleApplicationCredentials="${GOOGLE_APPLICATION_CREDENTIALS}" \
    --set app.container.env.googleProjectID="${GOOGLE_CLOUD_PROJECT_ID}" \
    --set app.container.env.db.databaseRegion="${DATABASE_REGION}" \
    --set app.container.env.db.databaseInstance="${DATABASE_INSTANCE}" \
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
    $DEPLOY_RELEASE_NAME \
    ./charts