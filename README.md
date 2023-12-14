### SMS Campaign Manager

The SMS Campaign Manager (`sms-campaign-manager`) is a microservice for managing `Campaigns`, `Segments`, and `Templates` for an SMS messaging service. This application utilizes **MongoDB** and **gqlgen**

#### Structure

```
sms-campaign-manager/
├── api/ # GraphQL API definition and resolvers
│ ├── campaigns/ # Campaign-related business logic
│ ├── segments/ # Segment-related business logic
│ └── templates/ # Templates-related business logic
├── config/ # Configuration files & database connection
├── graph
│ ├── generated # A package containing generated runtime code
│ │ └── generated.go
│ ├── model # A package for all graph models, generated or custom
│ │ └── models_gen.go
│ ├── resolver.go # The root graph resolver type. Not regenerated
│ ├── schema.graphqls # Schema definition file
│ └── schema.resolvers.go # the resolver implementation for schema.graphql
└── server.go # The entry point to the application
├── gqlgen.yml # The gqlgen config file
├── .gitignore # gitignore file
└── README.md # Documentation
```

## Usage

```
go run server.go
```

### GraphQL

#### gqlgen

```
go run github.com/99designs/gqlgen init
```

```
go run github.com/99designs/gqlgen generate
```

## Deployment

### Local

Start `minikube`

```
minikube start
```

Direct `minikube` to use the `docker` env. Any `docker build ...` commands after this command is run will build inside the `minikube` registry and will not be visible in Docker Desktop. `minikube` uses its own docker daemon which is separate from the docker daemon on your host machine. Running `docker images` inside the `minikube` vm will show the images accessible to `minikube`

```
eval $(minikube docker-env)
```

```
docker build -t sms-campaign-manager-image:latest .
```

#### Environment Variables (if needed)

```
kubectl create secret generic mongodb-secret --from-env-file=./.env
```

```
kubectl apply -f ./k8s/sms-campaign-manager.deployment.yaml
```

```
kubectl apply -f ./k8s/sms-campaign-manager.service.yaml
```

```
kubectl get deployments
```

```
kubectl get pods
```

```
minikube service sms-campaign-manager-service
```

After running the last comment the application will be able to be accessed in the browser at the specified port that `minikube` assigns.

#### Troubleshooting

```
minikube ssh 'docker images'
```

```
kubectl logs <pod-name>
```

```
kubectl logs -f <pod-name>
```
