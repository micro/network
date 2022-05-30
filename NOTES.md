# Notes

On architecture and services

## Architecture

- gateway - p2p network connectivity
- exchange - token swap mechanism
- vm - wasm virtual machine

# Services

Services to extend the network to work on top of Micro

- Identity - A way to create and manage identities for nodes
- [Trust](#trust-service) - Reputation based scoring of nodes and services
- Token - Generation and distribution of tokens
- Gateway - The gatekeeper of all requests from node to node

# Trust Service

The trust service is a reputation based scoring system

### Overview

Trust in a system is an inherently difficult thing to predict. Even as something is secure, running and performant 
it's unclear whether a service is actually producing useful data or correct results. Trust is something that's 
earned overtime. What we propose a reputation based scoring system that establishes the trust score of every 
service based on an aggregated number of factors [uptime, performance, response].
