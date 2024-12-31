# gip

CLI tool for work-in-progress commits

## Overview

### What is gip?

gip is a lightweight command-line tool designed to streamline work-in-progress (WIP) commits for developers. If you often find yourself juggling micro-commits during development, gip helps you manage them efficiently without complicating your Git workflow.

### Why Use gip?

Not everyone wants to use a full-fledged Git client like Kraken or GitTower. gip is perfect for developers who prefer working directly from the command line but still want enhanced tooling to simplify their Git workflow. It’s built for convenience, speed, and minimalism.

### Key Features

- **WIP Commit**: Quickly make a commit for all the work you currently have in-progress.
- **WIP Rebasing**: Rebase all recent "WIP Commits" into a single commit with a provided message.

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
