# Net Monitor

Net Monitor is a simple CLI tool to monitor network interfaces activity. It provides real-time statistics on network usage.

## Table of Contents

- [Net Monitor](#net-monitor)
  - [Table of Contents](#table-of-contents)
  - [Installation](#installation)
  - [Usage](#usage)
  - [Commands](#commands)
    - [`monitor`](#monitor)
  - [Development](#development)
    - [Prerequisites](#prerequisites)
    - [Running the Project](#running-the-project)
    - [Project Structure](#project-structure)
  - [Dependencies](#dependencies)
  - [License](#license)

## Installation

To build the project, you need to have Go installed. Clone the repository and run the following command to build the executable:

On Windows:

```powershell
go build -o bin/nm.exe
```

On Linux/Unix:

```bash
go build -o bin/nd
```

Alternatively, you can use the provided VS Code task to build the project:

Open the Command Palette (Ctrl+Shift+P).
Select Tasks: Run Build Task.

## Usage

After building the project, you can run the executable to start monitoring network interfaces:

`nm monitor`

## Commands

### `monitor`

The `monitor` command starts monitoring network interfaces and displays real-time statistics.

## Development

### Prerequisites

- Go 1.23.0 or higher
- VS Code (optional, but recommended)

### Running the Project

To run the project using VS Code, you can use the provided launch configuration:

1. Open the Command Palette (Ctrl+Shift+P).
2. Select Debug: Start Debugging.

This will start the `monitor` command.

### Project Structure

- `cmd/`: Contains the CLI commands.
  - `monitor.go`: Implements the monitor command.
  - `root.go`: Defines the root command and initializes subcommands.
- `main.go`: Entry point of the application.

## Dependencies

The project uses the following dependencies:

`github.com/shirou/gopsutil`: For fetching network statistics.
`github.com/spf13/cobra`: For building the CLI.

## License

This project is licensed under the MIT License. See the [LICENSE](./LICENSE) file for details.
