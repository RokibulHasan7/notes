# Networking Introduction

## OSI Model

- The OSI model is a conceptual framework for describing how two systems communicate over a network.
- It was meant to be a protocol suite to power networks but lost to TCP/IP.
- The OSI model description is a complex and exact way of saying networks have layers like cakes or onions.
- The OSI model breaks the responsibilities of the network into ```seven distinct layers```.
- The layers encapsulate information from the layer below it, these layers are Application, Presentation,
  Session, Transport, Network, Data Link and Physical.
  - ```Application```: Top layer of the OSI model and is the one the end user interacts with every day. The single
    biggest interface is HTTP. Other example of the application layer that we use daily are DNS, SSH and SMTP.
    Those application are responsible for displaying and arranging data requested and sent over the network.
  - ```Presentation```: It can be referred to as the **syntax layer**. Encryption is also done at this layer.
  - ```Session```: It builds, manages, and terminates the connections between the local and remote applications.
  - ```Transport```: The transport layer controls a given connection's reliability through flow control, segmentation
    and desegmentation, and error control.
  - ```Network```: The Network layer performs routing functions and might also perform fragmentation and reassembly
    while reporting delivery errors. **Routers operate at this layer**.
  - ```Data Link```: This layer is responsible for the host-to-host transfers on the same network. It defines the 
    protocols to create and terminate the connections between two devices.
  - ```Physical```: This layer converts data in the form of digital bits into electrical, radio or optical signals.
    The wire signaling protocols are also defined at this layer.
- Each layer takes data from the previous layer and encapsulates it to make its Protocol Data Unit(PDU).
  The PDU is used to describe the data at each layer. PDUs are also part of TCP/IP.


### OSI Model Summary

| Layer Number | Layer Name | Protocol Data Unit | Function Overview                                                                                                        |
|--------------|---------|------|--------------------------------------------------------------------------------------------------------------------------|
| 7            | Application | Data | High-level APIs and application protocols like HTTP, DNS, and SSH.                                                       |
| 6            | Presentation | Data | Character encoding, data compression, and encryption/ decryption.                                                        |
| 5            | Session | Data | Continuous data exchanges between nodes are managed here: how much data to send, when to send more.                      | 
| 4            | Transport | Segment, Datagram | Transmission of data segments between endpoints on a network, including segmentation, acknowledgement, and multiplexing. | 
| 3            | Network | Packet | Structuring and managing addressing, routing, and traffic control for all endpoints on the network.                      |
| 2            | Data Link | Frame | Transmissio of data frames between two nodes connected by a physical layer.                                              | 
| 1            | Physical | Bit  | Sending and receiving of bitstreams over the medium.                                                                     | 


## TCP/IP

- TCP/IP creates a heterogeneous network with open protocols that are independent of the operating system and
  architectural difference.
- Layers of TCP/IP:
  - Application: Represent data user to user, plus encoding and dialog control.
  - Transport: TCP and UDP are the primary protocols of the Transport layer that provide host-to-host communication
    services for applications. There are 65535 ports.
  - Internet: The internet, or Network layer, is responsible for transmitting data between networks. For an outgoing
    packet, it selects the next-hop host and transmits it to that host by passing it to the appropriate link-layer.
    A checksum ensures that the information in a received packet in accurate, but this layer does not validate data integrity.
  - Link: This layer includes protocols for moving packets between two Internet layer hosts. **Ethernet is the dominant
    protocol at this layer.**
  - Physical Layer: Controls the hardware devices and media that make up the network.

## TCP

- TCP is a connection-oriented, reliable protocol, and it provides flow control and multiplexing.
- TCP is considered connection-oriented because it manages the connection state through the life cycle of the connection.
- TCP is known as a host-to-host layer protocol.
- TCP uses a **three-way handshake**.
- TCP is a stateful protocol.
- ```tcpdump``` allows administrators and users to display all the packets processed on the system and filter them
  out based on many TCP segment header details.
- The Transport layer does not offer any security protection for data transiting the network. TLS adds additonal
  security on top of TCP.

## TLS

- TLS adds encryption to TCP.
- HTTP transactions can be completed without TLS but are not secure from eavesdroppers on the wire.
- TLS, much like TCP, uses a handshake to establish encryption capabilities and exchange keys for encryption.

## UDP

- UDP is an excellent choice for applications that can withstand packet loss such as voice and DNS.

## Others

- Internet relies on Border Gateway Protocol(BGP).
- The Link layer is responsible for connectivity to the local network.
- MAC addresses have two parts: the organization unit identifier (OUI) and the NIC-specific parts.
- In Kubernetes, we are mostly interested in IPv4 and ARP (Address Resolution Protocol) packets. IPv6 has recently been introduced to K8s in the
  1.19 release.
- All devices on the network keep a cache of ARP addresses for fast lookups for those known hosts, so it does not
  have to send an ARP request for every frame the host wants to send out.
- 