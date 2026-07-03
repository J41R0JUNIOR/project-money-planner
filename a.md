
ESPECIFICAÇÃO DO SISTEMA — GERENCIADOR FINANCEIRO PESSOAL (SERVERLESS)

1. VISÃO GERAL
Objetivo:
Desenvolver um sistema pessoal para controle financeiro que permita:
- Registrar entradas e saídas
- Visualizar saldo atual
- Projetar saldo futuro
- Simular decisões financeiras

Problema:
“Posso gastar esse dinheiro agora sem comprometer meu futuro?”

Escopo (MVP):
Inclui:
- Cadastro de transações
- Consulta de histórico
- Projeção de saldo
- Simulação de gastos

Não inclui:
- Integração bancária
- Multiusuário
- Notificações
- UI avançada

2. ARQUITETURA
Cliente → API Gateway → Lambda (Go) → DynamoDB

3. MODELO DE DADOS
Transaction:
- ID
- Type (INCOME | EXPENSE)
- Amount
- Date
- Description
- IsRecurring

Projection:
- Date
- Balance

DynamoDB:
PK = USER#<id>
SK = DATE#<timestamp>

4. FLUXOS

Fluxo principal:
Início → Receber transação → Validar → Salvar → Buscar transações → Ordenar → Calcular saldo → Retornar → Fim

Fluxo de simulação:
Início → Receber simulação → Buscar transações → Adicionar fictícia → Ordenar → Calcular → Retornar → Fim

5. CASOS DE USO

UC-01 Registrar Transação:
Usuário envia dados → Sistema valida → Sistema salva

UC-02 Listar Transações:
Usuário solicita → Sistema consulta → Retorna lista

UC-03 Obter Projeção:
Sistema busca → Ordena → Calcula → Retorna

UC-04 Simular Gasto:
Usuário envia → Sistema simula → Retorna resultado

6. ENDPOINTS
POST /transactions
GET /transactions
GET /projection
POST /simulate

7. REGRAS DE NEGÓCIO
- Entradas somam
- Saídas subtraem
- Ordenação por data obrigatória
- Simulação não persiste
- Projeção sempre recalculada

8. ALGORITMO
Saldo acumulado:
balance += INCOME
balance -= EXPENSE

9. EVOLUÇÕES FUTURAS
- Recorrência automática
- Categorias
- Alertas
- Dashboard
- Machine learning

10. RISCOS
- Datas inconsistentes
- Crescimento de dados
- Complexidade de recorrência

11. ROADMAP
Fase 1: lógica em Go
Fase 2: API
Fase 3: DynamoDB
Fase 4: AWS deploy

CONCLUSÃO:
Sistema focado em previsão de fluxo financeiro pessoal.
