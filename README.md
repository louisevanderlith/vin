# vin
Vehicle Database and VIN Decoder

## Run with Docker
* $ docker build -t avosa/vin:dev .
* $ docker rm VINDEV
* $ docker run -d -e RUNMODE=DEV -p 8095:8095 --network mango_net -v db/:/db/ --name VINDEV avosa/vin:dev
* $ docker logs VINDEV

# API
``GET v1/lookup/WAUZZZ8E88A025765``