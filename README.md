# Micro Network

Micro Communication Protocol (MUCP)

## Overview

The Micro network is a protocol for service to service communication. 
It enables distributed systems to be built over platform boundaries using 
a header based transport called MUCP. The protocol behaves much like SMTP, 
allowing gateways to relay messages beyond a single domain to create a 
network of networks. 

## Motivation

The current state of internet protocols have not evolved in decades. We've
since relied heavily on HTTP for service based communication where it was 
initially intended simply for web pages. While frameworks and technologies 
have emerged to improve API based consumption of infrastructure services 
we have yet to define anything for end user consumption.

## Goals

A few expressive goals for the protocol

- simplicity by design and human readable
- extensible via federation not protocol changes
- consumable by clients and services alike

## Status

Initial MVP implementation of MUCP protocol 0.1 in development.

## Protocol

See the [PROTOCOL](PROTOCOL.md) doc for more info.

## License

Apache 2.0
