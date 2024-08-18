# Decentralized Chat Application

## Disclaimer
This project is a work in progress and is not yet functional. The goal is to build a decentralized chat application using Go and Libp2p, with a strong focus on privacy, security, and scalability. The project is designed with Clean Architecture and Domain-Driven Design (DDD) principles to ensure modularity and testability.

## Table of Contents
- [Introduction](#introduction)
- [Features](#features)
- [Technology Stack](#technology-stack)
- [Clean Architecture](#clean-architecture)
- [Domain-Driven Design (DDD)](#domain-driven-design-ddd)
- [Flow Diagram](#flow-diagram)
- [Getting Started](#getting-started)
    - [Prerequisites](#prerequisites)
    - [Installation](#installation)
    - [Running the Application](#running-the-application)
- [Contributing](#contributing)
- [License](#license)

## Introduction

This project is a fully decentralized chat application built using Go and Libp2p. Each instance of the application acts as a node in the network, allowing for peer-to-peer communication without a central server. The application is designed with a strong focus on privacy, security, and scalability, utilizing modern software architecture principles.

## Features

- **Decentralized**: No central server; each user is a node in the network.
- **Peer-to-Peer Communication**: Leverages libp2p for direct communication between nodes.
- **Encrypted Messaging**: Ensures privacy through end-to-end encryption.
- **DHT for Peer Discovery**: Utilizes a Distributed Hash Table (DHT) for finding peers.
- **Cross-Platform**: Compatible with desktop and mobile environments.
- **Extensible Architecture**: Built using Clean Architecture and Domain-Driven Design (DDD) principles.

## Technology Stack

You can find more about the tech stack in my mermaid diagram.

```mermaid
classDiagram
    class Go
    class Dart
    class Rust

    class LibP2P
    class WebRTC
    class Flutter
    class Tauri

    class BoltDB

    class Encryption {
        End-to-End Encryption
    }

    class CentralNode {
        Optional Central Node
    }

    Go --> LibP2P
    Dart --> Flutter
    Rust --> Tauri
    LibP2P --> BoltDB
    WebRTC --> BoltDB
    Go --> Encryption
    Dart --> Encryption
    Rust --> Encryption
``` 

## Clean Architecture

The project follows the principles of Clean Architecture, ensuring that the business logic is separated from the implementation details, making the codebase more modular and testable.

```mermaid
classDiagram
    direction TB

%% Top-Level Directories
    class cmd {
        +main.go
    }

    class internal {
        +domain/
        +usecase/
        +interface/
        +infra/
    }

    class pkg {
        +utility packages (e.g., logger, config)
    }

%% Domain Layer
    class domain {
        +chat/
        +user/
    }

    class chat {
        +entity.go
        +service.go
        +repository.go
    }

    class user {
        +entity.go
        +service.go
        +repository.go
    }

%% Use Case Layer
    class usecase {
        +chat/
        +user/
    }

    class chatUseCase {
        +sendMessage.go
        +createChannel.go
        +startVoiceCall.go
    }

    class userUseCase {
        +registerUser.go
        +authenticateUser.go
    }

%% Interface Adapters Layer
    class interface {
        +http/
        +websocket/
        +cli/
    }

    class http {
        +handler.go
        +router.go
    }

    class websocket {
        +handler.go
        +connection.go
    }

    class cli {
        +command.go
    }

%% Infrastructure Layer
    class infra {
        +persistence/
        +network/
        +security/
    }

    class persistence {
        +boltDB.go
    }

    class network {
        +libp2p.go
        +webrtc.go
    }

    class security {
        +encryption.go
        +keypair.go
    }

%% Relationships
    cmd --> internal : Uses
    internal --> domain : Contains Domain Logic
    internal --> usecase : Contains Application Logic
    internal --> interface : Adapts Interfaces
    internal --> infra : Implements Infrastructure

    domain --> chat : Defines Entities, Repositories, Services
    domain --> user : Defines Entities, Repositories, Services

    usecase --> chatUseCase : Implements Chat Use Cases
    usecase --> userUseCase : Implements User Use Cases

    interface --> http : Implements HTTP Handlers
    interface --> websocket : Implements WebSocket Handlers
    interface --> cli : Implements CLI Commands

    infra --> persistence : Implements Persistence (e.g., BoltDB)
    infra --> network : Implements Networking (e.g., LibP2P, WebRTC)
    infra --> security : Implements Security (e.g., Encryption, KeyPair)

    chatUseCase --> chat : Uses Chat Domain Logic
    userUseCase --> user : Uses User Domain Logic

    http --> chatUseCase : Calls Chat Use Cases
    websocket --> chatUseCase : Calls Chat Use Cases
    cli --> chatUseCase : Calls Chat Use Cases

    http --> userUseCase : Calls User Use Cases
    websocket --> userUseCase : Calls User Use Cases
    cli --> userUseCase : Calls User Use Cases

    persistence --> chat : Implements Chat Repository
    persistence --> user : Implements User Repository
    network --> chat : Implements Chat Networking
    security --> chat : Implements Chat Encryption

%% Additional Notes
    class Notes {
        +cmd: contains the entry point of the application, such as main.go, which sets up the server and orchestrates the initialization of other layers.
        +domain: domain contains the core business logic. Each subdirectory (e.g., chat, user) represents a different domain or subdomain.
        +usecase: usecase contains application-specific business rules, organized by feature or module (e.g., chat, user). Each use case corresponds to a specific application service or operation.
        +interface: interface contains the code that adapts the input/output for different interfaces, such as HTTP, WebSocket, and CLI.
        +infra: infra contains the implementations of infrastructure concerns like databases, networking, and security.
    }
```

## Domain-Driven Design (DDD)

The application is structured around DDD, focusing on core business concepts and ensuring that the domain logic is at the heart of the application.

```mermaid
classDiagram
    class User {
        +id: string
        +displayName: string
        +keyPair: KeyPair
    }

    class Message {
        +id: string
        +sender: User
        +receiver: User
        +content: string
        +timestamp: int64
    }

    class Channel {
        +id: string
        +name: string
        +isPrivate: bool
        +members: List<User>
        +messages: List<Message>
    }

    class KeyPair {
        +publicKey: string
        +privateKey: string
    }

    class ChannelPermissions {
        +canSendMessage: bool
        +canInviteUsers: bool
    }

    class Chat {
        +users: List<User>
        +channels: List<Channel>
    }

    class CommunicationSession {
        +participants: List<User>
        +isActive: bool
    }

    class MessagingService {
        +sendMessage(sender, receiver, content): Message
        +broadcastMessage(sender, channel, content): Message
    }

    class ChannelService {
        +createChannel(name, isPrivate): Channel
        +joinChannel(user, channel): bool
        +leaveChannel(user, channel): bool
    }

    class MediaService {
        +startVoiceCall(participants): CommunicationSession
        +startVideoCall(participants): CommunicationSession
    }

    User --> KeyPair
    Channel --> ChannelPermissions
    Chat --> Channel
    CommunicationSession --> User
    MessagingService --> Message
    ChannelService --> Channel
    MediaService --> CommunicationSession
```

## Flow Diagram
This diagram outlines the flow of data and operations within the application, from user input to peer communication.

```mermaid
sequenceDiagram
    participant User
    participant App as App (Mobile/Desktop)
    participant P2P as P2P Network (LibP2P)
    participant WebRTC
    participant DB as BoltDB
    participant Enc as Encryption

    User ->> App: Launch App
    App ->> Enc: Generate KeyPair
    Enc -->> App: Return KeyPair

    User ->> App: Register with Display Name
    App ->> P2P: Connect to P2P Network

    User ->> App: Initiate Chat with Peer
    App ->> P2P: Find Peer
    P2P -->> App: Peer Found

    User ->> App: Send Message
    App ->> Enc: Encrypt Message
    Enc -->> App: Encrypted Message
    App ->> P2P: Send Encrypted Message
    P2P -->> App: Message Delivered

    User ->> App: Start Voice/Video Call
    App ->> WebRTC: Initialize WebRTC Session
    WebRTC ->> Enc: Secure Media Stream
    WebRTC -->> App: Media Stream Established
    App ->> P2P: Notify Peers of Call

    App ->> DB: Store Chat History
```

## Getting Started

### Prerequisites
- Go: Install Go (version 1.22.5 or higher recommended)
- Git: Install Git
- Optional (for now): Docker for containerized deployment

## Installation
1. Clone the Repository:
```bash
git clone https://github.com/rafagomes/decentralized-chat-backend.git
```
2. Install Dependencies:
```bash
go mod tidy
```

## Running the Application
To start the application, you can run the following command:

```bash
go run cmd/p2pchat/main.go -port 8080
```

To start multiple nodes, simply open new terminals and specify different ports:
```bash
go run cmd/p2pchat/main.go -port 8081
```

## Contributing
Contributions are welcome! Please open an issue or submit a pull request for any improvements or bug fixes.

### Steps to Contribute
- Steps to Contribute:
- Fork the repository.
- Create a new branch: git checkout -b feature/my-new-feature.
- Commit your changes: git commit -am 'Add some feature'.
- Push to the branch: git push origin feature/my-new-feature.
- Submit a pull request.

## License
This project is licensed under the MIT License - see the LICENSE file for details.

### Summary:
- **Introduction**: Provides an overview of the project.
- **Features**: Lists the key features of the application.
- **Technology Stack, Clean Architecture, DDD, and Flow Diagrams**: Visually explains the project structure and concepts using Mermaid diagrams.
- **Getting Started**: Simple instructions to install and run the application.
- **Contributing**: Encourages contributions with clear steps.
- **License**: Clarifies the licensing under MIT.