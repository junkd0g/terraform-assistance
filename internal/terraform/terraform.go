package terraform

import (
	"encoding/json"
	"fmt"
	"os/exec"
)

func RunTerraformPlan(terraformDir string) (*Plan, error) {
    cmdInit := exec.Command("terraform", "init")
    cmdInit.Dir = terraformDir
    if output, err := cmdInit.CombinedOutput(); err != nil {
        return nil, fmt.Errorf("terraform init failed: %s, %w", string(output), err)
    }

    cmdPlan := exec.Command("terraform", "plan", "-out=plan.out")
    cmdPlan.Dir = terraformDir
    if output, err := cmdPlan.CombinedOutput(); err != nil {
        return nil, fmt.Errorf("terraform plan failed: %s, %w", string(output), err)
    }

    cmdShow := exec.Command("terraform", "show", "-json", "plan.out")
    cmdShow.Dir = terraformDir
    planOut, err := cmdShow.CombinedOutput()
    if err != nil {
        return nil, fmt.Errorf("terraform show failed: %s, %w", string(planOut), err)
    }

    var plan Plan
    if err := json.Unmarshal(planOut, &plan); err != nil {
        return nil, fmt.Errorf("failed to unmarshal plan: %w", err)
    }

    return &plan, nil
}