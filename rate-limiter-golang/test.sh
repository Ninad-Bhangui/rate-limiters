#!/bin/bash
if [ $# -ne 1 ]; then
  echo "Usage: $0 <no_of_concurrent_requests>"
  exit 1
fi

concurrent_req=$1

echo "Calling curl $concurrent_req times"
# Send 4 concurrent requests to your API within 1 second
make_request() {
  response=$(curl -s -o /dev/null -w "%{http_code}%" -X GET http://localhost:3000)
  if [[ $response =~ ^[2][0-9][0-9][%]$ ]]; then
    echo "Success"
  else
    echo "Failure"
  fi
}
status_file=$(mktemp /tmp/statusfile.XXXXXX)
for ((i=1; i <= concurrent_req; i++)); do
    # Send a request in the background with "&"
    make_request &>> "$status_file" &
done

# Wait for all background jobs to finish
wait

success_count=$(grep -c "Success" "$status_file")
failure_count=$(grep -c "Failure" "$status_file")
echo "Successful requests: $success_count"
echo "Failed requests: $failure_count"
