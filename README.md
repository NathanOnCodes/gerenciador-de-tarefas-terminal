# Gerenciador de Tarefas em Go

### Flags:

O programa aceita as seguintes flags:
* `-add`: Adiciona uma nova tarefa.
* `-del`: Deleta uma tarefa existente.
* `-complete`: Marca uma tarefa como concluída.
* `-list`: Lista todas as tarefas existentes.

Primeiro execute o comando go build -o "nome que desejar, exemplo = gerenciador-de-tarefas" main.go
você precisa ter a linguagem go instalada, caso não tenha siga esse tutorial: <a href="https://go.dev/doc/install"> tutorial de instalação </a>
<br>

o comando -add, sua sintaxe: 
```
$ ./gerenciador-de-tarefas -add "Comprar leite"

```
<br>

o comando -del, sua sintaxe: 
```
$ ./gerenciador-de-tarefas -del 1

```
<br>

o comando -complete, sua sintaxe: 
```
$ ./gerenciador-de-tarefas -complete 1

```
<br>

o comando -list, sua sintaxe: 
```
$ ./gerenciador-de-tarefas -list

```

## Cross-compilação

Este projeto foi feito para ser executado em diferentes plataformas, por isso é necessário compilar o código para cada arquitetura desejada. sendo assim, não deixei um executável pré-compilado.
