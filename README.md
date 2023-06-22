# go-rabbitmq

Note: The project about x-delay is in watermill-direct folder <br><br>

To start rabbitmq: `brew services start rabbitmq`# go-rabbitmq <br>
Rabbitmq Server(Local): `http://localhost:15672/`
<br><br>
For docker run: `docker compose up` <br>
If you want to access terminal run `docker-compose exec rabbitmq bash`

<br><br>
Map:  my-exchange --> my-queue-tmp --> my-exchange-old --> my-queue <br>
my-exchange => new config for config delay and pub <br>
my-queue-tmp => just tmp <br>
my-exchange-old => new config for consuming and old config for publishing <br>
my-queue => old config for consuming <br>