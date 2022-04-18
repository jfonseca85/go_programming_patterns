# Error Handling in Go

![error_Handling](error-handling.png)

O tratamento de erros sempre foi um problema que a programação deve enfrentar, se o tratamento de erros for bem-feito, 
a estabilidade do código será ótima. Diferentes linguagens possuem diferentes formas de lidar com ocorrências. 
O mesmo vale para o Go e, neste artigo, vamos discutir onde o Go dá errado, especialmente os enlouquecedores if err != nil.

As funções na linguagem Go suportam vários valores de retorno, portanto, a semântica de negócios (valores de retorno de negócios)
e a semântica de controle (valores de retorno de erro) podem ser diferenciadas na ‘interface’ de retorno.
Muitas funções na linguagem Go retornarão dois valores de result e err, então:

- O parâmetro é basicamente o parâmetro de entrada, e a interface de retorno separa o resultado do erro,
o que torna clara a semântica da interface da função;
- Além disso, se os parâmetros de erro na linguagem Go devem ser ignorados, eles precisam ser ignorados explicitamente,
e variáveis como _ são usadas para ignorá-los;
- Além disso, como o retornado error é uma interface (com apenas um método Error(), retornando um string),
você pode estender o tratamento de erros personalizado.

Além disso, se uma função retornar vários tipos diferentes error, você também poderá usar o seguinte método:

````go
if err != nil {
  switch err.(type) {
    case *json.SyntaxError:
      ...
    case *ZeroDivisionError:
      ...
    case *NullPointerError:
      ...
    default:
      ...
  }
}
````

Podemos ver que a forma de tratamento de erros na linguagem Go é essencialmente a verificação do valor de retorno,
mas também considera alguns dos benefícios das exceções - a extensão dos erros.

## Resource cleanup

Após ocorrer um erro, a limpeza de recursos precisa ser feita.
Diferentes linguagens de programação possuem diferentes modos de programação para limpeza de recursos:

- **Linguagem C** - usa *goto fail;* o método para limpar num local centralizado (para um artigo interessante, veja "Pensando no BUG de baixo nível da Apple")
- **Linguagem C++** - geralmente usa o *modo RAII*, através do modo proxy orientado a objetos,
entrega os recursos que precisam ser limpos para uma classe proxy e, em seguida, resolve no destruidor.
- **Linguagem Java** - a limpeza pode ser feita num bloco finally.
- **Go language** - use a palavra - *defer* chave para limpar.

Aqui está um exemplo de limpeza de recursos em Go:

````go
func Close ( c io. Closer ) {
    err := c.Close ()
    if err != nil {
        log. Fatal ( err )
    }
}
func main () {
    r, err := Open ( "a" )
    if err != nil {
        log.Fatalf ( "error opening ' a'\n" )
    }
    defer Close ( r ) // Use the defer keyword to close the file when the function exits. 
    r, err = Open ( "b" )
    if err != nil {
        log.Fatalf ( "error opening ' b'\n" )
    }
    defer Close ( r ) // Use the defer keyword to close the file when the function exits. 
}
````
## Error Check  Hell

Bem, quando se trata *if err !=nil* do , esse código pode realmente fazer as pessoas vomitarem.
Então, existe alguma boa maneira? Vamos primeiro olhar para o seguinte código de falha.

````go
func parse(r io.Reader) (*Point, error) {

    var p Point

    if err := binary.Read(r, binary.BigEndian, &p.Longitude); err != nil {
        return nil, err
    }
    if err := binary.Read(r, binary.BigEndian, &p.Latitude); err != nil {
        return nil, err
    }
    if err := binary.Read(r, binary.BigEndian, &p.Distance); err != nil {
        return nil, err
    }
    if err := binary.Read(r, binary.BigEndian, &p.ElevationGain); err != nil {
        return nil, err
    }
    if err := binary.Read(r, binary.BigEndian, &p.ElevationLoss); err != nil {
        return nil, err
    }
}
````

Para resolver este problema, podemos usar a forma de programação funcional, o seguinte exemplo de código:

````go
func parse(r io.Reader) (*Point, error) {
    var p Point
    var err error
    read := func(data interface{}) {
        if err != nil {
            return
        }
        err = binary.Read(r, binary.BigEndian, data)
    }

    read(&p.Longitude)
    read(&p.Latitude)
    read(&p.Distance)
    read(&p.ElevationGain)
    read(&p.ElevationLoss)

    if err != nil {
        return &p, err
    }
    return &p, nil
}
````

Podemos ver no código acima que redefinimos uma função extraindo o mesmo código usando Closure,
então muito *if err!=nil* do processamento foi removido. Mas haverá um problema, ou seja,
há uma *err* variável e uma função interna, que não parece muito limpa.

Então, podemos fazer um pouco mais limpo *bufio.Scanner()*:

````go
scanner := bufio.NewScanner(input)

for scanner.Scan() {
    token := scanner.Text()
    // process token
}

if err := scanner.Err(); err != nil {
    // process the error
}
````

A partir do código acima, podemos ver que scanner ao operar a E/S subjacente, não há case *if err !=nil* e há uma *scanner.Err()*
verificação após sair do loop. Parece usar a estrutura do caminho. Para emulá-lo, podemos refatorar nosso código para ficar assim:

Primeiro, defina uma struct e uma member function

````go
type Reader struct {
    r   io.Reader
    err error
}

func (r *Reader) read(data interface{}) {
    if r.err == nil {
        r.err = binary.Read(r.r, binary.BigEndian, data)
    }
}
````

Então, o nosso código pode ficar assim:

````go
func parse(input io.Reader) (*Point, error) {
    var p Point
    r := Reader{r: input}

    r.read(&p.Longitude)
    r.read(&p.Latitude)
    r.read(&p.Distance)
    r.read(&p.ElevationGain)
    r.read(&p.ElevationLoss)

    if r.err != nil {
        return nil, r.err
    }

    return &p, nil
}
````

Com a implementação acima, nossa " Fluent Interface" é fácil de manusear. Do seguinte modo:
````go
package main

import (
  "bytes"
  "encoding/binary"
  "fmt"
)

// The length is not enough, one less weight
var b = []byte {0x48, 0x61, 0x6f, 0x20, 0x43, 0x68, 0x65, 0x6e, 0x00, 0x00, 0x2c} 
var r = bytes.NewReader(b)

type Person struct {
  Name [10]byte
  Age uint8
  Weight uint8
  err error
}
func (p *Person) read(data interface{}) {
  if p.err == nil {
    p.err = binary.Read(r, binary.BigEndian, data)
  }
}

func (p *Person) ReadName() *Person {
  p.read(&p.Name) 
  return p
}
func (p *Person) ReadAge() *Person {
  p.read(&p.Age) 
  return p
}
func (p *Person) ReadWeight() *Person {
  p.read(&p.Weight) 
  return p
}
func (p *Person) Print() *Person {
  if p.err == nil {
    fmt.Printf("Name=%s, Age=%d, Weight=%d\n",p.Name, p.Age, p.Weight)
  }
  return p
}

func main() {   
  p := Person{}
  p.ReadName().ReadAge().ReadWeight().Print()
  fmt.Println(p.err)  // EOF Error
}
````
Acredito que deva entender essa técnica, mas seu cenário de uso só pode simplificar o tratamento de erros sob a 
operação contínua do mesmo objeto de negócios. Para vários objetos de negócios, vários métodos ainda *if err != nil* são necessários.

## packaging error

Por fim, mais uma coisa, precisamos encapsular o erro, em vez de devolvê-lo à camada superior secamente err, precisamos adicionar algum contexto de execução.

Normalmente, usamos *fmt.Errorf()* para fazer isso, por exemplo:

````go
if err != nil {
   return fmt.Errorf("something failed: %v", err)
}
````

Além disso, entre os desenvolvedores (Go), é uma prática mais comum envolver o erro em outro erro, preservando o conteúdo original:

````go
type authorizationError struct {
    operation string
    err error   // original error
}

func (e *authorizationError) Error() string {
    return fmt.Sprintf("authorization failed during %s: %v", e.operation, e.err)
}
````

Claro, uma maneira melhor seria através de um método de acesso padrão, desta forma, é melhor usar uma interface,
causer como o Cause() método implementado na interface, para expor o erro original para inspeção adicional:

````go
type causer interface {
    Cause() error
}

func (e *authorizationError) Cause() error {
    return e.err
}
````

A boa notícia aqui é que esse código não precisa mais ser escrito, existe uma biblioteca de erros de terceiros ( github.com/pkg/errors ),
para esta biblioteca, posso ver a sua existência onde quer que eu vá, então, é basicamente o padrão de fato. O exemplo de código é o seguinte:

````go
import "github.com/pkg/errors"

//wrapper of error
if err != nil {
    return errors.Wrap(err, "read failed")
}

// Interface Cause
switch err := errors.Cause(err).(type) {
case *MyError:
    // handle specifically
default:
    // unknown error
}
````

## Reference article

- [Golang Error Handling lesson by Rob Pike](http://jxck.hatenablog.com/entry/golang-error-handling-lesson-by-rob-pike)
- [Errors are values](https://blog.golang.org/errors-are-values)