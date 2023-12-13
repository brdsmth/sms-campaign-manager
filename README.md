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
