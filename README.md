# rcon support for drone

  I write it in kotlin , but jdk images so fat , so I try to implement with golang. 
  Also , only tested in minecraft , any feature will be add later...

```shell
docker run -e PLUGIN_ADDRESS=127.0.0.1:28016        \
           -e PLUGIN_PASSWORD=shinobi9              \
           -e PLUGIN_TIMEOUT=30                     \
           -e PLUGIN_COMMANDS="say hello,say hi"   shinobi9/drone-rcon-go:<tag>
```