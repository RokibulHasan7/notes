# JSON Web Token

- We know that in session-based authentication, the session details are saved on the server. However, in a distributed
  system, it is not necessary that a request from a given user will always go to the same server. It’s quite possible
  that one request is handled by one particular server and the next request is handled by another server. In this case,
  we can't use session-based authentication as we can't save the session info on both servers.
- Storing and retrieving the session information from the database or memory is a costly process. Each time a new user
  authenticates, we need to store their information. And whenever a user sends a sessionId with the request then we
  need to validate it from the database or memory. This leads to a lot of back and forth.
- Token Based Authentication: Within the token payload, you can easily specify what resources a user can access. 
  For example, if a third-party API wants to access my Gmail account then I can provide a token that will allow that 
  API to collect only my contact information from Gmail. It will not be able to access other resources.
- If a client suspects that the token is stolen, then they can logout from the browser. On doing this, the token for 
  that user will be deleted from the browser storage and will be added to the blacklist present on the server. 
  When the hacker sends a request with the stolen JWT, the server will find it in the blacklist and throw an 
  unauthorized error.
- 