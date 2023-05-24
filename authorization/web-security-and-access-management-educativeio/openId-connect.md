# OpenID Connect

- OpenID Connect is an extension of OAuth. It is a thin layer above OAuth which adds support for authentication.
- Participants in OpenID Connect:
    - End User: End User refers to the entity for which the client is requesting identity information. This is called Resource Owner in OAuth.
    - Relying Party: This is the party that relies on the authorization server to provide the identity of the End User. This is called client in OAuth.
    - Identity Provider: This is a server that provides identity information about the End User. This is called Authorization Server in OAuth.
- Identity Token: Encodes the user's authentication information.
- The identity token contains only basic information about the user. To get the complete user information, the client
  must send the access token to UserInfo endpoint.
- There are four types of scopes defined in OpenID Connect. When a particular **scope** is requested, then all the
  attributes under that scope are returned.
    - email: Fetch the user's email info.
    - phone: Fetch the user's phone info.
    - profile: Fetch the user's default info.
    - address: Fetch the user's address.<br>
  There is one more scope value, i.e. ```openid```. This value is mandatory if the client app needs an identity token
  in the response.
- ```Claims``` are name/value pairs that contain information about a user.
- The authorization server defines some endpoints that are used by the client to request some data. The core endpoints are:
    - Authorization endpoint
    - Token endpoint
    - UserInfo endpoint
- **Can user information be fetched from the ```UserInfo``` endpoint by sending the access token in the request even if
  ```openid``` was not provided in the ```scope``` field when an access token was requested?**
    - NO.<br>
  Let's say that while sending a request to ```token endpoint```, the ```scope``` value is "openid email". The client sends
  this request and gets an access token. If the client sends this access token to the UserInfo endpoint, it will get only
  email information. It will not get an address or any other information.
- Implicit Code flow uses the authorization endpoint only. It does not use the token endpoint.
- To get the ```identity token``` the value of the scope field must contain openid.
- In the implicit flow, we get the access token and identity token from the ```authorization endpoint```. This is faster
  but is not secure.
- In the hybrid flow, the client gets immediate access to the identity token from the ```authorization endpoint``` itself.
  The client also gets the authorization code from the ```authorization endpoint```. Later it fetches the access token
  from the token endpoint which can be used to get further user info.