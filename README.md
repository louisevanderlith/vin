# vin
Vehicle Database and VIN Decoder

## Run with Docker
* $ docker build -t avosa/vin:latest .
* $ docker rm VINDEV
* $ docker run -d -e RUNMODE=DEV -p 8095:8095 --network mango_net --name VINDEV avosa/vin:latest
* $ docker logs VINDEV