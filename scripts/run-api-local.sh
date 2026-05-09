#!/usr/bin/env bash
set -euo pipefail

ROOT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"

if [[ -f "${ROOT_DIR}/.env" ]]; then
  set -a
  # shellcheck disable=SC1091
  source "${ROOT_DIR}/.env"
  set +a
else
  echo "Missing ${ROOT_DIR}/.env. Copy .env.example to .env first." >&2
  exit 1
fi

cd "${ROOT_DIR}/api"
exec go run .
