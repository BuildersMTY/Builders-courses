#!/bin/bash
# Verify the server accepts TCP connections on port 8080.
# Tries to connect with a short timeout — success means the listener is open.

sleep 0.5  # give the server a moment to start

# Attempt a TCP connection using /dev/tcp (bash built-in)
timeout 2 bash -c "echo '' > /dev/tcp/localhost/8080" 2>/dev/null
if [ $? -eq 0 ]; then
    echo "PASS: TCP connection accepted on :8080"
    exit 0
fi

# Fallback: try with nc if /dev/tcp is not available
if command -v nc &>/dev/null; then
    echo "" | nc -w 2 localhost 8080 2>/dev/null
    if [ $? -eq 0 ]; then
        echo "PASS: TCP connection accepted on :8080"
        exit 0
    fi
fi

echo "FAIL: Could not connect to port 8080"
exit 1
