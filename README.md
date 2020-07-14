## Sample repository for the demo of Google Cloud Build and OpenFaas running on GKE

>This repo has also been tested for deploying OpenFaas function from GCB to OpenFaas running on minikube with Docker Hub as the container registry instead of Google Cloud's Container Registry and exposed with Inlets-pro, the possibilities are endless

Repo contains:

./hello-world/ -> Contains the handler for the function

./hello-world.yml -> YAML file containing the config generated by faas-cli new

.gitignore -> Files to ignore for version control

README.md -> The file you're reading right now.