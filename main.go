package main

import (
	"flag"
	"fmt"
	"log"
	"os/exec"
	"strings"
)

func parseMessageFlag() string {
	message := flag.String("m", "", "Commit message")
	flag.Parse()

	return *message
}

func makeWorkInProgressCommit() {
	if err := runGitCommand("add", "-A"); err != nil {
		log.Fatalf("Failed to track git files: %v", err)
	}

	if err := runGitCommand("commit", "-m", "wip"); err != nil {
		log.Fatalf("Failed to make wip commit: %v", err)
	}

	fmt.Println("You've made a work-in-progress commit!")
}

func squashWorkInProgressCommits(message string) {
	commitsToSquash, err := countSquashableCommits()
	if err != nil {
		log.Fatalf("Failed to determine commits to squash: %v", err)
	}

	if commitsToSquash == 0 {
		log.Println("You have no work-in-progress commits!")
		return
	}

	if err := squashCommits(commitsToSquash, message); err != nil {
		log.Fatalf("Failed to squash work-in-progress commits: %v", err)
	}
	fmt.Println("You've squashed your work-in-progress commits!")
}

func parsePreviousGitMessages(back int) ([]string, error) {
	output, err := runGitCommandOutput("log", "-n", fmt.Sprintf("%d", back), "--pretty=%s")
	if err != nil {
		return nil, fmt.Errorf("Failed to get git history: %w", err)
	}
	return strings.Split(strings.TrimSpace(output), "\n"), nil
}

func countSquashableCommits() (int, error) {
	back := 10
	for {
		commitMessages, err := parsePreviousGitMessages(back)
		if err != nil {
			return 0, err
		}

		if count := findWipCommitCount(commitMessages); count > 0 {
			return count, nil
		}

		back++
	}
}

func findWipCommitCount(commitMessages []string) int {
	for i, message := range commitMessages {
		if message != "wip" {
			return i
		}
	}
	return len(commitMessages)
}

func squashCommits(back int, message string) error {
	if err := runGitCommand("reset", "--soft", fmt.Sprintf("HEAD~%d", back)); err != nil {
		return fmt.Errorf("Failed to reset commits: %w", err)
	}
	if err := runGitCommand("commit", "-m", message); err != nil {
		return fmt.Errorf("Failed to create squashed commit: %w", err)
	}
	return nil
}

func runGitCommand(args ...string) error {
	cmd := exec.Command("git", args...)
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("git %s: %w", strings.Join(args, " "), err)
	}
	return nil
}

func runGitCommandOutput(args ...string) (string, error) {
	cmd := exec.Command("git", args...)
	output, err := cmd.Output()
	if err != nil {
		return "", fmt.Errorf("git %s: %w", strings.Join(args, " "), err)
	}
	return string(output), nil
}

func main() {
	message := parseMessageFlag()
	if message == "" {
		makeWorkInProgressCommit()
	} else {
		squashWorkInProgressCommits(message)
	}
}
