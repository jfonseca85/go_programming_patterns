# Pattern Decorator

![decorator](decorator.png)

Neste Post iremos discutir sobre o **Pattern Decorator** este pattern pode facilmente montar algumas funções para outras funções,
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

Podemos ver que usamos uma função de ordem superior decorator().
Ao chamar, primeiro passamos a função Hello(), e depois ela retorna uma função anônima.
Essa função anônima não apenas executa seu próprio código, mas também chama a função passada Hello().

Essa codificação é semelhante à do Python, mas, infelizmente, o Go não suporta sintático @decorator. Então, é um pouco feio na chamada.
Claro, se você quiser tornar o código mais fácil de ler, você pode fazer isso:

````go
hello := decorator(Hello)
hello("Hello")
````
Vejamos outro exemplo e calculemos o tempo de execução:

````go
package main

import (
  "fmt"
  "reflect"
  "runtime"
  "time"
)

type SumFunc func(int64, int64) int64

func getFunctionName(i interface{}) string {
  return runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
}

func timedSumFunc(f SumFunc) SumFunc {
  return func(start, end int64) int64 {

    defer func(t time.Time) {
      fmt.Printf("--- Time Elapsed (%s): %v ---\n", 
          getFunctionName(f), time.Since(t))
    }(time.Now())

    return f(start, end)
  }
}

func Sum1(start, end int64) int64 {
  var sum int64
  sum = 0
  if start > end {
    start, end = end, start
  }
  for i := start; i <= end; i++ {
    sum += i
  }
  return sum
}

func Sum2(start, end int64) int64 {
  if start > end {
    start, end = end, start
  }
  return (end - start + 1) * (end + start) / 2
}

func main() {

  sum1 := timedSumFunc(Sum1)
  sum2 := timedSumFunc(Sum2)

  fmt.Printf("%d, %d\n", sum1(-10000, 10000000), sum2(-10000, 10000000))
}
````
Algumas coisas a serem observadas sobre o código acima:
1) Existem duas funções Sum, a função Sum1() é simplesmente um loop e a função Sum2() usa a fórmula de dados.
(Nota: início e fim podem ter números negativos)
2) O código usa a máquina de reflexão da linguagem Go para obter o nome da função.
3) A função de decorador é timedSumFunc()

Saída após a execução:

````shell
$ go run time.sum.go
--- Time Elapsed (main.Sum1): 3.557469ms ---
--- Time Elapsed (main.Sum2): 291ns ---
49999954995000, 49999954995000
````

## An example related to HTTP

Vejamos um exemplo relacionado de manipulação de solicitações HTTP.

Vejamos primeiro o código de um servidor HTTP simples.

````go
package main

import (
    "fmt"
    "log"
    "net/http"
    "strings"
)

func WithServerHeader(h http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        log.Println("--->WithServerHeader()")
        w.Header().Set("Server", "HelloServer v0.0.1")
        h(w, r)
    }
}

func hello(w http.ResponseWriter, r *http.Request) {
    log.Printf("Recieved Request %s from %s\n", r.URL.Path, r.RemoteAddr)
    fmt.Fprintf(w, "Hello, World! "+r.URL.Path)
}

func main() {
    http.HandleFunc("/v1/hello", WithServerHeader(hello))
    err := http.ListenAndServe(":8080", nil)
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}
````

O padrão decorador é usado acima no código WithServerHeader().
A função é um Decorator, que passa em um *http.HandlerFunc* e retorna uma versão reescrita.
O exemplo acima é relativamente simples, você pode adicionar um cabeçalho de resposta com WithServerHeader().

Portanto, podemos escrever várias dessas funções.
Conforme mostrado abaixo, alguns gravam cabeçalhos de resposta HTTP, alguns gravam cookies de autenticação,
alguns verificam cookies de autenticação e alguns log.

````go
package main

import (
    "fmt"
    "log"
    "net/http"
    "strings"
)

func WithServerHeader(h http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        log.Println("--->WithServerHeader()")
        w.Header().Set("Server", "HelloServer v0.0.1")
        h(w, r)
    }
}

func WithAuthCookie(h http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        log.Println("--->WithAuthCookie()")
        cookie := &http.Cookie{Name: "Auth", Value: "Pass", Path: "/"}
        http.SetCookie(w, cookie)
        h(w, r)
    }
}

func WithBasicAuth(h http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        log.Println("--->WithBasicAuth()")
        cookie, err := r.Cookie("Auth")
        if err != nil || cookie.Value != "Pass" {
            w.WriteHeader(http.StatusForbidden)
            return
        }
        h(w, r)
    }
}

func WithDebugLog(h http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        log.Println("--->WithDebugLog")
        r.ParseForm()
        log.Println(r.Form)
        log.Println("path", r.URL.Path)
        log.Println("scheme", r.URL.Scheme)
        log.Println(r.Form["url_long"])
        for k, v := range r.Form {
            log.Println("key:", k)
            log.Println("val:", strings.Join(v, ""))
        }
        h(w, r)
    }
}
func hello(w http.ResponseWriter, r *http.Request) {
    log.Printf("Recieved Request %s from %s\n", r.URL.Path, r.RemoteAddr)
    fmt.Fprintf(w, "Hello, World! "+r.URL.Path)
}

func main() {
    http.HandleFunc("/v1/hello", WithServerHeader(WithAuthCookie(hello)))
    http.HandleFunc("/v2/hello", WithServerHeader(WithBasicAuth(hello)))
    http.HandleFunc("/v3/hello", WithServerHeader(WithBasicAuth(WithDebugLog(hello))))
    err := http.ListenAndServe(":8080", nil)
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}
````

## Pipeline with multiple decorators
Em uso, é necessário empilhar funções camada por camada, o que não parece ser muito bonito,
se forem necessários mais decoradores, o código ficará mais feio. Bem, podemos refatorar um pouco.

Ao refatorar, precisamos primeiro escrever uma função utilitária - para percorrer e chamar cada decorador:

````go
type HttpHandlerDecorator func(http.HandlerFunc) http.HandlerFunc

func Handler(h http.HandlerFunc, decors ...HttpHandlerDecorator) http.HandlerFunc {
    for i := range decors {
        d := decors[len(decors)-1-i] // iterate in reverse
        h = d(h)
    }
    return h
}
````
Então, podemos usá-lo como abaixo.
````go
http.HandleFunc("/v4/hello", Handler(hello,
                WithServerHeader, WithBasicAuth, WithDebugLog))
````

Este código não seria mais fácil de ler? A função de pipeline também sai.


