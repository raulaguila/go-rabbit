<h1 align="center">
    Go Rabbit
</h1>

<p align="center">Service simulating RFID reader posting tags readed on rabbitmq queue by MQTT topic, with the "backend" receiving, processing and publishing on other MQTT topic.</p>

<ul>
    <li>Generate the ".env" file running "env.sh" script</li>
    <li>Start docker-compose</li>
    <li>Access rabbitmq management <a href="http://127.0.0.1:15672">on this address</a> with user "admin" and password "admin"</li>
    <li>Create a rabbitmq queue named ".tags." without the quotes</li>
    <li>On the exchange "amq.topic" create a bind to queue ".tags." and routing key: ".tagsreader"</li>
    <li>Execute application:</li>
</ul>

```bash
go run cmd/main/main.go
```

<ul>
    <li>Watch the topic: "/tags/detected" on an mqtt client using mqtt version 3.1 when connecting. ex: <a href="https://mqttx.app/downloads">MQTTX</a></li>
</ul>
