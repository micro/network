# Micro Communication Protocol (MUCP)

Specification for the **micro communication protocol** (mucp). 

## Overview

The micro communication protocol is a simple design specification for how micro services sends and receives messages over any transport. 
Our intention is to define an open protocol for service to service communication on top of any system. This also provides us the ability 
to simplify the creation of clients and servers which accept micro based requests. 

Ideally client should be implemented using the micro proxy http interface to standardise creation.

## Contents

- [Protocol](#protocol)
- [Discovery](#discovery)

## Protocol

The protocol is a very simple transport agnostic set of headers and an encoded message body. The protocol supports request/response, 
bidirectional streaming and asynchronous message broadcasting. Where the transport or broker accepts headers (such as http) the message headers 
will be encoded in the transport headers. Otherwise the entire header and message will be encoded in an envelope. Our preference is to use protobuf 
but the protocol should scan for a starting json delimiter `{` to know whether to decode to json.

The protocol covers 3 forms of communication: 

- [Request](#request) - Sending a request and synchronously receiving a response
- [Stream](#stream) - Maintaining an open connection over which messages are passed back and forth
- [Publish](#publish) - Asynchronously broadcast events to topics with multiple interested parties

### Request

Request/Response communication allows a single request to be sent and a response to be received. The request and response 
are of identical format to correlate one to one mapping. A request should be passed with a unique id, name of the service, 
the endpoint being called and the content-type. 

An example request.

```
{
	Header: {
		"Micro-Id": "d02d5da0-14dc-11e9-ab14-d663bd873d93",
		"Micro-Service": "greeter",
		"Micro-Endpoint": "Say.Hello",
		"Content-Type": "application/protobuf",
	}
	Body: []byte(...)
}
```

In the event of an error we return it as a header. This may also be returned in the body.

```
{
	Header: {
		"Micro-Id": "d02d5da0-14dc-11e9-ab14-d663bd873d93",
		"Micro-Service": "greeter",
		"Micro-Endpoint": "Say.Hello",
		"Micro-Error": {"id":"greeter.Say.Hello","code":500,"detail":"Failed greeting","status":"Internal Server Error"},
	}
}
```

### Stream

A stream is a long live connection over which messages are passed back and forth. This could be request response or streaming updates 
such as gps location from a client to the server. A stream uses identical request/response semantics except it also includes a 
stream id.

```
{
        Header: {
                "Micro-Id": "d02d5da0-14dc-11e9-ab14-d663bd873d93",
		"Micro-Stream": "user.1"
                "Micro-Service": "geolocation",
                "Micro-Endpoint": "Gps.Update",
                "Content-Type": "application/protobuf",
        }
        Body: []byte(...)
}
```

### Publish

Messages can be broadcast asynchronously to a topic. This requires no knowledge of subscribers or interested parties a head of time. 
It provides a method for notification of events without requiring a response. In the event no subscribers exist, the messages 
can be saved in an inbox until subscribers are present to retrieve the messages at a later time.

An example broadcast message.

```
{
	Header: {
		"Micro-Id": "d02d5da0-14dc-11e9-ab14-d663bd873d93",
		"Micro-Topic": "events",
		"Content-Type": "application/protobuf",
	}
	Body: []byte(...)
}
```

In the event you want to subscribe to a topic you must specify a queue.

```
{
        Header: {
                "Micro-Id": "d02d5da0-14dc-11e9-ab14-d663bd873d93",
                "Micro-Topic": "events",
                "Micro-Queue": "customer",
        }
}
```

## Headers

Supported headers in the protocol (all prefixed with `Micro-`

- Service
- Method
- Endpoint
- Network
- Topic
- Event
- Queue
- Id
- Stream

### In Progress

We are looking at optional extra headers for routing

- `Micro-Protocol` to specify the protocol for the endpoint. 
- `Micro-Stream` to segregate streams on the same connection.
- `Micro-Channel` or `Micro-Queue` to specify a specific queue to segregate by.
- `Micro-Method` to indicate a request, response, stream or message
- `Micro-Event` defines the type of event sent to a topic

## Methods

Methods are the actions which can be taken. Much like HTTP GET, POST, PUT, PATCH, DELETE.

All methods are transported as the `Micro-Method` header.

### List of methods

- Connect - Connect to the network
- Close - Disconnect from the network
- Route - Advertise known routes
- Call - Make a request
- Stream - Create a stream
- Publish - Broadcast a message
- Subscribe - Subscribe to messages
- Keepalive - Heartbeat with peers
- Discover - Ask for a route to a service
- Send - Make a payment
- Receive - Receive a payment
- Listen - Start accepting messages
- Announce - Broadcast an announcement
- Ban - Ban a node or service
- Block - Drop messages
- Event - Observed events

## Discovery

Service and topic discovery is a non concern for the time being. Internally our implementation of micro supports multicast DNS, gossip and consul. We believe 
in closed networks users should have the choice to use whatever is a preference for their architectures. The [mucp](https://github.com/micro/network/tree/main/mucp) 
directory is the location for our default implementation. 

In regard to discovery across the public internet. We believe networks should be served via DNS SRV records. 
