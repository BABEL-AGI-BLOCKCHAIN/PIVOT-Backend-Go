```markdown
# PIVOT-Backend-Go

## Prerequisites
Before you begin, make sure your system has the following software installed:
- [Go 1.23.4](https://golang.org/dl/)
- [Git](https://git-scm.com/)
- [Make](https://www.gnu.org/software/make/)

### 1. Clone the repository
First, use Git to clone this repository:
```sh
git clone https://github.com/BABEL-AGI-BLOCKCHAIN/PIVOT-Backend-Go.git
cd PIVOT-Backend-Go
```

### 2. Install Go dependencies
Ensure you have Go 1.23.4 installed, then run the following command to install the required dependencies:
```sh
go mod tidy
```

### 3. Configuration file
Create a 

config.json

 file in the project root directory and configure it according to your needs. Here is a sample configuration file:
```json
{
    "listener": {
        "confirmation": 0, // Number of confirmation blocks, usually used to ensure transactions are confirmed by multiple blocks, 0 means no confirmation wait
        "endpoint": "your-endpoint", // RPC endpoint of the Ethereum node, e.g., URL of Infura or Alchemy
        "startHeight": 7658466, // Block height to start listening from
        "blockTime": 12, // Block time interval (seconds), used for the timer interval
        "fetchLimit": 5, // Limit on the number of blocks to fetch events from each time
        "TopicContractAddr": "0x9b764159249880e2d6B9a7F86495371c45aB69bC" // Address of the smart contract to listen to
    },
    "db": {
        "dsn": "testuser:123456@tcp(localhost:3306)/testdb?charset=utf8mb4&parseTime=True&loc=Local", // Database connection string, including username, password, host, port, and database name
        "driverName": "mysql", // Database driver name, e.g., mysql, postgres, etc.
        "maxOpenNum": 200, // Maximum number of open connections in the database connection pool
        "maxIdleNum": 20 // Maximum number of idle connections in the database connection pool
    }
}
```

### 4. Build and run
Use the `make` command to build and run the project:
```sh
make run
```

## Usage
After configuring 

config.json

, use the following command to start the listener:
```sh
make run
```

## Directory structure
Here is the main directory structure of the project:
```
PIVOT-Backend-Go/
├── LICENSE
├── Makefile
├── README.md
├── common
│   ├── config
│   │   └── config.go
│   └── database
│       ├── config.go
│       ├── db.go
│       └── db_test.go
├── config.json
├── contract
│   ├── PivotTopic.go
│   └── abi
│       └── backend_abi.go
├── go.mod
├── go.sum
├── listener
│   └── listener.go
├── logic
│   ├── event_parser.go
│   ├── event_update.go
│   └── fetcher.go
├── main.go
├── orm
│   ├── connection.go
│   ├── rawEvents.go
│   ├── schema.go
│   ├── topic.go
│   └── user.go
├── relayer
│   └── relayer.go
├── run
├── types
│   └── schema.go
└── utils
    └── utils.go
```

## FAQ
### 1. How to install Go?
Please refer to the [official Go documentation](https://golang.org/doc/install) for installation instructions.

### 2. How to resolve dependency issues?
Run the `go mod tidy` command to install and tidy up dependencies.

### 3. How to view logs?
Logs will be output to the console, and you can view them in the terminal during runtime.

## License
This project is licensed under the MIT License. For more details, please refer to the 

LICENSE

 file.
```