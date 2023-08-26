<h1 align="center">
    Go Rabbit
</h1>

<p align="center">Service simulating RFID reader posting tags readed on a rabbitmq queue by MQTT topic, with the "backend" receiving, processing and publishing on another MQTT topic.</p>

* Generate the `.env` file running `env.sh` script
* Start docker-compose
* Access rabbitmq management [on this address](http://127.0.0.1:15672) with user `admin` and password `admin`
* Create a rabbitmq queue named `.tags.` without the quotes
* On the exchange "amq.topic" create a bind to queue `.tags.` and routing key: `.tags.reader`
* Execute application:

```bash
go run cmd/main/main.go
```

* Watch the topic: `/tags/detected` with an mqtt client using mqtt version 3.1 when connecting. ex: [MQTTX](https://mqttx.app/downloads)
