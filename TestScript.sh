#!/bin/bash

# Base URL for the API
BASE_URL="http://localhost:8080"

# Initialize an array to hold failed tests
failed_tests=()

# Function to test an endpoint
test_endpoint() {
    local method=$1
    local endpoint=$2
    local expected_code=$3

    echo "Testing $method $endpoint..."
    response_code=$(curl -s -o /dev/null -w "%{http_code}" -X "$method" "$BASE_URL$endpoint")

    if [[ "$response_code" -ne "$expected_code" ]]; then
        echo "Failed: Expected $expected_code but got $response_code for $method $endpoint"
        failed_tests+=("$method $endpoint")
    else
        echo "Success: $method $endpoint returned $response_code"
    fi
}

# Test public routes
test_endpoint "GET" "/" 200
test_endpoint "GET" "/product" 200
test_endpoint "GET" "/terms" 200

# Test dashboard routes
test_endpoint "GET" "/dashboard" 200
test_endpoint "GET" "/collections" 200
test_endpoint "GET" "/settings" 200

# Test product management routes
test_endpoint "POST" "/products" 201  # Assuming creation returns 201
test_endpoint "GET" "/products" 200

# Replace <ID> with an actual product ID for the following tests
PRODUCT_ID=1
test_endpoint "GET" "/products/$PRODUCT_ID" 200
test_endpoint "PUT" "/products/$PRODUCT_ID" 200
test_endpoint "DELETE" "/products/$PRODUCT_ID" 204  # Assuming deletion returns 204

# Summary of failed tests
if [ ${#failed_tests[@]} -eq 0 ]; then
    echo "All tests passed successfully!"
else
    echo "Some tests failed:"
    for test in "${failed_tests[@]}"; do
        echo "- $test"
    done
fi