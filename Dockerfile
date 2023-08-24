FROM rabbitmq:3.12-management

RUN rabbitmq-plugins enable rabbitmq_mqtt

RUN chown -R rabbitmq:rabbitmq /var/lib/rabbitmq /etc/rabbitmq &&\
    chmod 777 /var/lib/rabbitmq /etc/rabbitmq
