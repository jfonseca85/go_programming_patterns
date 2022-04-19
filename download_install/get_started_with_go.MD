# Tutorial: Get started with Go

<p align="center"><img src="assets/images/get-started-hello-world.png?raw=true" width="600" height="480"></p>

## Conteúdo
- [Pré-requisitos](#pr-requisitos)
- [Install Go](#install-go)
- [Nosso Hello World](#nosso-hello-world)

Neste tutorial, iremos fazer uma inicialização para Programação em Go. Passando por estas etapas:

- Instalação do Go ( Caso ainda não tenha feito)
- Escrevendo um exemplo simples *"Hello, world!"*
- Usando o comando go para executar nosso código

> **Nota:** Para outros tutoriais, clique neste [link](https://go.dev/doc/tutorial/index.html)

## Pré-requisitos

- Experiência prévia em programação. Usamos um código bem simples, porém o conhecimento sobre funções orá ajudar.
- Uma ferramenta ára editar o código. Podemos usar qualquer editor de texto. A maioria dos editores tem um bom suporte
para o Go. Os mais populares são VSCode (gratuito), Goland(pago) e Vim (gratuito)
- Um terminal de comando. Go funciona bem em qualquer terminal no Linux, Mac e , Windows (cmd e PowerShell)

## Install Go

Neste post mostramos fazer o [Download e instalação](install.md) do Go

## Nosso Hello World

Começando bem simples, com o *Hello, World!"

### Abra seu terminal favorito e crie uma pasta chamada *tutoriais-golang\hello-world*

````shell
mkdir tutoriais-golang
mkdir tutoriais-golang\hello-world
cd tutoriais-golang\hello-world
````

### Habilitando o rastreamento de dependências

Quando o código em Go importa outros módulos. O gerenciamento é feito através do próprio código fonte.
Esse módulo é definido no **go.mod** que escanea e fornece os pacotes. Devemos manter o arquivo **go.mod** dentro do nosso projeto
inclusive no repositório git.

Para bahilitar o rastreamento de dependência devemos executar este comando

````shell
go mod init *<meu-modulo>*
````

Normalmente, inserimos o mesmo nome do repositório git do nosso módulo. Por exemplo, o caminho do módulo pode ser github.com/mymodule.
Desde modo você consegue publicar seu modulo para outros devs usarem. COnsulte esta documentação para saber mais 
sobre [Gerenciamento de dependências](https://go.dev/doc/modules/managing-dependencies#naming_module)

Para o nosso tutorial, vamos usar assim example/hello

````shell
$ go mod init example/hello
go: creating new go.mod: module example/hello
````

### Criando o *hello.go*

Crie um arquivo **hello.go** no seu editor de texto para escrever nosso código.

Cole o código abaixo dentro do arquivo **hello.go**

````go
//a package is a way to group functions, and it's made up of all the files in the same directory
package main

import "fmt" //which contains functions for formatting text, including printing to the console.

//main - Implement a main function to print a message to the console
//A main function executes by default when you run the main package.
func main() {
    fmt.Println("Hello, World!")
}
````
### Executando no código

````shell
$ go run .
Hello, World!
````

Chegamos ao final do nosso tutorial mostrando como fazer um código simples usando o Golang.
