#  Loggers Package

The loggers package is a comprehensive logging solution for Go applications, providing easy-to-use functionalities for logging messages at different levels and integrating with various logging systems like Loki, Sentry, and Prometheus.

### Features:
- Multiple log levels: Info, Debug, Warn, Error, and Fatal.
- Versatile logging functions supporting both simple and detailed messages.
- Integration with Loki for log aggregation.
- Integration with Sentry for error tracking.
- Prometheus integration for monitoring and metrics.
- Context and structured logging capabilities.
- Easy setup and configuration through environment variables.

## Getting Started:

## Installation:
Install the loggers package in your Go project using:
```bash
go get github.com/pworld/loggers
```

## Configuration:
The loggers package can be configured primarily through environment variables. These variables allow you to control various aspects of logging behavior, such as enabling or disabling specific loggers and configuring their settings. Here are the key environment variables you should set:

- LOKI_ACTIVE: Set this to "1" to enable logging to Loki. If this is not set or set to any other value, logging to Loki will be disabled.
- SENTRY_ACTIVE: Set this to "1" to enable logging to Sentry. As with Loki, not setting this or setting it to a different value will disable Sentry logging.
- LOKI_CLIENT: This should be set to the URL or address of your Loki client/server to enable proper communication with Loki.
- SENTRY_DSN: The Data Source Name (DSN) for Sentry. This is a key piece of information Sentry needs to communicate with your application.
You can set these variables directly in your operating system, through a configuration file, or within your application before initializing the logger.

## Initialization:
Before you start logging messages, the loggers package might require some initialization steps, especially if you are using Sentry. Here’s how you can initialize Sentry:

### Call InitSentry Function:
If your loggers package has a function like InitSentry, call this at the beginning of your main function. This function will set up Sentry with the appropriate DSN and other configurations.

### Defer the Cleanup Function:
InitSentry should return a cleanup function that you need to call when your application exits. Use the defer statement right after the initialization to ensure that this cleanup function is called at the end.

Here’s an example:
```go
package main

import (
    "github.com/pworld/loggers"
    "os"
)

func main() {
    // Set environment variables (or ensure they are set before running the application)
    os.Setenv("SENTRY_ACTIVE", "1")
    os.Setenv("SENTRY_DSN", "<your-sentry-dsn>")

    // Initialize Sentry
    cleanup := loggers.InitSentry()
    defer cleanup()

    // ... rest of your application code
}

```
## Usage:
To effectively use the loggers package in your Go applications, simply import the package and call its logging functions as needed. Below is an example demonstrating how to use some of the primary functions:

````go
package main

import "github.com/pworld/loggers"

func main() {
  logger := loggers.NewLogger()

  // Simple info log
  logger.Info("Application started")

  // Detailed debug log
  logger.Debug("Debugging data loaded", "main", "/debug", 200)

  // Error log with structured data
  if err := someFunction(); err != nil {
    logger.Error("Error encountered", "someFunction", "/errorPath", 500, err)
  }
}

func someFunction() error {
  // Example function
  return nil
}
````
## API Reference
The loggers package provides various functions to log messages at different levels. Here is a brief overview of some key functions:

- InfoLog(description, method, path string, status int, args ...interface{})
Logs an informational message.
    - description: Message to log.
    - method: The method in which the log is recorded.
    - path: API or function path.
    - status: HTTP status code or relevant status indicator.
    - args: Additional arguments to format the message.

- DebugLog(description, method, path string, status int, args ...interface{})
- Similar to InfoLog, used for logging debug-level messages.
- ErrorLog(description, method, path string, status int, args ...interface{})
Used for logging error-level messages, typically when an error occurs.
- Other functions 
Similar to the above, but for different log levels like Trace, Fatal, etc.

## Package Level Logger:
This code include a package-level default logger, allowing direct access to the Info method (and other logging methods) without needing to create a new Logger instance each time.

### Sample Code
```go
package main

import "github.com/pworld/loggers"

func main() {
    // Directly use the Info method from the loggers package
    loggers.Info("This is an info message")

    // You can also pass additional arguments if your Info method supports them
    loggers.Info("This is an info message with additional context", "more data", 123)
}

```
There can be potential conflicts or unintended behavior if you mix direct calls to package-level logging functions (like loggers.Info(...)) with calls made through an instance of the Logger struct (like logger.Info(...) from logger := loggers.NewLogger()). 
These conflicts or issues generally arise from differences in the state or configuration of the loggers.

To avoid these issues, it's generally a good practice to choose one logging approach and stick with it consistently throughout your application. If you decide to use a package-level logger, ensure that all parts of your application use this global logger. 
If you prefer creating individual logger instances, avoid using the package-level functions and ensure that each part of your application either receives its own logger instance or uses a shared instance that is passed around appropriately
## Contributing:
Contributions to the loggers package are welcome. If you're interested in contributing, you can follow these steps:

- Fork the Repository: Create your own fork of the loggers repository.
- Make Your Changes: Implement your feature or fix.
- Write Tests: Ensure your changes are covered by tests.
- Submit a Pull Request: Open a PR against the main loggers repository for review.
- Before contributing, please review the contribution guidelines, typically found in a CONTRIBUTING.md file in the repository.

## License:
The loggers package is released under the Apache v2.0. This permissive license allows for commercial use, modification, distribution, and private use.

## Contact:
Feel free to contribute or ask questions to make the loggers package more robust and user-friendly.