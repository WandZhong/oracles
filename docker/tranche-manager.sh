#!/bin/bash

# Check that the environment variable has been set correctly
if [ -z "$SECRETS_BUCKET_NAME" ]; then
  echo >&2 'error: missing SECRETS_BUCKET_NAME environment variable'
  exit 1
fi

# Load the S3 secrets file contents into the environment variables
eval $(aws s3 cp s3://${SECRETS_BUCKET_NAME}/tranche-manager.txt - | sed 's/^/export /')


exec /tranche-manager -pg-user "$POSTGRES_USER" -pg-pwd "$POSTGRES_PASSWORD" -pg-db "$POSTGRES_DATABASE" -pg-addr "$POSTGRES_HOST:$POSTGRES_PORT" -port 8000
