#!/bin/bash
# Verify path traversal attempts are blocked with 403.
# Uses raw TCP to bypass HTTP client path normalization.
# The platform has already spawned $BUILDERSMTY_BINARY and waited for readiness.

set -e

RESPONSE=$(echo -e "GET /../../../etc/passwd HTTP/1.1\r\nHost: localhost\r\n\r\n" | timeout 3 nc localhost 8080)

if echo "$RESPONSE" | grep -q "403"; then
    echo "PASS: Path traversal blocked with 403"
    exit 0
fi

echo "FAIL: Path traversal was not blocked (expected 403)"
echo "Response: $RESPONSE"
exit 1
