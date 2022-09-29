<p align="center">
  <a href="https://fullcycle.com.br/" target="blank"><img src="https://fullcycle.com.br/wp-content/themes/fullcycle/assets/images/fullcycle-logo.svg"/></a>
</p>

# PFA Go

## Descrição do Projeto

Esse projeto corresponde a um producer e um consumer de uma fila do rabbitmq utilizando golang.
O Prometheus fica acessando o rabbitmq a cada tanto segundos e pega as metricas do rabbit mq

## Passos para execução do desafio

1 - Passo > Subir os containers no docker

```bash
docker-compose up -d
```

2 - Passo > Acessar o Prometheus na porta 9090
Ir na aba status/targets e analisar que o prometheus esta vivisualindo o rabbitmqq

3 - Passo > Acessar o Grafana na porta 3000
Logar com o username admin e senha admin.
Apos isso eh necessario configurar o grafana pra acessar o prometheus e buscar as informacoes do rabbitmq



4 - Passo > Acessar o Rabbitmq na porta 15672
Logar com o username guest e senha guest e criar a fila de maneira manual
Ao mandar uma mensagem pro rabbitmq vc manda ela pra uma exchange e essa exchange manda a mensagem pra sua respectiva fila. Vc precisa conectar a fila na exchange


5 - Passo > Acessar o go_app e executar os producers

5 - Passo > Acessar o go_app e executar os producers


```bash
 docker run -it --rm luique97/codeeducation:latest
```

# Author

<table>
   <tr>
      <td align="center">
         <a href="http://github.com/lhfam97/">
            <img src="https://github.com/lhfam97.png" width="100px;" alt="Imagem de Luís Henrique Machado"/>
            <br />
            <sub>
               <b>Luís Henrique Machado</b>
            </sub>
          </a>
          <br />
          <a href="https://www.linkedin.com/in/luís-henrique-machado-98037a127/" title="Linkedin">@luishenriquemachado</a>
          <br />
          <a href="https://github.com/lhfam97/fastfeet-api/commits?author=lhfam97" title="Code">💻</a>
      </td>
   </tr>
</table>
