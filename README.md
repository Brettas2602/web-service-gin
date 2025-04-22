## 🧠 Sobre este Projeto

Este projeto é uma **API REST** escrita em **Go (Golang)** utilizando o framework **Gin**. Ele foi criado como exemplo didático para alunos de Engenharia de Software que estão começando a aprender sobre:

- Como construir uma API;
- Arquitetura **MVC** adaptada para Go;
- O uso de **ponteiros** e **receivers**;
- O papel de cada camada da aplicação (controller, usecase, repository).

---

## 📚 O que é uma API?

Uma **API** (Application Programming Interface) é um conjunto de regras que define como sistemas se comunicam. No caso desta aplicação, a API permite criar, buscar, atualizar e deletar produtos via requisições HTTP.

---

## 🧱 Arquitetura usada: MVC + UseCase

O padrão utilizado é uma variação da arquitetura **MVC (Model-View-Controller)**, adaptada para aplicações em Go. Também incluímos a camada **UseCase**, que separa a lógica de negócio.

**Camadas:**

| Camada        | Função                                                                 |
|---------------|------------------------------------------------------------------------|
| `model`       | Define as estruturas de dados (ex: `Product`).                         |
| `controller`  | Lida com as requisições HTTP (entrada e saída da API).                |
| `usecase`     | Contém a lógica de negócio (ex: criar um produto).                    |
| `repository`  | Faz a comunicação com o banco de dados.                               |

---

## 🧭 O que são **Receivers**?

Em Go, uma **função com receiver** é como um "método" de uma struct. Exemplo:

```go
func (pc *ProductController) CreateProduct(c *gin.Context)
```

Isso quer dizer que `CreateProduct` é um método da struct `ProductController`. É como em POO, só que explícito.

---

## 📌 O que são **ponteiros**?

Ponteiros são formas de acessar e modificar diretamente o valor original de uma variável, ao invés de copiar seus dados.

```go
func criarProduto(p *Product)
```

- O `*Product` indica que estamos usando um **ponteiro**.
- Isso economiza memória e permite modificar o valor original.

---

## 📦 Explicando cada camada e seus arquivos

### 🔸 `model/product.go`

Define a estrutura básica do objeto que será manipulado na API.

```go
type Product struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}
```

---

### 🔸 `controller/product_controller.go`

Responsável por receber as requisições HTTP, validar os dados e retornar as respostas. Não contém lógica de negócio nem acesso ao banco.

```go
func (pc *ProductController) CreateProduct(c *gin.Context)
```

- **Receiver:** `(pc *ProductController)` é como "dizemos" que essa função pertence à struct `ProductController`.
- **Parâmetro `*gin.Context`**: contém os dados da requisição HTTP e métodos para responder ao cliente.

---

### 🔸 `usecase/product_usecase.go`

Aqui fica a **lógica de negócio**. Exemplo: se for necessário validar um preço, aplicar desconto, ou verificar se um produto já existe — tudo isso seria feito aqui.

```go
func (pu *ProductUsecase) CreateProduct(product *model.Product) (*model.Product, error)
```

- Recebe um ponteiro de `Product`, o que evita cópias desnecessárias.
- Retorna também um ponteiro, permitindo que possamos retornar `nil` em caso de erro.

---

### 🔸 `repository/product_repository.go`

Faz o acesso ao **banco de dados**. Cada função aqui executa um comando SQL específico.

```go
func (pr *ProductRepository) CreateProduct(product *model.Product) (int, error)
```

- Executa a `INSERT INTO` no banco.
- Retorna o ID gerado para o produto.

---

## ✅ Exemplo do fluxo completo

Quando você faz um `POST /products` com um JSON de produto:

1. O **controller** recebe os dados da requisição.
2. Ele chama o **usecase**, passando os dados do produto.
3. O **usecase** executa a lógica de criação.
4. Ele chama o **repository**, que insere no banco de dados.
5. O ID gerado volta para o **usecase**, que atualiza o produto.
6. Por fim, o **controller** envia a resposta para o cliente.

---

## 🛠 Requisitos para rodar

- Go 1.20 ou superior
- PostgreSQL
- Gin framework
