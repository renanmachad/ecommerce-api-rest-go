# Descrição do projeto
Este projeto é um exemplo de como implementar uma arquitetura limpa em Golang. A arquitetura limpa é uma abordagem de design de software que separa o código em camadas, cada uma com uma responsabilidade específica. Isso torna o código mais fácil de entender, manter e testar.

Este projeto usa as seguintes camadas:

# Entidades: 
Representam objetos do mundo real, como produtos ou clientes.
# Repositórios: 
Fornecem acesso aos dados, como um banco de dados ou um arquivo.
# Usecases: 
Implementam a lógica de negócios, como criar ou atualizar um produto.
Controladores: Recebem solicitações do usuário e as encaminham para os usecases.
Este projeto também inclui uma camada de utilitários que fornece funções comuns, como codificação e decodificação JSON.

# Objetivos do projeto
Este projeto tem os seguintes objetivos:

Demonstrar como implementar uma arquitetura limpa em Golang.
Fornecer um exemplo de como usar os recursos do Golang.
Servir como um guia para aprendizado da linguagem.
Como usar o projeto
Para usar o projeto, siga estas etapas:

Clone o repositório do GitHub.
Instale as dependências com o comando go get.
Execute o projeto com o comando go run main.go.
O projeto irá iniciar um servidor HTTP na porta 8080. Você pode enviar solicitações ao servidor usando um navegador da web ou uma ferramenta de teste, como o Postman.


## TODO`S:
* Implementar testes unitários para todos os componentes.
* Implementar todas operações para todas entidades.
* Melhorar respostas JSON

Responsáveis:

* [renanmachad](https://github.com/renanmachad/)



# Exemplos de solicitações
Aqui estão alguns exemplos de solicitações que você pode enviar ao servidor:

Criar  um pedido:

POST /api/v1/order

Body:

{

"customer_id":1,

"total_amount":23,

"order_status":"PURCHASED"

}


## Contribuições

Contribuições são bem-vindas. Para contribuir, siga estas etapas:

1. Faça um fork do repositório do GitHub.
2. Faça suas alterações no fork.
3. Envie uma solicitação de pull para o repositório original.

## Licença

Este projeto é licenciado sob a licença MIT.
Você pode personalizar essa descrição para atender às suas necessidades específicas. Por exemplo, você pode adicionar mais informações sobre as tecnologias ou recursos usados no projeto. Você também pode adicionar instruções sobre como configurar o projeto ou como usá-lo.

Aqui estão algumas dicas para escrever um README eficaz:

Seja conciso e direto ao ponto. Os usuários devem ser capazes de entender rapidamente o que o projeto faz.
Use uma linguagem clara e concisa. Evite usar jargões ou terminologia técnica que possa ser confusa para os usuários.
Forneça links para documentação ou recursos adicionais. Isso ajudará os usuários a aprender mais sobre o projeto.
Espero que isso ajude!