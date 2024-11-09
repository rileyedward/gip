# gip

gip is the perfect CLI tool for developers who struggle with micro-commits, making it much simpler to use work-in-progress commits.

## Getting Started

### Prerequisites

Ensure you have the following prerequisites installed on your system. You can verify each installation by running the provided commands in your terminal.

1. **Go** is required for the application. Check if Go is installed by running:

   ```bash
   go version
   ```

### Installation

1. Build the application:

   ```bash
   go build -o gip main.go
   ```

2. Install the executable:

   ```bash
   sudo mv gip /usr/local/bin
   ```

## Usage

### Work-In-Progress Commit

To make a work-in-progress commit, simply run:

```bash
gip
```

This will make a git commit in your current directory with a message of "wip"

### Squash Work-In-Progress Commits

To squash your current stash of work-in-progress commits into a single commit, use the following command:

```bash
gip -m "[MESSAGE GOES HERE]"
```

This will squash the previous commits in your current directory with the "wip" message into a single commit with the message that you provide in the `-m` flag.
