<p align="center">
  <a href="https://fullcycle.com.br/" target="blank"><img src="https://fullcycle.com.br/wp-content/themes/fullcycle/assets/images/fullcycle-logo.svg"/></a>
</p>

# PFA Go

## Descrição do Projeto

Esse projeto corresponde a um producer e um consumer de uma fila do rabbitmq utilizando golang.
O Prometheus fica acessando o rabbitmq a cada tanto segundos e pega as metricas do rabbit mq

## Passos para execução do desafio

<h3>1 - Passo > Subir os containers no docker</h3>

```bash
docker-compose up -d
```
</br>
<h3>2 - Passo > Acessar o Prometheus na porta 9090</h3>

- Config prometheus.yml
<img  src="./assets/prometheus-config.png"/>

- Ir na aba status/targets e analisar que o prometheus esta visualzando o rabbitmq

<img  src="./assets/prometheus-target.png"/>
<img  src="./assets/prometheus-analysis.png"/>

</br>
<h3>3 - Passo > Acessar o Grafana na porta 3000</h3>
Logar com o username admin e senha admin.
Apos isso eh necessario configurar o grafana pra acessar o prometheus e buscar as informacoes do rabbitmq

- Adicionar novo data source

<img  src="./assets/add_data_source.png"/>

- Selecionar Prometheus

<img alt="" src="./assets/prometheus.png"/>

- Adicionar o caminho do Prometheus

<img src="./assets/prometheus.png"/>

- Acessar o site do grafana: <a href="https://grafana.com">https://grafana.com</a>  e buscar o dashboard do rabbitmq e copiar o id

<img src="./assets/dashboard_rabbitmq.png"/>
<img src="./assets/dashboard_rabbitmq_pt2.png"/>
<img src="./assets/dashboard_rabbitmq_pt3.png"/>

- Importar esse Id na pagina do grafana no localhost:3000
<img src="./assets/import_dashboard.png"/>
<img src="./assets/import_dashboard_pt2.png"/>

</br>
<h3>4 - Passo > Acessar o Rabbitmq na porta 15672</h3>
- Logar com o username guest e senha guest e criar a fila de maneira manual

- Ir na aba queue e adicionar a fila orders
<img src="./assets/add_new_queue.png"/>

- Ir na aba exchanges na exchange na exchange amq.direct e fazer o bind com a queue orders

<img src="./assets/bind_exchange_queue_pt1.png"/>
<img src="./assets/bind_exchange_queue_pt2.png"/>
 

</br>
<h3>5 - Passo > Acessar o go_app e executar os producers</h3>

<img src="./assets/executar_bash_container_goapp.png"/>
<img src="./assets/comando_executando_producer.png"/>

```bash
   go run cmd/producer/main.go
```

- Analisar no grafana após execução do producer:

<img src="./assets/rabbitmq_grafana_after_producer,png"/>


</br>
<h3>6 - Passo > Acessar o go_app e executar os producers</h3>

<img src="./assets/executar_bash_container_goapp.png"/>

```bash
   go run cmd/main.go
```

- Imagem do grafana após funcionamento do consumer

<img src="./assets/analisando_apos_executar_consumers.png"/>


<h3>7 - Passo > Acessar MySql e ver se as orders estão sendo cadastradas</h3>

<img src="./assets/connection_postgres.png"/>

<img src="./assets/database_mysqql.png"/>

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
