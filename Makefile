# Mailer Micro Service Makefile

# lint: runs a golangci-lint with the same settings as in the CI.
lint:
	golangci-lint run ./...

# check: executes a static check.
check:
	staticcheck ./...

# test: executes a test suite.
test:
	go test ./...

# build: builds application.
build:
	env GOOS=linux GOARCH=amd64 go build -o mailer-micro-service main.go

# run: starts an app.
run:
	go run main.go

# ks: starts an app in the kubernetes cluster
ks:
	cd kubernetes && kubectl apply -f app-service.yaml,postgres-service.yaml,app-deployment.yaml,postgres-deployment.yaml,postgres-data-persistentvolumeclaim.yaml,postgres-configmap.yaml

# kd: stops the app
kd:
	cd kubernetes && kubectl delete -f app-service.yaml,postgres-service.yaml,app-deployment.yaml,postgres-deployment.yaml,postgres-data-persistentvolumeclaim.yaml,postgres-configmap.yaml

# mkenv: allow to work in a minikube docker env
mkenv:
	eval $(minikube docker-env)

# mks: starts a minikube
mks:
	minikube start

# mkd: downs a minikube
mkd:
	minikube stop

# kbpa: starts a kubernetes proxy for app
kbpa:
	kubectl port-forward service/app 8080:8080

# kbpa: starts a kubernetes proxy for postgres
kbpp:
	kubectl port-forward service/postgres 5432:5432

