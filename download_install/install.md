# Download and install

<p align="center"><img src="assets/images/install-motorcycle.svg" width="auto" height="150"></p>

Neste tutorial iremos mostrar como instalar facilmente o Go.

Aqui você encontra outras forma de instalações.

- [Gerenciando as instalações Go](https://go.dev/doc/manage-install) -- Como instalar e desinstalar multiplas de versões.
- [Instalando a partir do código fonte](https://go.dev/doc/install/source) -- Como fazer o build e instalar na sua máquina a partir do código fonte.

## Go download

Clique no botão seguir para fazer download do instalador.

[Download Go for Windows](https://go.dev/dl/go1.18.1.windows-amd64.msi)
[Download Go for macOS](https://go.dev/dl/go1.18.1.darwin-amd64.pkg)
[Download Go for Linux](https://go.dev/dl/go1.18.1.linux-amd64.tar.gz)

Caso você esteja executando num ambiente diferente do Windows. Experimente [outros downloads](https://go.dev/dl/) 

> **Nota:** Por padrão, o comando go faz downloads e autentica usando Go module mirror e banco de dados Go checksum executados pelo Google. [Saiba mais](https://go.dev/dl)

## Go install

### Windows

1) Abra o arquivo MSI da pasta download e siga os passo a passo da instalação.
> Por padrão, o instalador instala o Go na pasta *Program Files or Program Files (x86)*. Caso necessário você pode alterá-lo.
Depois da instalação você precisa fechar e abrir o terminal para recarregae as variáveis de ambiente.

2) Valide a instalação do Go
   1) No **Windows**, clique no menu **Iniciar**
   2) Na caixa de pesquisa, digite *cmd*, e pressione a tecla **Enter**
   3) No Terminal de comando do windows, digite o seguinte comando
   ````shell
     $ go version
    ````
   4) O Comando imprime a versão do Go instalada na sua máquina

## [Vamos de código - Hello World](./get_started_with_go.md#nosso-hello-world)


