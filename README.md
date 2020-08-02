# a plugin for drone used to connect to server with rcon protocol

  I write it in kotlin , but jdk images so fat , so I try to implement with golang. 

```shell
docker run -e RCON_HOST=127.0.0.1        \
           -e RCON_PORT=28016            \
           -e RCON_PASSWORD=shinobi9     \
           -e RCON_COMMANDS="say hello"   shinobi9/drone-rcon-go:<tag>
```