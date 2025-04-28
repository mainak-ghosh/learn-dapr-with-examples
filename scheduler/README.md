# Dapr Simple Cron Job Scheduler Example

This example demonstrates how to use Dapr to create a scheduled job (cron job) that triggers an API endpoint at regular intervals.

## Components

- `main.go`: A Go HTTP server that handles both direct API calls and scheduled job triggers
- `resources/jobscheduler.yaml`: Dapr cron binding component configuration
- `resources/job.yaml`: Dapr job configuration that defines the schedule and actions

## Prerequisites

1. [Go](https://golang.org/doc/install) installed
2. [Dapr CLI](https://docs.dapr.io/getting-started/install-dapr-cli/) installed and initialized

## Project Structure

```
scheduler/
├── main.go                 # Go HTTP server
├── README.md              # This file
└── resources/
    ├── job.yaml           # Job configuration
    └── jobscheduler.yaml  # Cron binding component
```

## How It Works

1. The `jobscheduler.yaml` defines a cron binding component that triggers every minute
2. The `job.yaml` configures the job to use this binding
3. The Go server handles the scheduled triggers at the `/scheduled-job` endpoint

## Running the Example

1. Start the application with Dapr:

   ```bash
   dapr run \
     --app-id scheduler-app \
     --app-port 3000 \
     --resources-path ./resources \
     --dapr-http-port 3500 \
     --log-level debug \
     go run main.go
   ```

2. The scheduler will automatically trigger every minute, and you'll see logs in the console.

3. You can also manually trigger the API endpoint:
   ```bash
   curl -X POST http://localhost:3000/call-api
   ```

## Understanding the Code

- The cron binding is configured to run every minute (`* * * * *`)
- The Go server has two endpoints:
  - `/call-api`: For direct HTTP calls
  - `/scheduled-job`: For handling scheduled job triggers from Dapr
- When triggered, it logs the event and returns an OK response
- You can add your own API call logic in the `handleScheduledJob` function

## Logs

You can see the scheduler working by watching the logs. Successful triggers will show:

```
Received scheduled job trigger from [address]
Calling the external API...
```

## Troubleshooting

If the scheduler isn't working:

1. Check if Dapr is initialized: `dapr init`
2. Verify the application is running: `dapr list`
3. Check the logs with `--log-level debug` for detailed information
4. Ensure ports 3000 and 3500 are available
