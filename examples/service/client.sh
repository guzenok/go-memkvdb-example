#!/bin/bash

set -e

KEY="0xZ"
VAL="SDjfsdjfueairtwerjg iejfuihaugfih sdfgkdfjgX"

echo ">>> WRITE val '${VAL}' for key '${KEY}': "
curl -i -X POST -H "Content-type: application/octet-stream" -d "${VAL}" http://127.0.0.1:8080/v1/values/${KEY}
echo

echo "<<< READ key '${KEY}':"
curl -i http://127.0.0.1:8080/v1/values/${KEY}
echo

echo "<<< READ key '${KEY}' again"
curl -i http://127.0.0.1:8080/v1/values/${KEY}
echo
