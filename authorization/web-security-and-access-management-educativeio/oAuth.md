# OAuth

- OAuth is an authorization framework that allows a client app to retrieve information from another system using a 
  token which is valid for a limited time. The application users authorize the client app to retrieve information on 
  their behalf.
- OAuth should only be used for **authorization**.
- Why shouldn't OAuth be used for authentication?
    - During authentication, a client app will need some user information. The client app can send the token 
      to the resource server (Facebook in our example) to get the user information, but it is a bad idea because 
      there is no standard way to send the user information back to the client. If applications use OAuth for 
      authentication, then every implementation will be different, which will be a problem. This is the major 
      problem using OAuth for authentication.
    - OpenId Connect defines a standard way to return user information. It defines a UserInfo endpoint which 
      can be used to access user information.
- OAuth Terminologies:
    - Resource owner
    - Client
    - Resource server
    - Authorization server
    - Authorization grant
    - Authorization code
    - Access Token
    - Scope
- Steps of Authorization code grant type:
    - Authorization request
    - Authorization response
    - Token request
    - Token response
- **Why are there two steps to get the access token? Why do we need to first get the authorization code and then exchange
  it with the access token?**
  - There are two concepts:
    - Front Channel: Less secure browser/mobile app to the server channel.
    - Back Channel: Highly secure server to server communication channel.
  - It is not safe to share the client secret and get the access token on the front channel. Therefore, we first
    fetch the authorization code using the front channel and then request the access token using the back channel.
- **In Authorization Code grant flow the client app use the client_secret and authorization code to get the access code.**
