# Notes

On architecture and services

## Architecture

Core components that will be required to operate

- control - control plane to manage network
- gateway - p2p network connectivity
- exchange - token swap mechanism
- vm - wasm virtual machine

## Services

Services to extend the network to work on top of Micro

- Identity - A way to create and manage identities for nodes
- Trust - Reputation based scoring of nodes and services
- Token - Generation and distribution of tokens
- Gateway - The gatekeeper of all requests from node to node

## Trust Service

The trust service is a reputation based scoring system. Trust in a system is an inherently difficult thing to predict. 
Even as something is secure, running and performant it's unclear whether a service is actually producing useful data 
or correct results. Trust is something that's earned overtime. What we propose a reputation based scoring system that 
establishes the trust score of every service based on an aggregated number of factors [uptime, performance, response].

## Control

The control plane (ctrl) will manage the network, rules and any high level routing decisions

Make use of [https://github.com/juanfont/headscale](https://github.com/juanfont/headscale)

## Gateway

The gateway will establish and maintain p2p connectivity on the network and act as an outbound 
routing plane for local services.

It will potentially make use of [https://github.com/tailscale/tailscale/blob/main/tsnet](https://github.com/tailscale/tailscale/blob/main/tsnet)
