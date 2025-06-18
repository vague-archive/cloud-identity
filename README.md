# Void Identity Service

**[WIP]** Experimenting with an OIDC identity service using the open source
[Zitadel OIDC](https://github.com/zitadel/oidc?tab=readme-ov-file) provider

See patchnotes - https://www.notion.so/voidinc/Patch-Notes-5-16-24-6796b3fbd39f4b719066eddac8f61c97?pvs=4#10a17db2b54c40baa8beabd70be571ae
See discord thread - https://discord.com/channels/1174701754739347487/1237543021642387497

```bash
  > make download  # install dependencies
  > make test      # run tests
  > make run       # run the service locally (default localhost:3000)
  > make build     # build bin/server release executable
  > make clean     # remove bin/* release executables
```

## Manual
 
```bash
> bin/server --help

NAME:
   void-identity - The Void Identity Service.

USAGE:
   server --env ENV --host HOST --port PORT

GLOBAL OPTIONS:
   --env value, -e value     environment (dev | test | stage | prod) (default: "dev") [$ENV]
   --host value              server host (default: "localhost") [$HOST]
   --port value, -p value    server port (default: 3000) [$PORT]
   --message value, -m value provide a custom message (default: "Hello World") [$MESSAGE]
   --help, -h                show help
   --version, -v             print the version
```
