RemonPi
=========================================

## Requirements
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
