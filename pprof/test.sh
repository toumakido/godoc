#!/bin/bash
# Test script for pprof demo

set -e

echo "=== pprof Demo Test ==="
echo

# Check if server is running
if ! curl -s http://localhost:8080 > /dev/null; then
    echo "Error: Server is not running on :8080"
    echo "Please start the server with: go run main.go"
    exit 1
fi

echo "✓ Server is running"
echo

# Test CPU endpoint
echo "Testing CPU endpoint..."
curl -s http://localhost:8080/cpu
echo

# Test Memory endpoint
echo "Testing Memory endpoint..."
curl -s http://localhost:8080/memory
echo

# Test Goroutine endpoint
echo "Testing Goroutine endpoint..."
curl -s http://localhost:8080/goroutine
echo

# Test Block endpoint
echo "Testing Block endpoint..."
curl -s http://localhost:8080/block
echo

# Show pprof index
echo "Fetching pprof index..."
curl -s http://localhost:8080/debug/pprof/ | head -20
echo

# Show goroutine profile
echo "Goroutine profile (first 20 lines)..."
curl -s "http://localhost:8080/debug/pprof/goroutine?debug=1" | head -20
echo

echo "=== All tests completed ==="
echo
echo "To profile CPU for 10 seconds:"
echo "  go tool pprof http://localhost:8080/debug/pprof/profile?seconds=10"
echo
echo "To profile heap:"
echo "  go tool pprof http://localhost:8080/debug/pprof/heap"
