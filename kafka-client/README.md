Docker Commands to run kafka : 

1. First create a docker network . 
```docker network create kafka-network```
2. Then run zookeeper container
```docker run -d --name zookeeper --network kafka-network -p 2181:2181 zookeeper:latest```
3. Run kafka container
```
docker run -d --name kafka --network kafka-network -p 9092:9092 \
    -e KAFKA_LISTENERS=PLAINTEXT://0.0.0.0:9092 \
    -e KAFKA_ADVERTISED_LISTENERS=PLAINTEXT://localhost:9092 \
    -e KAFKA_OFFSETS_TOPIC_REPLICATION_FACTOR=1 \
    -e KAFKA_ZOOKEEPER_CONNECT=zookeeper:2181 \
    wurstmeister/kafka:latest
```
4. Enter into the kafka container 
```docker exec -it kafka /bin/bash```
5. Create this topic named "test-topic"
```kafka-topics.sh --create --topic test-topic --partitions 1 --replication-factor 1 --bootstrap-server localhost:9092```
<!-- This command should run inside the kafka container -->

All THE STEPS ARE NOW COMPLETE.

If you want to do same task without using any go code . Then basically you have to enter into the kafka container in two different terminals . Then run these commands : 

```kafka-topics.sh --create --topic test-topic --partitions 1 --replication-factor 1 --bootstrap-server localhost:9092```

```kafka-console-producer.sh --topic test-topic --bootstrap-server localhost:9092```

For consumer in another terminal 

```kafka-console-consumer.sh --topic test-topic --from-beginning --bootstrap-server localhost:9092```