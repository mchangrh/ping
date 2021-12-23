# mchangrh/ping

tiny webserver for ping testing

## flags
port          - port to listen on
ssl-key       - ssl key file
ssl-cert      - ssl cert file

## endpoints
```
/pixel.gif    - 36b gif for img timing
/echo/:text   - respond with text
/code/:code   - respond with http code
/version      - respond with version
/ping         - respond with "1"
/             - same as /ping
```