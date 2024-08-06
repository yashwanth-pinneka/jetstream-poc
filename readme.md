docker-compose up -d

Check connection,
docker-compose exec nats-client nats server check jetstream -s nats://sys:password@nats-server:4222

Publish to a stream,
docker-compose exec nats-client nats stream add --config /etc/nats/stream-config.json -s nats://sys:password@nats-server:4222

Subscribe to a stream,
docker-compose exec nats-client nats sub mystream.test -s nats://sys:password@nats-server:4222

Stream info,
nats account info -s nats://sys:password@nats-server:4222

Add kv pair,
nats kv add my-kv -s nats://sys:password@nats-server:4222
