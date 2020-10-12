# go-kafka-labs
Studies for event driven architecture using go

# Setting up 

```sh
docker-compose up --force-recreate
```

# Creating demo topic

```sh
docker-compose exec kafka  kafka-topics --create --topic demo.orders.new --partitions 3 --replication-factor 3 --if-not-exists --zookeeper zookeeper:32181
```


# Consuming test 

```sh
docker-compose exec kafka kafka-console-consumer --bootstrap-server localhost:29092 --topic demo.orders.new --from-beginning --max-messages 100 
```
