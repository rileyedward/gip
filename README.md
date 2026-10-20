# gip

CLI tool for work-in-progress commits.

> **Status:** Complete — no further development planned, but ready to download and use.

## Overview

### What is gip?

gip is a lightweight command-line tool for developers who make lots of small work-in-progress (WIP) commits. One command commits everything in progress with a `wip` message, and another squashes those WIP commits into a single commit with a real message, without leaving the terminal or reaching for a full Git client like GitKraken or Tower.

### Key Features

- **WIP Commit**: Quickly commit all the work you currently have in progress.
- **WIP Squash**: Rebase all recent WIP commits into a single commit with the message you provide.

## Getting Started

### Prerequisites

1. **Go** is required to build the tool. Check it by running:

    ```bash
    go version
    ```

### Installation

1. Build the executable:

    ```bash
    go build -o gip main.go
    ```

2. Move it somewhere on your `PATH`:

    ```bash
    sudo mv gip /usr/local/bin
    ```

## Usage

### Make a Work-in-Progress Commit

```bash
gip
```

This makes a Git commit in the current directory with the message `wip`.

### Squash Work-in-Progress Commits

```bash
gip -m "Your commit message"
```

This squashes the recent commits in the current directory that have the `wip` message into a single commit with the message you pass to `-m`.

## Contributing

No new features are planned, but it's meant to be used: bug reports through [Issues](../../issues) are welcome, and small fixes may be merged when I get to them. For anything bigger, you're welcome to fork it and take it in your own direction under the license below.

## License

Released under the [MIT License](LICENSE).
