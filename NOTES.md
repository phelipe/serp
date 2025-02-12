A tarefa consiste em pegar informações das páginas html do google e fornecere em um formato json
We need get information of google pages

To solve the problem first we need to analize the html code of the google. Basically the html has css class names random generated, wich means, the class name can change, so we need a way to get the information  without the necessity of scrap using class names.

Problabbly a better aproach is to use the role atribute, this atribute is used in screen readers so the value shoud have a meaningfull name.

# portugues

Para resolver esse problema primeiro temos que analisar o código html do google. Basicamente temos que todas as classes CSS do código estão com um nome randomizado, isso normalmente é feito por frameworks para dificultar nome repetido de classes. Sendo assim, usar o nome das classes não será uma boa abordagem, visto que uma simples mudança de uma cor por parte do google pode fazer com que todos os nomes das classes mudem novamente, assim sendo temos que buscar outra forma de obter os dados.

Páginas da web de grandes sites buscam seguir o mínimo de acessibilidade, apresentando assim recursos para possibilitar a leitura das informações por leitores de tela, então fazer o uso de propriedades do html associadas com a parte de acessibilidade podem ser um caminho. Verificando o código temos que existe o uso do atributo role, e é através deste que iniciaremos o nossos testes com o web scrapping (mais informações sobre o assunto podem ser obtidas no seguinte [link](https://www.w3.org/TR/wai-aria/#introroles) ).

Escolhida nossa abordagem inicial, vamos então para o próximo passo, iniciar o nosso código. Aqui farei em linguagem GO, por motivos, como já dito antes, Ruby ainda não é uma das linguagens que domino então vou para uma das que venho utilizando recentemente.

