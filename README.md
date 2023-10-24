# Golang CLI for Managing Agencies

This is a simple Golang command-line interface (CLI) application for managing agency information. It allows you to perform various operations, such as creating new agencies, listing agencies by region, updating agency information, and getting statistics about agencies in a specific region.

## Prerequisites

Before using this CLI, make sure you have Go installed on your system. If not, you can download it from the [official Go website](https://golang.org/dl/).

## Usage

1. **Compile and Run the CLI**

   To use this CLI, you need to compile and run the Go code.

    ```shell
    go run main.go
    ```
2. Available Commands


    get: Get agency information by agency ID.
    
    create: Create a new agency and enter agency information.
    
    list: List agencies in a specific region.
    
    update: Update agency information by agency ID.
    
    status: Get statistics about agencies in a specific region.
    Enter one of these commands when prompted.

## Example Workflow

1. Run the CLI using go run main.go.
2. Enter a command (e.g., create) and follow the prompts to create a new agency.
3. Use the list command to view agencies in a specific region.
4. Use the update command to modify agency information.
5. Get statistics about agencies in a region using the status command.

## License

This project is licensed under the MIT License. See the [LICENSE](https://opensource.org/license/mit/) file for details.