# Pattern Decorator

![decorator](decorator.png)

Neste Post iremos conversar sobre o Pattern Decorator este partnnern pode facilmente montar algumas funções para outras funções,
o que pode tornar seu código mais simples, e também pode fazer algumas "pequenas funções".
O código é mais reutilizável, de modo que as funções no código podem ser montadas livremente como brinquedos Lego.

O objetivo principal deste artigo é mudar a maneira de pensar de programação procedural para programação funcional.
Levar mais pessoas a jogar programação funcional. Portanto, o modo de programação do decorador da linguagem Go é na verdade o modo de programação funcional.

No entanto, deve-se notar que a linguagem Go não tem muito "sugar", e é uma linguagem estática fortemente tipada sem uma máquina virtual.
Portanto, é impossível obter um código decorador elegante como Java e Python. Claro, talvez seja eu que sou ignorante,
se você souber de mais maneiras de escrever, por favor me avise. Desde já, obrigado.

## Content
1) [Simple example](#simple-example)
2) [An example related to HTTP](#an-example-related-to-http)
3) [Pipeline with multiple decorators](#pipeline-with-multiple-decorators)
4) [Generic decorator](#generic-decorator)

## Simple example

Vejamos primeiro um exemplo:

````go
package main

import "fmt"

func decorator(f func(s string)) func(s string) {

    return func(s string) {
        fmt.Println("Started")
        f(s)
        fmt.Println("Done")
    }
}

func Hello(s string) {
    fmt.Println(s)
}

func main() {
        decorator(Hello)("Hello, World!")
}
````

## An example related to HTTP
## Pipeline with multiple decorators
## Generic decorator
