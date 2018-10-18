# ubicaciones_crud
API de ubicaciones y luegares, Integración con CI
ubicaciones_crud master/develop
 ## Requirements
Go version >= 1.8.
 ## Preparation:
    Para usar el API, usar el comando:
        - go get github.com/udistrital/ubicaciones_crud
 ## Run
 Definir los valores de las siguientes variables de entorno:
  - `UBICACIONES_HTTP_PORT`: Puerto asignado para la ejecución del API
 - `UBICACIONES_CRUD__PGUSER`: Usuario de la base de datos
 - `UBICACIONES_CRUD__PGPASS`: Clave del usuario para la conexión a la base de datos  
 - `UBICACIONES_CRUD__PGURLS`: Host de conexión
 - `UBICACIONES_CRUD__PGDB`: Nombre de la base de datos
 - `UBICACIONES_CRUD__SCHEMA`: Esquema a utilizar en la base de datos
 ## Example:
UBICACIONES_HTTP_PORT=8095 UBICACIONES_CRUD__PGUSER=postgres UBICACIONES_CRUD__PGPASS=password UBICACIONES_CRUD__PGURLS=localhost UBICACIONES_CRUD__PGDB=local UBICACIONES_CRUD__SCHEMA=core_new bee run
 ## Model DB
![image](https://github.com/udistrital/ubicaciones_crud/blob/develop/modelo_ubicaciones_crud.png).
