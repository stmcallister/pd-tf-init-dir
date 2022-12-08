# pdtf

Utility to create Terraform code for managing a PagerDuty account. Currently the only command is `mkdir`

`pdtf mkdir <directory_name>`

This command creates a directory with the name that is passed as an argument. Inside the directory two files are created: `main.tf` and `terraform.tf`. 

`terraform.tf` has the `required_providers` block defined with `PagerDuty/pagerduty` as the source with the following code.

```hcl
terraform {
  required_providers {
    pagerduty = {
      source = "PagerDuty/pagerduty"
    }
  }
}
```


The command then opens the newly created directory in [VS Code](https://code.visualstudio.com/).

Building the binary requires [Taskfile](https://taskfile.dev/).

Create an issue or open a PR if there are other commands you would like to see.
