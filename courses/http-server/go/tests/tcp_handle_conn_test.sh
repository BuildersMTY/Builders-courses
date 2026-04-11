#!/bin/bash
# Verify handleConnection processes a request without crashing the server.
# The platform has already spawned $BUILDERSMTY_BINARY and waited for readiness.

set -e

# Send a minimal HTTP request and let the connection close
echo -e "GET / HTTP/1.1\r\nHost: localhost\r\n\r\n" | timeout 3 nc localhost 8080 >/dev/null || true

# Verify the server is still alive after handling the request
timeout 2 bash -c "echo '' > /dev/tcp/localhost/8080"
echo "PASS: handleConnection processed request without crashing"
