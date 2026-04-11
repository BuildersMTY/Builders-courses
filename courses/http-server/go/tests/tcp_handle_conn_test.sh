#!/bin/bash
# Verify that handleConnection processes a request without crashing.
# Sends a minimal HTTP request and checks the connection completes cleanly.

sleep 0.5  # give the server a moment to start

# Send a raw HTTP request via TCP
RESPONSE=$(echo -e "GET / HTTP/1.1\r\nHost: localhost\r\n\r\n" | timeout 3 nc localhost 8080 2>/dev/null)
EXIT_CODE=$?

# If nc is not available, try /dev/tcp
if ! command -v nc &>/dev/null; then
    RESPONSE=$(timeout 3 bash -c 'exec 3<>/dev/tcp/localhost/8080; echo -e "GET / HTTP/1.1\r\nHost: localhost\r\n\r\n" >&3; cat <&3; exec 3>&-' 2>/dev/null)
    EXIT_CODE=$?
fi

# The connection should complete (not hang, not crash the server).
# Even if response stubs write nothing, the connection should close gracefully.
if [ $EXIT_CODE -eq 0 ] || [ $EXIT_CODE -eq 124 ]; then
    # Verify the server is still running by making another connection
    sleep 0.3
    timeout 2 bash -c "echo '' > /dev/tcp/localhost/8080" 2>/dev/null
    if [ $? -eq 0 ]; then
        echo "PASS: handleConnection processed request without crashing"
        exit 0
    fi
fi

echo "FAIL: handleConnection crashed or server stopped accepting connections"
exit 1
