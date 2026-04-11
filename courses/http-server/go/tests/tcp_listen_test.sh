#!/bin/bash
# Verify the student's binary accepts TCP connections on port 8080.
# The platform has already spawned $BUILDERSMTY_BINARY and waited for readiness.

set -e

timeout 2 bash -c "echo '' > /dev/tcp/localhost/8080"
echo "PASS: TCP connection accepted on :8080"
