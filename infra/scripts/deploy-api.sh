#!/usr/bin/env bash
# Build, push, and deploy the API container to ECS Fargate.
# Usage: AWS_REGION=us-east-1 ACCOUNT_ID=123456789012 ./deploy-api.sh [tag]

set -euo pipefail

: "${AWS_REGION:?AWS_REGION is required}"
: "${ACCOUNT_ID:?ACCOUNT_ID is required}"

TAG="${1:-$(git rev-parse --short HEAD)}"
REPO="${ACCOUNT_ID}.dkr.ecr.${AWS_REGION}.amazonaws.com/shrinkr/api"
IMAGE="${REPO}:${TAG}"

echo ">> logging in to ECR"
aws ecr get-login-password --region "${AWS_REGION}" \
  | docker login --username AWS --password-stdin "${ACCOUNT_ID}.dkr.ecr.${AWS_REGION}.amazonaws.com"

echo ">> building ${IMAGE}"
docker build -t "${IMAGE}" -f api/Dockerfile api

echo ">> pushing ${IMAGE}"
docker push "${IMAGE}"

echo ">> registering new task definition"
NEW_TD_ARN=$(aws ecs register-task-definition \
  --cli-input-json "file://infra/ecs/api-task-definition.json" \
  --query 'taskDefinition.taskDefinitionArn' --output text)

echo ">> updating service to ${NEW_TD_ARN}"
aws ecs update-service \
  --cluster shrinkr \
  --service shrinkr-api \
  --task-definition "${NEW_TD_ARN}" \
  --force-new-deployment >/dev/null

echo ">> done."
