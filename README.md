# Graph IFC

Projeto para encontrar a distânci entre caminhos em um grafo representando locais e distâncias no IFC Videira. Este projeto utiliza Go e Gin para construir uma API que oferece funcionalidades de encontrar o caminho mais curto e listar todos os caminhos possíveis entre dois pontos dados.

---

## Requisitos

- Go 1.22 ou superior

---

## Instalação e Execução

1. Clone o repositório:

```bash
git clone https://github.com/Renan-Parise/graph-ifc.git
```

2. Entre na pasta do projeto:

```bash
cd graph-ifc
```

3. Instale as dependências:

```bash
go mod tidy
```

4. Execute o projeto:

```bash
go run main.go
```

O servidor estará rodando em `127.0.0.1:8181`.

## Endpoints

### `POST /findpaths`

Encontra o caminho mais curto entre dois pontos e retorna a distância e o caminho percorrido, além de listar todos os caminhos possíveis entre os dois pontos.

#### Payload

```json
{
    "from": "origem",
    "to": "destino"
}
```

#### Exemplo

```bash
curl --location '127.0.0.1:8181/findpaths' \
--header 'Content-Type: application/json' \
--data '{
    "from": "ginásio",
    "to": "guarita"
}'
```

### Importante!

- O código ignora letras maiúsculas e minúsculas, então `ginásio` e `Ginásio` são considerados iguais.
- Caso você insira um local que não exista no grafo, o servidor retornará um erro 500 - Internal Server Error com a mensagem `source node not found`.
- Para o desenvolvimento deste código, foi aplicado os algoritmos de Dijkstra e DFS.