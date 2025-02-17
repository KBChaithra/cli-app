# CLI App

## Overview

CLI application to run processes defined in YAML format and log the results. Each process has a unique name, parameters, and tasks defined within it. Tasks are executed sequentially and can either run local commands, remote SSH commands, or generate files using templates.

## Prerequisites

Before you begin, ensure you have met the following requirements:
- Installed and configured the latest [Golang](https://golang.org/doc/install).
- Installed SSH CLI.
- A development environment like a Golang IDE or text editor such as [VS Code](https://code.visualstudio.com/).
- An available GCP Virtual Machine (contact your buddy).
- Installed `delve` for debugging (optional).

## Installation

Follow these steps to install and set up the project:

1. **Clone the repository:**

   
   git clone https://github.com/yourusername/cli-app.git
   cd cli-app
  

2. **Install dependencies:**

   Ensure all Go dependencies are installed:

   go mod tidy
  

## Usage

To use the CLI application, follow these steps:

1. **Build the application:**
  
   go build -o cli-app
  

2. **Run the application:**

   
   ./cli-app run -c config.json -t /path/to/yaml/folder
   

   - `-c`: Path to the configuration file (JSON) which specifies the process.
   - `-t`: Path to the folder where YAML process definitions are stored.

## Configuration

### Configuration File (JSON)

The configuration file specified by the `-c` flag should be a JSON file containing the process name and its parameters. Example of `config.json`:

json
{
  "processName": "exampleProcess",
  "parameters": {
    "param1": "value1",
    "param2": "value2"
  }
}


## Task Classes

The following task classes are implemented:

- **localCmd**: Executes commands in the local terminal.
- **sshCmd**: Executes commands over SSH on a remote machine.
- **writefile**: Uses the Go template language to create and save a file.

### Running Tasks

Tasks defined in a process are executed sequentially. Parameters defined at the process level are passed down to tasks using Go templating.

## Logging

The application logs its state into a file and also outputs logs to the terminal. The log file is created in the same directory as the executable with a timestamp in the filename.

## Testing

To run tests, use the following command:


go test ./...


## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
