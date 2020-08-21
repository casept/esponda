# Test bot, please ignore

## Build

* With [Nix](https://nixos.org/download.html) (Linux/MacOS):
    * If you want a shell with all dependencies present
        ```shell
        nix-shell --command "go build && ./esponda -t <your-bot-token>"
        ```
    * If you want to run it on NixOS, use the provided service.

* On other platforms, figure out how to install and configure `go` and `imagemagick`.

## Usage

```shell
$ ./esponda --help
Usage of ./esponda:
  -t string
    	Bot Token
```