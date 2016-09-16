Generate server and client certificates for setting up remote docker access.
----------

This is the same as `docker-machine regenerate-certs`, bar actually installing certificates on the remote machine.

Usage
----------

[Download](../../releases) the binaries suitable for your platform. Then `chmod +x` and

```
./generate-docker-certs.xxx.xxx -ip 1.2.3.4
```

This will create `./output` folder full of certificates.
