# K8S Visitor Mode

![k8s_visitor_mode](k8s_visitor_mode.png)

Este artigo pretende discutir principalmente um padrão de programação usado nos comandos kubectl
– *Visitor* (Nota: Na verdade, kubectl são principalmente usados dois, um é o Builder e o outro é o Visitor).
Originalmente, o Visitor é um modelo de design muito importante no modelo de *design orientado a objetos* [consulte a entrada Wikipedia Visitor Pattern](https://en.wikipedia.org/wiki/Visitor_pattern),
esse padrão é uma maneira de separar o algoritmo da estrutura do objeto operacional.
O resultado prático dessa separação é a capacidade de adicionar novas operações a uma estrutura de objeto existente sem modificar a estrutura,
uma abordagem que segue o princípio *aberto/fechado*.

Neste artigo, vamos nos concentrar em como implementar esse padrão usando uma abordagem funcional.

## Content
1) [Simple example](#simple-example)
2) [k8s related background](#k8s-related-background)
3) [Implementation of kubectl](#implementation-of-kubectl)
   1) [Visitor pattern definition](#visitor-pattern-definition)
   2) [Name Visitor](#name-visitor)
   3) [Other Visitor](#other-visitor)
   4) [Log Visitor](#log-visitor)
   5) [consumer code](#consumer-code)
   6) [Visitor decorator](#visitor-decorator)

## Simple example
## k8s related background
## Implementation of kubectl
### Visitor pattern definition
### Name Visitor
### Other Visitor
### Log Visitor
### consumer code
### Visitor decorator