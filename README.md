<h1 align="center">
    Go Rabbit
</h1>

<p align="center">Service simulating RFID reader posting tags readed on a rabbitmq queue by MQTT topic, with the "backend" receiving, processing and publishing on another MQTT topic.</p>

&nbsp;

#### Run

* `sudo chmod +x env.sh && ./env.sh`
* `docker compose up -d --build`
* Access rabbitmq management [on this address](http://127.0.0.1:15672) with user `admin` and password `admin`
* Create a rabbitmq queue named `.tags.`.
* On the exchange `amq.topic` create a bind to queue `.tags.` and routing key: `.tags.reader`
* Execute application:

```bash
go run cmd/main/main.go
```

* Watch the topic `/tags/detected` with a mqtt client using mqtt version 3.1 to connecting
* Suggested mqtt client: [MQTTX](https://mqttx.app/downloads)
