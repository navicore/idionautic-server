#!/bin/bash

# Arguments
URL=$1                   # Server URL (no default, required)
RATE=${2:-60}            # Posts per minute (default: 60)
COUNT=${3:-100}          # Number of posts per process (default: 100)
PROCESSES=${4:-5}        # Number of concurrent processes (default: 5)

if [ -z "$URL" ]; then
  echo "Usage: $0 <URL> [RATE] [COUNT] [PROCESSES]"
  echo "Example: $0 http://localhost:8080/ingest 120 50 10"
  exit 1
fi

generate_payload() {
  local timestamp=$(date -u +"%Y-%m-%dT%H:%M:%SZ")  # ISO 8601 format

  cat <<EOF
{
  "eventType": "$(shuf -n1 -e functionCall widgetRendered tabOpened)",
  "target": "$(shuf -n1 -e calculateMetrics renderChart loadTab)",
  "count": $((RANDOM % 10 + 1)),
  "timestamp": "$timestamp"
}
EOF
}

send_post() {
  payload=$(generate_payload)

  # Debugging: Print the payload before sending it
  echo "Sending payload: $payload"

  response=$(curl -s -o /dev/null -w "%{http_code}" -X POST -H "Content-Type: application/json" -d "$payload" "$URL")
  
  if [ "$response" -ne 202 ]; then
    echo "Error: Received HTTP $response. Payload was: $payload"
    exit 1
  fi
}

# Function for a single process to send telemetry posts
run_process() {
  local rate=$1
  local count=$2
  local delay_ms=$((60000 / rate)) # delay in ms based on rate
  local posts_sent=0

  while [ $posts_sent -lt $count ]; do
    send_post
    posts_sent=$((posts_sent + 1))

    # Calculate random delay in seconds with 3 decimal places for compatibility
    delay_sec=$(awk "BEGIN {printf \"%.3f\", ($RANDOM % $delay_ms) / 1000}")
    sleep "$delay_sec"
  done
}

# Export the function for use by xargs
export -f send_post generate_payload run_process
export URL

# Run the processes with xargs
seq $PROCESSES | xargs -I{} -n1 -P$PROCESSES bash -c "run_process $RATE $COUNT"
