# Micro Network

An open network for Micro services.

## Overview

The Micro network is a protocol for communication between services across Micro deployments. 
It enables distributed and federated clusters to be built over platform boundaries using 
a header based transport called MUCP. The protocol behaves much like SMTP, allowing gateways 
to relay messages beyond a single domain to create a network of networks. Using this for 
services should enable cross cloud and cross org collaboration.

## Status

Currently in the early design phase.

## Protocol

See the [PROTOCOL](PROTOCOL.md) doc for more info.

## Implementation

See the [mucp](mucp) directory for the Go implementation.

## License

Apache 2.0
