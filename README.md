# ubicaciones_crud
API de ubicaciones y lugares

Integración con

 - `CI`
 - `AWS Lambda - S3`
 - `Drone 1.x`
 - `ubicaciones_crud master/develop`

## Requerimientos
Go version >= 1.8.

## Preparación
Para usar el API, usar el comando:

 - `go get github.com/planesticud/ubicaciones_crud`

## Ejecución
Definir los valores de las siguientes variables de entorno:

 - `UBICACIONES_HTTP_PORT`: Puerto asignado para la ejecución del API
 - `UBICACIONES_CRUD__PGUSER`: Usuario de la base de datos
 - `UBICACIONES_CRUD__PGPASS`: Clave del usuario para la conexión a la base de datos  
 - `UBICACIONES_CRUD__PGURLS`: Host de conexión
 - `UBICACIONES_CRUD__PGDB`: Nombre de la base de datos
 - `UBICACIONES_CRUD__SCHEMA`: Esquema a utilizar en la base de datos

## Ejemplo
UBICACIONES_HTTP_PORT=8085 UBICACIONES_CRUD__PGUSER=user UBICACIONES_CRUD__PGPASS=password UBICACIONES_CRUD__PGURLS=localhost UBICACIONES_CRUD__PGDB=bd UBICACIONES_CRUD__SCHEMA=schema_new bee run

## Modelo BD
![image](https://github.com/planesticud/ubicaciones_crud/blob/develop/modelo_ubicaciones_crud.png).
