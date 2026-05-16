# Introdução a CI com GitHub Actions: build, testes e empacotamento Docker de uma aplicação Go

- Proposto e criado por: Pietro Piva Vieira
- Revisado por: Prof. Rodrigo Mansilha

Este repositório é o template para a atividade introdutória de CI (Continuous Integration).

A aplicação Go já está pronta porque o foco do exercício é completar a pipeline no GitHub Actions.

## Requisitos para a atividade

Para realizar testes locais, é necessário:

1. Instalar e configurar Go 1.26;
2. Instalar e configurar Docker;

## Objetivo

Completar o arquivo [`ci.yml`](.github/workflows/ci.yml) para que a pipeline:

1. compile a aplicação Go;
2. execute os testes automatizados.

Cada uma dessas ações deve ser feita com apenas um comando.

## O que já vem pronto

- uma aplicação Go mínima em [`cmd/app/main.go`](cmd/app/main.go);
- uma lógica simples em [`internal/greeting/greeting.go`](internal/greeting/greeting.go);
- um teste automatizado em [`internal/greeting/greeting_test.go`](internal/greeting/greeting_test.go);
- a estrutura inicial do workflow em [`.github/workflows/ci.yml`](.github/workflows/ci.yml).

## O que você deve fazer

Antes de iniciar a atividade, baixe este repositório como `.zip` disponibilizado pelo próprio GitHub e crie o seu próprio repositório (público) no GitHub. Faça um commit com esta estrutura inicial.

### Prática

Edite o arquivo [`.github/workflows/ci.yml`](.github/workflows/ci.yml) e substitua os `TODOs` por passos válidos do GitHub Actions para:

1. buildar a aplicação;
2. rodar os testes.

### Questionário

Crie um arquivo chamado `ANSWERS.md` e responda as seguintes perguntas:

1. Explique como a _pipeline_ é disparada no GitHub Actions. Na sua resposta, cite especificamente o que você configurou no arquivo `ci.yml`.
2. O que é um _runner_ no GitHub Actions e qual o seu papel na execução da _pipeline_?
3. Qual a diferença entre _buildar_ a aplicação inteira como binário e _buildar_ a imagem Docker?
4. Por que usar Docker em uma _pipeline_ CI pode ser útil?
5. Altere temporariamente o código para fazer um teste falhar.
  - O que aconteceu com o _pipeline_?
  - Em qual etapa ele falhou?
  - Anexe um _print_ do erro.

## Como validar localmente

Antes de subir sua solução, você pode testar os mesmos comandos no seu terminal:

```bash
go build ./...
go test ./...
```

## Critério de sucesso

A atividade principal está concluída quando:

- O _workflow_ executa sem erro no GitHub Actions;
- A aplicação compila;
- Os testes passam;
- Um _print_ da execução da _pipeline_ com sucesso é incluído no repositório.

## Desafios bônus

### Docker

O repositório também inclui um [Dockerfile](Dockerfile) pronto. Como desafio optativo, você pode adicionar uma nova etapa ao workflow para buildar a imagem Docker da aplicação.

Objetivo do bônus:

- Buildar a imagem Docker usando o Dockerfile fornecido;
- Garantir que o build acontece após o sucesso do _build_ e dos testes;

### CI em Pull Requests

Como segundo desafio optativo, você pode alterar novamente o arquivo [`ci.yml`](.github/workflows/ci.yml) para fazer com que a _pipeline_ execute também em _Pull Requests_.

Objetivo do bônus:

- Modifique o _workflow_ para também executar em eventos de `pull_request`;
- Crie uma _Pull Request_ no seu repositório para validar a execução da _pipeline_ (não é necessário realizar o `merge`, apenas demonstrar a execução).

## Dica

Se a pipeline falhar, utilize os logs do GitHub Actions para identificar o problema. Eles indicam exatamente qual etapa falhou e por quê.

## Materiais importantes

- [Documentação - GitHub Actions](https://docs.github.com/pt/actions);

## Materiais extras

Deixo aqui, materiais extras para que vocês possam entender mais das práticas de integração contínua. Não são obrigatórios, mas são conteúdos bacanas 😁.

- [Leitura: Continuous Integration - Martin Fowler (em inglês)](https://martinfowler.com/articles/continuousIntegration.html);
- [Leitura: Continuous Integration - Atlassian](https://www.atlassian.com/br/continuous-delivery/continuous-integration);
- [Vídeo: Integração Contínua - Código Fonte TV](https://www.youtube.com/watch?v=nI3IjYcBGiU).

## Feedback

Seu feedback é importante para mim! Se possível, por favor, encaminhe sua opinião para [meu e-mail institucional](mailto:pietrovieira.aluno@unipampa.edu.br).
