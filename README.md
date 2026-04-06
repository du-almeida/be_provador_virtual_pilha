# 👗 Provador Virtual: Engenharia de Dados com LIFO (Stack)

Este projeto é uma implementação prática de **Estrutura de Dados Lineares**, simulando a lógica de um provador de roupas de um e-commerce de moda. O objetivo é demonstrar o domínio técnico sobre **Pilhas (Stacks)**, **Orientação a Objetos** e **Gestão de Memória** utilizando a linguagem **Go**.

---

## 🧠 O Conceito: Por que uma Pilha?

No universo da moda, quando sobrepomos peças de roupa, a última peça vestida é obrigatoriamente a primeira a ser retirada para não danificar o look. Na computação, chamamos esse comportamento de **LIFO (Last-In, First-Out)**.

Neste projeto, o "Provador" não é apenas uma lista comum; ele é um **Objeto Encapsulado** que protege a integridade das regras de negócio do ateliê.

---

## 🛠️ O que foi aplicado (Ficha Técnica)

### 1. Estrutura de Dados: Stack (Pilha)
* **Operação Push (Adicionar):** Inserção de uma nova peça no topo do look com custo de processamento **O(1)** (Tempo Constante).
* **Operação Pop (Desfazer):** Remoção cirúrgica da última peça adicionada, também com custo **O(1)**, garantindo performance máxima independente do tamanho do guarda-roupa.

### 2. Orientação a Objetos em Go
* **Structs & Encapsulamento:** A lista de roupas (`pecasVestidas`) é privada. Isso impede que agentes externos alterem o estado do provador sem passar pelas travas de segurança.
* **Method Receivers:** Uso de ponteiros (`*ProvadorVirtual`) para manipular o objeto real na memória Heap, evitando cópias desnecessárias de dados.

### 3. Regras de Negócio e Segurança
* **Trava de Overflow:** Implementação de um limite máximo de **5 peças**. Se a cliente tentar exceder o limite, o sistema dispara um erro controlado, evitando desperdício de memória RAM.
* **Programação Defensiva:** Validação de estados vazios para impedir "Panics" (crashes) no servidor durante a operação de remoção.

---

## 💻 Como o código funciona (Lógica de Engenharia)

Para facilitar o entendimento de quem lê o código, utilizei o padrão de **Pseudo-Código Comentado**, onde cada instrução é explicada sob a ótica do "Por Quê":

```go
// POR QUE: Trava de segurança para evitar acúmulo infinito de dados.
// PSEUDO-CÓDIGO: Verifica se o tamanho já atingiu o limite de 5 peças.
if tamanho >= 5 {
    return errors.New("Limite atingido!")
}
```

---

## 🚀 Como Executar

### Executar Diretamente
```bash
cd "/be_provador_virtual_pilha"
go run .
```

### Compilar para Executável
```bash
go build .
```
Isso gera um executável `be_provador_virtual_pilha` (ou `.exe` no Windows) que pode ser rodado sem Go instalado.

---

## 📁 Estrutura de Arquivos

```
be_provador_virtual_pilha/
│
├── main.go                          # Código principal (sem comentários)
├── go.mod                           # Arquivo de módulo Go
├── README.md                        # Esta documentação
│
└── codigo_comentado_estudo/         # Pasta com versões educacionais
    ├── sem_interacao/
    │   └── provador_sem_interacao.go
    └── com_interacao/
        └── atelie_com_interacao.go  # Versão com menu comentado
```

---

## 📚 Visualizando a Pilha em Ação

A estrutura LIFO funciona assim:

```
[TOP]    ← Primeira a sair (Pop)
  🧥 Cinto
  👖 Calça
  👕 Camiseta
  🩱 Sutiã
[BASE]   ← Última a sair
```

**Sequência de Adição:**
1. Sutiã (1ª)
2. Camiseta (2ª)
3. Calça (3ª)
4. Cinto (4ª) ← Está no topo

**Com Desfazer:**
Remove Cinto → Remove Calça → Remove Camiseta → Remove Sutiã

---

## 🎯 Menu de Opções

| Opção | Ação | Saída Esperada |
|-------|------|---|
| 1 | Adicionar uma peça | `Sucesso! Você vestiu: [nome]` |
| 2 | Desfazer última peça | `✂️ Ação desfeita! Você tirou: [nome]` |
| 3 | Ver Look Completo | Lista de roupas em ordem (topo → base) |
| 4 | Sair | `👋🏾 Saindo do provador... Volte sempre!` |

---

## ⚙️ Limites do Sistema

- **Máximo de peças:** 5 itens
- **Mínimo para desfazer:** 1 peça
- **Validação:** Nome da peça não pode ser vazio