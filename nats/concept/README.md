# NATS Concepts

## Pull Consumer and Push Consumer

- A pull-based consumer is a type of consumer that explicitly requests messages from the server by using a fetch method,
  and it allows the client to have more control over the rate of messages it receives. In a pull-based system, the client 
  explicitly requests a specific number of messages, and the server responds with the requested messages. The client 
  acknowledges the receipt of each message, which removes it from the server's queue. Pull-based systems are often used 
  when the client wants to process messages at its own pace and has the capability to handle large batches of messages.
- A push-based consumer is a type of consumer where the server pushes messages to the client as soon as they arrive. 
  In a push-based system, the client subscribes to a topic, and the server pushes messages to the client as they become 
  available. The client has no control over the rate of messages it receives and must be able to handle messages as they
  arrive. Push-based systems are often used when the client needs real-time access to data and wants to receive messages
  as soon as they are available.
- In NATS, the default type of consumer is push-based, but it also supports pull-based consumers through JetStream. 
  JetStream is a message broker for NATS that supports both push and pull consumers, and it provides features such as 
  message persistence, stream retention policies, and more.

## Misc

- Synchronous request-reply can be useful in scenario where a response is needed before proceeding with further processing.
- ```Durable consumer``` means it will remember its place in the stream and continue consuming from where it left off if 
  it disconnects and reconnects later.
- ```AckPolicy: nats.AckExplicitPoilicy``` specifies that the consumer will manually ack each message it receives.
- 