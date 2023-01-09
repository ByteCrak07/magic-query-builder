# Magic Query Builder

This is a command-line interface (CLI) application written in [Go](https://go.dev) that can be used to generate queries.

## Steps to run the app

### 1. Building the CLI 🛠️

```bash
# Clone the repository
git clone https://github.com/ByteCrak07/magic-query-builder.git

# Change into the project directory
cd magic-query-builder

# Build the executable (optional)
go build .
```

### 2. Running the CLI 💻

To run the CLI, open a terminal and navigate to the project directory.

```bash
# To launch the interactive shell
./magic-query-builder

# To read queries from a file (refer example.txt)
./magic-query-builder <file>

# To see the usage instructions, use the --help flag
./magic-query-builder --help
```

To run the CLI without building the executable use:

```bash
# Interactive shell
go run .

# Help
go run . --help
```

## Made with [<img height="20" style="margin-bottom:-2px" src="https://go.dev/images/go-logo-blue.svg">](https://go.dev)
