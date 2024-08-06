# Start the services

docker-compose up -d

# Check connection

docker-compose exec nats-client nats server check jetstream -s nats://sys:password@nats-server:4222

# Create a stream

docker-compose exec nats-client nats stream add --config /etc/nats/stream-config.json -s nats://sys:password@nats-server:4222

# Publish to a stream

docker-compose exec nats-client nats pub mystream.test 'Hello, NATS!' -s nats://sys:password@nats-server:4222

# Subscribe to a stream

docker-compose exec nats-client nats sub mystream.test -s nats://sys:password@nats-server:4222

# Stream info

docker-compose exec nats-client nats stream info mystream -s nats://sys:password@nats-server:4222

# Add kv pair

docker-compose exec nats-client nats kv add my-kv -s nats://sys:password@nats-server:4222

# Create a consumer

docker-compose exec nats-client nats consumer add mystream mystream-consumer --ack explicit -s nats://sys:password@nats-server:4222

# Read messages from the stream

docker-compose exec nats-client nats consumer next mystream mystream-consumer -s nats://sys:password@nats-server:4222
