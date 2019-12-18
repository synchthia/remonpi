RemonPi
=========================================

[![Go Report Card](https://goreportcard.com/badge/github.com/synchthia/remonpi)](https://goreportcard.com/report/github.com/synchthia/remonpi)
[![FOSSA Status](https://app.fossa.com/api/projects/git%2Bgithub.com%2Fsynchthia%2Fremonpi.svg?type=shield)](https://app.fossa.com/projects/git%2Bgithub.com%2Fsynchthia%2Fremonpi?ref=badge_shield)

Remote on RaspberryPi, RemonPi.  \
This is a simple Remote Management application....

![ScreenShot](https://raw.githubusercontent.com/synchthia/remonpi/master/docs/images/screenshot.png)

## Requirements
> Remon Pi required HexPi for Send IR Signal, etc.  \
> https://github.com/synchthia/hexpi

## How to build
```bash
# Build frontend & build
make clean frontend build
```

## Environment Variables
### HTTP_PORT
* Listening HTTP Port (default: 8080)

### REMONPI_VENDOR
* Vendor (Manufactor) name of Remote
    * Currently available only `mitsubishi`

### REMONPI_MODEL
* Remote Controller Model
    * kgsa3-c: KGSA3 (C Mode)

### REMONPI_DATABASE_PATH
* RemonPi LocalData storage (state saving)

### HEXPI_ADDRESS
* HexPi Endpoint Address (ex, `http://localhost:8081`)
