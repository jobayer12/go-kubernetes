## Go version 1.21.1

## Install the dependencies using the below command:
`go mod tidy`

## Generate the swagger documentation
`swag init --parseDependency true`

Note: Make sure `swag` already installed on your local machine


## Run the project using the below command:
`go run main.go`

Note: Before running the project, Ensure you can access the target Kubernetes cluster using kubectl. The kubeconfig file should be in the user’s home directory .kube folder for the setup to work as expected.

