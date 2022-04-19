# Interface Programming

<p align="center"><img src="interface_programming.png?raw=true" width="300" height="150"></p>

Neste artigo, abordarei algumas técnicas básicas e pontos-chave dos padrões de programação Go,
que tornarão mais fácil para você dominar a programação Go.

Hoje iremos conversar sobre o Parttern Interface Programming usando *Golang*.
A seguir Mostramos dois exemplos simples que imprimem os dados da struct Person.
No primeiro exemplo usamos uma função e no segundo usamos "Member Function" *Print()*

```go

type Person struct {
    Name string
    Sex  string
    Age  int
}

//PrintPerson - Function that prints data as Person structure
func PrintPerson(p *Person) {
fmt.Printf("Name=%s, Sex=%s, Age=%d\n",
p.Name, p.Sex, p.Age)
}

//Print - Member function ( Receiver ) that prints data as Person structure
func (p *Person) Print() {
    fmt.Printf("Name=%s, Sex=%s, Age=%d\n",
    p.Name, p.Sex, p.Age)
}

func main() {
    var p = Person{
    Name: "Rafael Oliveira",
    Sex:  "Male",
    Age:  35,
    }

    PrintPerson(&p)
    p.Print()
}
```
Qual maneira prefere? Na linguagem *Go* a maneira de usar a "Member function" é chamada "Receiver", pois originalmente é *Person* ficando fortemente acoplado.
O Mais importante neste tipo de método podemos realizar a Interface Programming permitindo a abstração usada principalmente em "polimorfismo".
Aqui quero conversar sobre Interface Programming em Go.
Primeiro, vamos dar uma olhada no seguinte código:

```go
package main

import "fmt"

//Country - Structure that encapsulates the Country name
type Country struct {
	Name string
}

//City - Structure that encapsulates the City name
type City struct {
	Name string
}

//Printable - A Country and City interface Both implement PrintS() interface methods and produce themselves.
type Printable interface {
	PrintStr()
}

//PrintStr -Applying polymorphism in the PrintStr function to the Country structure
func (c Country) PrintStr() {
	fmt.Println(c.Name)
}

//PrintStr -Applying polymorphism in the PrintStr function to the PrintStr structure
func (c City) PrintStr() {
	fmt.Println(c.Name)
}

//main function that runs our scenario
func main() {
	c1 := Country{"São Paulo"}
	c2 := City{"Belém"}
	c1.PrintStr()
	c2.PrintStr()
}

```

Podemos notar que esta implementação usa a interface *Printable* de Country e City que implementam a função
*PrintStr()*. No entanto, os códigos são iguais. Vamos melhorar esta codificação?

Podemos usar "struct embedding" para fazer isso, conforme mostrado no código a seguir:

```go

//WithName - Struct embedding
type WithName struct {
	Name string
}

//Country - Structure that encapsulates the Country name
type Country struct {
	WithName
}

//City - Structure that encapsulates the City name
type City struct {
	WithName
}

//Printable - A Country and City interface Both implement PrintS() interface methods and produce themselves.
type Printable interface {
	PrintStr()
}

//PrintStr -Applying polymorphism in the PrintStr function to the WithName structure
func (w WithName) PrintStr() {
	fmt.Println(w.Name)
}

//main function that runs our scenario
func main() {
	c1 := Country{WithName{Name: "São Paulo"}} //messy startup
	c2 := City{WithName{Name: "Belém"}}       //messy startup
	c1.PrintStr()
	c2.PrintStr()
}
```

No entanto, um problema aparace com a inclusão do *WithName*, é que a inicialização fica um pouco confusa.
Então, temos uma maneira melhor? Aqui está a solução.

````go
//Country - Structure that encapsulates the Country name
type Country struct {
	Name string
}

//City - Structure that encapsulates the City name
type City struct {
	Name string
}

//Stringable - A Country and City interface Both implement PrintS() interface methods and produce themselves.
type Stringable interface {
	ToString() string
}

//ToString - Applying polymorphism in the ToString function to the Country structure
func (c Country) ToString() string {
	return "Country = " + c.Name
}

//ToString - Applying polymorphism in the ToString function to the City structure
func (c City) ToString() string {
	return "City = " + c.Name
}

//PrintStr - Function that receives the Stringable interface, which can be used by Country and City implementations
func PrintStr(p Stringable) {
	fmt.Println(p.ToString())
}

//main - function that runs our scenario
func main() {
	d1 := Country{"USA"}
	d2 := City{"Los Angeles"}
	PrintStr(d1)
	PrintStr(d2)
}
````

Podemos notar no código acima o uso da interface *Stringable* , usamos esta interface para desacoplar a o tipo de negócio ( City ou Country )
da lógica de controle *PrintStr*. Uma vez, que esta função recebe como paramêtro a interface *Stringable*, podendo ser usada pelas implementações Country e City.

Existem muitos exemplos desse padrão de programação na biblioteca padrão do Go, io.Readsendo ioutil.ReadAllo play de e , 
onde io.Read é uma interface, e você precisa implementar um de seus Read(p []byte) (n int, err error) métodos de interface.
Desde que atenda a essa escala, ele pode ser usado por ioutil.ReadAll este método.
Esta é a regra de ouro da programação orientada a objetos - "Programe para uma interface, não para uma implementação".

## Interface Integrity Check

Além disso, podemos ver que o programador da linguagem Go não verifica rigorosamente se um objeto implementa todos 
os métodos de uma interface, conforme mostra o exemplo a seguir:

```go
//Shape - Shape - Interface to abstract the polygon type
type Shape interface {
    Sides() int
    Area() int //Method not implemented
}

//Square - Structure that implements a Shape interface
type Square struct {
    len int
}
//Slices - Get the number of sides of the square
func (s* Square) Sides() int {
    return 4
}

//main - function that runs our scenario
func main() {
    s := Square{len: 5}
    fmt.Printf("%d\n",s.Sides())
}
```

Podemos ver que todos os métodos da interfaceSquare não estão implementados . 
Embora o programa possa rodar, a forma de programação não é rigorosa. Se precisarmos aplicar todos os métodos da interface *Shape*,
o que devemos fazer?

Existe uma prática padrão no círculo de programação da linguagem Go:

````go
var _ Shape = (*Square)(nil)
````

Declare uma _variável (ninguém a usa), que converterá um nil (ponteiro nulo)  de Square para em Shape, de modo que,
se o método de interface relevante não for implementado, o compilador informará um erro:

>cannot use (*Square)(nil) (type *Square) as type Shape in assignment: *Square does not implement Shape (missing Area method)
