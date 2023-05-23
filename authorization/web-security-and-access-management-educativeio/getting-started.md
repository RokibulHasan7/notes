# Getting Started

- Since XSS works by injecting malicious code into a website, the website owners should make sure that all user inputs
  are validated before they are stored into the database. The theory here is to treat all data or input 
  as malicious until they pass certain criteria, like **type and length requirements**.
- The ```SameSite``` flag in cookies is a relatively new method of preventing CSRF attacks and improving web application security. 
- If the session cookie is marked as a ```SameSite``` cookie, it is only sent along with requests that originate
  from the same domain. Therefore, even if the user clicks on the hyperlink provided by the attacker,
  the cookies will not be sent.
- Best practice to avoid a **CSRF attack**:
  - Try not to use multiple websites at the same time. If you are logged in into a website in one browser tab and using a
    malicious website in another tab, then the chances of CSRF attack increase.
  - Do not allow browsers to remember passwords.
- Types of Denial of Service Attacks:
  - Flood Attack
    - ICMP flood
    - A SYN flood
  - Crash Attack
- Distributed denial of service: In a distributed DoS attack, multiple systems orchestrate a synchronized DoS attack to 
  a single target. With this method, the target is attacked from many locations at once instead of being attacked from 
  a single location.
- Steps to stop DoS attack:
  - Black Hole Routing: The benefit is that it gives some time to application owners to look into the origin of the
    attack and take appropriate action.
  - Rate Limiting: This may lead to some valid requests being denied, the benefit of this method is that the system 
    will not be overwhelmed
