2C_vehicle_ms

### Montar servicio a docker

Ejecutar las siguientes instucciones:

1. Subir la bd a rancher e inicializarla:

`docker-compose up`

2. Subir el programa a Rancher

`docker build -t vehicle-ms .`

3. Iniciar el programa

`docker run --name vehicle-ms -p 6004:6004 vehicle-ms`

**Aclaración:** Los comandos se deben ejecutar desde la terminal, en la ruta del proyecto (1. en una terminal y 2.,3. desde otra terminal).

### Ejecutar los microservicios

1. Ejecutar microservicio vehicle-db

`docker start vehicle-db`

2. Ejecutar microservicio vehicle-ms

`docker start vehicle-ms`
