# Reverse Proxy
This is a simple example that shows how to build a configuration-driven [reverse proxy](https://en.wikipedia.org/wiki/Reverse_proxy) in Go. This was written for educational purpose, by no mean is meant to be used in production ;)

## configuration

The proxy by default uses the `config.json` file inside the main folder, you can define your own config file and pass the folder path using the `--config_path path/to/config.json` flag.

Here is the default `config.json`

```json
{
  "ssl": true,
  "default_url": "localhost",
  "default_port": 1333,
  "rules": [
    {
      "matcher": "/auth",
      "downstream_port": 1331
    },
    {
      "matcher": "/api",
      "downstream_port": 1332
    }
  ]
}
```


#### Description

| Field | Description|
|:---|:---|
| ssl | _[bool]_  if `true` the proxy uses the `HTTPS` protocol and listen on port `443`, otherwise it uses `HTTP` on port `8080`|
| default_url | _[string]_ default url for the downstream servers |
| default_port | _[int64]_ default port for the downstream servers |
| roles | _[[]object]_ Array of rules that describe the matcher and the downstream port and URL where forward the request to.<ul><li>__matcher__: _[string]_ describes the URL path that, when matched, the proxy uses this rules</li><li>__downstream_port__: _[int64]_ Port of the downstream server</li><li>__downstream_url__: _[string]_ URL of the downstream server</li></ul> |


## Running
We can test the `reverse-proxy` using the default configuration file `config.json`.

In a terminal console we can start the reverse-proxy using the following command.

```bash
# Start the reverse Proxy with default config file
> go build && ./reverse-proxy
2020/01/01 19:32:41 Server listening on... https://localhost:443

# Start the reverse Proxy with specific config file
go build && ./reverse-proxy --config_path config.json
2020/01/01 19:32:41 Server listening on... https://localhost:443
```

In order to test our default configuration we need to start 3 different downstream servers, for each of them we should open a new terminal console and rung the following commands:

```bash
# On a new terminal: server 1
> go build && ./server -p 1331
Server listening on port 1331

# On a new terminal: server 2
> go build && ./server -p 1332
Server listening on port 1332

# On a new terminal: server 3
> go build && ./server -p 1333
Server listening on port 1333
```



Here is how it works.

[![asciicast](https://asciinema.org/a/QhAoBtORwRz906q68t1BcAvKb.svg)](https://asciinema.org/a/QhAoBtORwRz906q68t1BcAvKb)
