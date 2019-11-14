RemonPi
=========================================

## Requirements
[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2Fsynchthia%2Fremonpi.svg?type=shield)](https://app.fossa.io/projects/git%2Bgithub.com%2Fsynchthia%2Fremonpi?ref=badge_shield)

> Remon Pi required HexPi for Send IR Signal, etc.
> https://github.com/synchthia/hexpi

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


## License
[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2Fsynchthia%2Fremonpi.svg?type=large)](https://app.fossa.io/projects/git%2Bgithub.com%2Fsynchthia%2Fremonpi?ref=badge_large)