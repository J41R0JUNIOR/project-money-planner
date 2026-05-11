# Project Money Planner

Sistema em Go para planejamento financeiro pessoal com foco em arquitetura serverless na AWS.

## Objetivo do MVP

- Registrar transacoes (entrada e saida)
- Consultar historico
- Projetar saldo
- Simular impacto de gastos sem persistir a simulacao

## Arquitetura alvo

Cliente -> API Gateway -> Lambda (Go) -> DynamoDB (single-table)

## Estado atual do projeto

- Modelos de dominio base organizados em `src/domain`
- Modelos de persistencia DynamoDB organizados em `src/domain/dynamodb`
- Interfaces de repositorio em `src/repositories`
- Implementacoes DynamoDB em `src/repositories/dynamodb_repositories`
- Infra DynamoDB (client e criacao de tabela) em `src/infra/aws`
- Entry point de API/Lambda ainda nao implementado (`main.go` esta vazio)

## Estrutura atual

```text
.
├── go.mod
├── go.sum
├── main.go
└── src
    ├── domain
    │   ├── models.go
    │   └── dynamodb
    │       └── models.go
    ├── repositories
    │   ├── UserRepository.go
    │   ├── TransactionRepository.go
    │   ├── ProjectionRepository.go
    │   └── dynamodb_repositories
    │       ├── UserDynamoRepository.go
    │       ├── TransactionDynamoRepository.go
    │       ├── ProjectionDynamoRepoitory.go
    │       └── mappers.go
    └── infra
        └── aws
            ├── DynamoDBClient.go
            └── DynamoDBCreateTable.go
```

## Modelagem DynamoDB (single-table)

- Partition Key: `PK`
- Sort Key: `SK`
- Convencao:
  - `PK = USER#<id>`
  - `SK = PROFILE` para perfil de usuario
  - `SK = DATE#<timestamp>#TX#<id>` para transacoes
  - `SK = DATE#<yyyy-mm-dd>#PROJECTION` para projecoes

## Dependencias AWS

Atualmente a infra DynamoDB do projeto usa AWS SDK for Go v1 (`github.com/aws/aws-sdk-go`).

## Como validar localmente

```bash
go mod tidy
go build ./...
```

## Proximos passos recomendados

1. Implementar os metodos dos repositorios DynamoDB (hoje estao com `panic("unimplemented")`)
2. Criar camada de servico para saldo/projecao/simulacao
3. Implementar handler Lambda em `main.go` com os endpoints do MVP
4. Adicionar IaC (SAM/CloudFormation/Terraform) para deploy serverless

## Observacoes

- O nome do arquivo `ProjectionDynamoRepoitory.go` esta com typo em "Repoitory".
- Se optar por SDK v2, convem migrar `src/infra/aws` para manter consistencia.
