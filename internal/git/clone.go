package git

import (
	"fmt"
	"os"
	"os/exec"
)

func CloneRepo(repoURL, token, clonePath string) error {
    if _, err := os.Stat(clonePath); os.IsNotExist(err) {
        if err := os.MkdirAll(clonePath, 0755); err != nil {
            return fmt.Errorf("failed to create clone path: %w", err)
        }

        cmd := exec.Command("git", "clone", 
            fmt.Sprintf("https://%s@%s", token, repoURL),
            clonePath,
        )
        if output, err := cmd.CombinedOutput(); err != nil {
            return fmt.Errorf("failed to clone repo: %s, %w", string(output), err)
        }
    } else {
        // If the folder exists, pull the latest changes
        cmd := exec.Command("git", "-C", clonePath, "pull")
        if output, err := cmd.CombinedOutput(); err != nil {
            return fmt.Errorf("failed to pull repo: %s, %w", string(output), err)
        }
    }
    return nil
}