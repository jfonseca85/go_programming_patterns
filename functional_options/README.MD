# Pattner Functional Options
![functional_options](functional_options.png)

Neste artigo, vamos discutir o padrão de programação Functional Options. Este é um caso de aplicação de programação funcional,
e as habilidades de programação também são ótimas, sendo atualmente o padrão de programação mais popular na linguagem Go.
No entanto, antes de podermos discutir formalmente esse padrão, precisamos primeiro dar uma olhada no tipo de problema que está sendo resolvido.

## Content
1) [configuration options problem](#configuration-options-problem)
2) [Configure Object Schema](#configure-object-schema)
3) [Builder mode](#builder-mode)
4) [Functional Options](#functional-options)
5) [Reference documentation](#reference-documentation)

## Configuration options problem

Durante a programação, muitas vezes precisaremos configurar um objeto (ou entidade de negócios).
Por exemplo, a seguinte entidade comercial (observe que este é apenas um exemplo):

````go
type Server struct {
    Addr     string
    Port     int
    Protocol string
    Timeout  time.Duration
    MaxConns int
    TLS      *tls.Config
}
````

Nesta strcut Server , podemos ver:
- Para ter um endereço IP de escuta *Addr* e número de porta *Port*, essas duas opções de configuração são necessárias
(claro, o endereço IP e o número da porta podem ter valores padrão, quando o usamos como exemplo aqui,
pensamos não haver valor padrão e não pode estar vazio, obrigatório).

- Em seguida, existem os campos *Protocol*, *Timeout* e os *MaxConns* campos que não podem ser vazios, mas possuem valores padrão,
como: protocol is tcp, timeout 30 seconds e número máximo 1024 de links.

- *TLS* Há também um link seguro, que requer a configuração do certificado relevante e da chave privada. Isso pode estar vazio.

Portanto, para uma configuração como a acima, precisamos de uma variedade *Server*,
conforme mostrado abaixo:

````go
func NewDefaultServer(addr string, port int) (*Server, error) {
  return &Server{addr, port, "tcp", 30 * time.Second, 100, nil}, nil
}

func NewTLSServer(addr string, port int, tls *tls.Config) (*Server, error) {
  return &Server{addr, port, "tcp", 30 * time.Second, 100, tls}, nil
}

func NewServerWithTimeout(addr string, port int, timeout time.Duration) (*Server, error) {
  return &Server{addr, port, "tcp", timeout, 100, nil}, nil
}

func NewTLSServerWithMaxConnAndTimeout(addr string, port int, maxconns int, timeout time.Duration, tls *tls.Config) (*Server, error) {
  return &Server{addr, port, "tcp", 30 * time.Second, maxconns, tls}, nil
}
````

Como o Go não oferece suporte a funções sobrecarregadas, você precisa usar nomes de função diferentes para diferentes opções de configuração.

## Configure Object Schema

A maneira mais comum de resolver esse problema é usar um objeto de configuração como este:

````go
type Config struct {
    Protocol string
    Timeout  time.Duration
    Maxconns int
    TLS      *tls.Config
}
````

Nós movemos as opções não obrigatórias para uma estrutura, então o Serverobjeto se tornou:

````go
type Server struct {
    Addr string
    Port int
    Conf *Config
}
````

Portanto, precisamos apenas de uma função NewServer(), e precisamos construir o objeto Config antes de usá-lo.

````go
func NewServer(addr string, port int, conf *Config) (*Server, error) {
    //...
}

//Using the default configuratrion
srv1, _ := NewServer("localhost", 9000, nil) 

conf := ServerConfig{Protocol:"tcp", Timeout: 60*time.Duration}
srv2, _ := NewServer("locahost", 9000, &conf)
````

Geralmente este código não é ruim, podemos parar por aí. No entanto, para os aspirantes a programadores com limpeza,
eles podem ver que uma das coisas ruins é que o Config não é necessário, então precisa julgar se é *nil ou Vazio* - Config{}
o que faz com que nosso código fique verboso.

## Builder mode

Se você for um programador Java, aqueles familiarizados com padrões de projeto naturalmente usarão o padrão Builder.
Por exemplo o seguinte código:

````go
User user = new User.Builder()
  .name("Jorge Luis")
  .email("jlfonseca@gmail.com")
  .nickname("JLFonseca")
  .build();
````

Seguindo o padrão acima, podemos reescrever o código anterior no código a seguir 
(Observação: o código a seguir não considera o tratamento de erros. Para obter mais informações sobre o tratamento de erros, consulte "Go Programming Patterns: Error Handling "):

````go
//Use a constructor class to do the packaging
type ServerBuilder struct {
  Server
}

func (sb *ServerBuilder) Create(addr string, port int) *ServerBuilder {
  sb.Server.Addr = addr
  sb.Server.Port = port
  //Other code sets the default value of other members
  return sb
}

func (sb *ServerBuilder) WithProtocol(protocol string) *ServerBuilder {
  sb.Server.Protocol = protocol 
  return sb
}

func (sb *ServerBuilder) WithMaxConn( maxconn int) *ServerBuilder {
  sb.Server.MaxConns = maxconn
  return sb
}

func (sb *ServerBuilder) WithTimeOut( timeout time.Duration) *ServerBuilder {
  sb.Server.Timeout = timeout
  return sb
}

func (sb *ServerBuilder) WithTLS( tls *tls.Config) *ServerBuilder {
  sb.Server.TLS = tls
  return sb
}

func (sb *ServerBuilder) Build() (Server) {
  return  sb.Server
}
````

Então pode ser usado da seguinte forma

````go
sb := ServerBuilder{}
server, err := sb.Create("127.0.0.1", 8080).
  WithProtocol("udp").
  WithMaxConn(1024).
  WithTimeOut(30*time.Second).
  Build()
````

O método acima também é muito claro. Não há necessidade de uma classe Config adicional.
O método de chamada de função encadeada é usado para construir um objeto.
Só é necessário adicionar uma classe Builder adicional.
Esta classe Builder parece ser um pouco redundante Parece que podemos fazer isso diretamente na Serverconstrução O Construtor faz exatamente isso.
Mas pode ser um pouco problemático ao lidar com erros (precisar adicionar um membro de erro à estrutura do servidor, o que destrói a "pureza" da estrutura do servidor), é pior do que uma classe wrapper.

Se quisermos omitir essa estrutura wrapper, então é nossa vez de usar **Functional Options**, programação funcional.

## Functional Options

Primeiro, vamos definir um tipo de função:

````go
type Option func(*Server)
````

Então, podemos definir um conjunto de funções de forma funcional da seguinte forma:

````go
func Protocol(p string) Option {
    return func(s *Server) {
        s.Protocol = p
    }
}
func Timeout(timeout time.Duration) Option {
    return func(s *Server) {
        s.Timeout = timeout
    }
}
func MaxConns(maxconns int) Option {
    return func(s *Server) {
        s.MaxConns = maxconns
    }
}
func TLS(tls *tls.Config) Option {
    return func(s *Server) {
        s.TLS = tls
    }
}
````

O conjunto de código acima passa um parâmetro e retorna uma função que define o parâmetro seu próprio Server. Por exemplo:

- Quando chamamos uma dessas funções com MaxConns(30)
- Seu valor de retorno é uma função de func(s* Server) { s.MaxConns = 30 }.

Bem, agora vamos definir uma função NewServer() com um parâmetro variável options que pode passar várias funções acima e,
em seguida, usar um loop for para definir nosso objeto Server.

````go
func NewServer(addr string, port int, options ...func(*Server)) (*Server, error) {

  srv := Server{
    Addr:     addr,
    Port:     port,
    Protocol: "tcp",
    Timeout:  30 * time.Second,
    MaxConns: 1000,
    TLS:      nil,
  }
  for _, option := range options {
    option(&srv)
  }
  //...
  return &srv, nil
}
````

Então, quando criamos um objeto Server, podemos fazer isso.

````go
s1, _ := NewServer("localhost", 1024)
s2, _ := NewServer("localhost", 2048, Protocol("udp"))
s3, _ := NewServer("0.0.0.0", 8080, Timeout(300*time.Second), MaxConns(1000))
````

Portanto, no futuro, quando você quiser jogar com código semelhante, é altamente recomendável usar o método Functional Options,
que traz pelo menos os seguintes benefícios:

- Programação intuitiva
- Altamente configurável
- Fácil de manter e estender
- Auto documentação
- Fácil de começar para recém-chegados
- Nada confuso (é nulo ou vazio)

## Reference documentation
- [“Self referential functions and design” by Rob Pike](http://commandcenter.blogspot.com.au/2014/01/self-referential-functions-and-design.html)