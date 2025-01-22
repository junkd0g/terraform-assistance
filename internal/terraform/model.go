package terraform

type Plan struct {
    FormatVersion    string          `json:"format_version"`
    TerraformVersion string          `json:"terraform_version"`
    Variables        map[string]Variable
    PlannedValues    *Values         `json:"planned_values,omitempty"`
    PriorState       *State          `json:"prior_state,omitempty"`
    Configuration    *Configuration  `json:"configuration,omitempty"`
    // There are also fields like "planned_changes", "resource_changes", etc.
    // You can add them here if you want to see the diffs or changes specifically.
}

type Values struct {
    Outputs    map[string]Output `json:"outputs,omitempty"`
    RootModule *Module           `json:"root_module,omitempty"`
}

type Module struct {
    Resources    []Resource `json:"resources,omitempty"`
    Address      string     `json:"address,omitempty"`
    ChildModules []Module   `json:"child_modules,omitempty"`
}

// Resource represents a Terraform resource in the plan
type Resource struct {
    Address       string                 `json:"address,omitempty"`
    Mode          string                 `json:"mode,omitempty"`
    Type          string                 `json:"type,omitempty"`
    Name          string                 `json:"name,omitempty"`
    ProviderName  string                 `json:"provider_name,omitempty"`
    SchemaVersion int                    `json:"schema_version,omitempty"`
    Values        map[string]interface{} `json:"values,omitempty"`
    // If you need deeper inspection of "values" (e.g. specific fields for an IAM policy),
    // define more typed structures and unmarshal accordingly.
}

// Output represents a Terraform output value
type Output struct {
    Sensitive bool        `json:"sensitive,omitempty"`
    Value     interface{} `json:"value,omitempty"`
}

// Variable is used to represent input variables (if needed)
type Variable struct {
    Value interface{} `json:"value,omitempty"`
}

// State is a partial representation of the prior state in the plan
type State struct {
    // For example:
    FormatVersion string  `json:"format_version,omitempty"`
    Values        *Values `json:"values,omitempty"`
    // You can add more fields if you want to inspect prior resources.
}

// Configuration holds the Terraform configuration structure
type Configuration struct {
    // For example, references to provider requirements, root module configuration, etc.
    ProviderConfig map[string]ProviderConfig `json:"provider_config,omitempty"`
    RootModule     *ModuleConfig            `json:"root_module,omitempty"`
    // And so on...
}

type ProviderConfig struct {
    Name   string                 `json:"name,omitempty"`
    Expressions map[string]interface{} `json:"expressions,omitempty"`
}

type ModuleConfig struct {
    Resources    []ConfigResource `json:"resources,omitempty"`
    ModuleCalls  map[string]ModuleCall `json:"module_calls,omitempty"`
    // Expand as needed...
}

type ConfigResource struct {
    Address     string                 `json:"address,omitempty"`
    Mode        string                 `json:"mode,omitempty"`
    Type        string                 `json:"type,omitempty"`
    Name        string                 `json:"name,omitempty"`
    Provider    string                 `json:"provider_config_key,omitempty"`
    Expressions map[string]interface{} `json:"expressions,omitempty"`
}

type ModuleCall struct {
    Source      string      `json:"source,omitempty"`
    Module      *ModuleConfig `json:"module,omitempty"`
}