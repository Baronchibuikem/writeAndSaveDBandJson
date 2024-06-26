# MongoDB Backup and Restore Utility

This project provides a set of utilities to back up MongoDB collections to JSON files and restore them from JSON files. The utilities are written in Go and make use of MongoDB's Go driver.

## Features

- **Backup MongoDB Collections to JSON**: Exports all collections from a specified MongoDB database to JSON files.
- **Restore JSON Files to MongoDB**: Imports data from JSON files into the specified MongoDB database.

## Prerequisites

- Go 1.21 or later
- MongoDB instance running locally or remotely
- Access to the terminal/command line

## Getting Started

### Installation

1. **Clone the Repository**:
    ```sh
    git clone https://github.com/yourusername/mongodb-backup-restore.git
    cd mongodb-backup-restore
    ```

2. **Install Dependencies**:
    ```sh
    go mod download
    go mod tidy
    ```

### Configuration

- Ensure your MongoDB instance is running.
- Modify the MongoDB URI and database name in the code if necessary.

### Directory Structure

- JSON files for backup will be saved in the `~/Documents/golang/mongodb_backup` directory.
- JSON files for restoration should be placed in the `~/Downloads/mongodb_backup` directory.

### Usage

#### Backup MongoDB Collections to JSON

To backup all collections in the specified MongoDB database to JSON files, use the `DBToJson` function.

**Example**:
```go
package main

import "yourmodule/cmd"

func main() {
    cmd.DBToJson()
}
```

### Run the script:
    ```sh
    go run main.go
    ```

### License

This project is licensed under the MIT License - see the LICENSE file for details.

### Contributing

Contributions are welcome! Please open an issue or submit a pull request with your changes.

### Acknowledgments
- MongoDB Go Driver
- Go