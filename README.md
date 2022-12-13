# pdtf

Utility to create Terraform code (HCL) for managing a PagerDuty account. Currently, the only command is `mkdir`

## Build Instructions

If you're new to Go, check out ["How to Build and Install Go Programs"](https://www.digitalocean.com/community/tutorials/how-to-build-and-install-go-programs)

## Running the Program
In the root directory (`pdtf/`) of the project run either `go run .` to run pdtf or `go build` to build the pdtf binary. The project also uses
[Taskfile](https://taskfile.dev/) to run the build process.
If you build the binary, you can run it by calling `./pdtf` in the project directory or add it to your PATH and run `pdtf` from anywhere.

## Usage
### `pdtf mkdir <directory_name>`

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

## Collaboration
Create an issue or open a PR if there are other commands you would like to see.
