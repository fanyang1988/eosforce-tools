# EOSForce Tools

Some tools for eosforce write by golang.

## 1. Install

```bash
git clone https://github.com/fanyang1988/eosforce-tools.git
cd eosforce-tools
go mod vendor
go install ./...
```

## 2. Tools

### 2.1 force-abi2hex

force-abi2hex can transfer *.abi file for eosforce to packed hex string, this can used by update contracts by muit-sign.

Usage:

```bash
force-abi2hex -abi ./sss.abi > sss.hex
```
