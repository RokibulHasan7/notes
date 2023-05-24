# HTTPS Basics

- Symmetric encryption is much faster than asymmetric encryption and is used to encrypt a large amount of data.
- The keys used in symmetric encryption are not very large, as the max length is ```256 bits```.
- In asymmetric encryption, the sender and receiver use a separate key to encrypt and decrypt the message. 
  This is also known as **PKI (Public Key Infrastructure)**.
- The keys used in asymmetric encryption are fairly large and can be around ```2048 bits```.
- ```SSL (Secure Sockets Layer)``` certificates create an encrypted environment between a client and a server. 
  A Secure Sockets Layer certificate (SSL certificate) is a small data file installed on a Web server that
  allows for a secure connection between the server and a web browser.
- SSL is a protocol that is used to secure the HTTP. SSL is deprecated now and Transport Layer Security (TLS)
  protocol is used instead. Most SSL certificates today also support the Transport Layer Security (TLS) protocol, 
  which is considered to be more secure than SSL.
- The application owner should install the SSL certificate on its web server. When an application is secured by an 
  SSL certificate then its URL starts with ```https``` instead of ```http```.
-  Cookies are limited to ```4kb``` in size.
- 