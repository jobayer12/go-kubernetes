## Go version 1.21.1

## Install the dependencies using the below command:
```shell
go mod tidy
```

## Generate the swagger documentation
```sh
swag init --parseDependency true
```
Note: Make sure `swag` already installed on your local machine

## Run the project using the below command:
<b>Note</b>: Before running the project, Ensure you can access the target Kubernetes cluster using kubectl. The kubeconfig file should be in the user’s home directory .kube folder for the setup to work as expected.
```sh
go run main.go
```
Then visit http://localhost:8080/docs/index.html to view the api list.



