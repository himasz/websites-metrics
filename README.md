# Website Metrics Monitoring
```
## Project Structure

websites_metrics/
├── config/
│   ├── config.go
│   ├── config_loader.go
│   └── db_config.json
├── main.go
├── models/
│   └── metric.go
├── repository/
│   ├── postgres_metrics_repository.go
│   └── metrics_repository_interface.go
├── scheduler/
│   ├── cron_scheduler.go
│   └── scheduler_interface.go
└── metrics/
    ├── metrics_calculator_interface.go
    └── url_metrics_calculator.go
```
### Getting Started

```
### Prerequisites

- Go 1.22.5
- PostgreSQL database (docker compose is there)

### Installation

1. Install dependencies:
    ```sh
    go mod tidy
    ```

### Configuration

1. Configure the PostgreSQL database connection in `config/db_config.json`:
    ```json
    {
        "database": {
            "user": "yourusername",
            "password": "yourpassword",
            "dbname": "yourdatabase",
            "port": 5432
        }
    }
    ```

2. Add URLs and their check intervals in `config/urls_config.json`:
    ```json
    {
        "urls": [
            {
                "url": "http://example.com",
                "regex": "Example Domain",
                "interval": "@every 1m"
            },
            {
                "url": "http://another-example.com",
                "regex": "",
                "interval": "@every 5m"
            }
        ]
    }
    ```

### Usage

1. Run the application:
    ```sh
    go run main.go
    ```

### Project Details

#### Configuration Loader

- Located in `config/config_loader.go`
- Implements `Load` method to load JSON configurations

#### Repository

- Located in `repository/`
- `IMetricsRepository` interface defines the `Save` method
- `MetricsRepository` struct implements `IMetricsRepository`

#### Scheduler

- Located in `scheduler/`
- `IScheduler` interface defines methods for scheduling tasks
- `CronScheduler` struct implements `IScheduler`

#### Metrics

- Located in `metrics/`
- `IMetricsCalculator` interface defines methods for metrics calculating
- `URLMetricsCalculator` struct implements `URLChecker`
