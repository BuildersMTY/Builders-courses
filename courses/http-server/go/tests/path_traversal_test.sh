#!/bin/bash
# Verify that path traversal attempts are blocked with 403.
# Uses raw TCP to bypass HTTP client path normalization.

sleep 0.5  # give the server a moment to start

# Send a path traversal request via raw TCP (bypasses client normalization)
RESPONSE=$(echo -e "GET /../../../etc/passwd HTTP/1.1\r\nHost: localhost\r\n\r\n" | timeout 3 nc localhost 8080 2>/dev/null)

if echo "$RESPONSE" | grep -q "403"; then
    echo "PASS: Path traversal blocked with 403"
    exit 0
fi

# Fallback: try with /dev/tcp
RESPONSE=$(timeout 3 bash -c 'exec 3<>/dev/tcp/localhost/8080; echo -e "GET /../../../etc/passwd HTTP/1.1\r\nHost: localhost\r\n\r\n" >&3; cat <&3; exec 3>&-' 2>/dev/null)

if echo "$RESPONSE" | grep -q "403"; then
    echo "PASS: Path traversal blocked with 403"
    exit 0
fi

echo "FAIL: Path traversal was not blocked (expected 403)"
echo "Response: $RESPONSE"
exit 1
