# go-kafka-labs
Studies for event driven architecture using go

# Setting up 

```sh
docker-compose up --force-recreate
```

# Creating demo topic

```sh
docker-compose exec kafka  kafka-topics --create --topic demo.orders.new --partitions 1 --replication-factor 1 --if-not-exists --zookeeper zookeeper:32181
```