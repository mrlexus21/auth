#!/bin/sh
until pg_isready -h pg -U auth-user-prod -d auth-prod; do
  echo "Waiting for PostgreSQL to start..."
  sleep 2
done
