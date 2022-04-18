# Functional Options
![functional_options](functional_options.png)

Neste artigo, vamos discutir o padrão de programação Functional Options. Este é um caso de aplicação de programação funcional,
e as habilidades de programação também são ótimas, sendo atualmente o padrão de programação mais popular na linguagem Go.
No entanto, antes de podermos discutir formalmente esse padrão, precisamos primeiro dar uma olhada no tipo de problema que está sendo resolvido.

## content
1) [configuration options problem](#configuration-options-problem)
2) [Configure Object Schema](#configure-object-schema)
3) Builder mode
4) Functional Options
5) Reference documentation

## configuration options problem

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