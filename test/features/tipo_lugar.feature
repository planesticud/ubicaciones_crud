Feature: Validate API responses
  UBICACIONES_CRUD
  probe JSON reponses


Scenario Outline: To probe route code response  /tipo_lugar
  When I send "<method>" request to "<route>" where body is json "<bodyreq>"
  Then the response code should be "<codres>"      

  Examples: 
  |method|route         |bodyreq               |codres       |
  |GET   |/v1/tipo_lugar|./files/req/Vacio.json|200 OK       |
  |GET   |/v1/tipo_luga |./files/req/Vacio.json|404 Not Found|
  |POST  |/v1/tipo_luga |./files/req/Vacio.json|404 Not Found|
  |PUT   |/v1/tipo_luga |./files/req/Vacio.json|404 Not Found|
  |DELETE|/v1/tipo_luga |./files/req/Vacio.json|404 Not Found|


Scenario Outline: To probe response route /tipo_lugar       
  When I send "<method>" request to "<route>" where body is json "<bodyreq>"
  Then the response code should be "<codres>"      
  And the response should match json "<bodyres>"

  Examples: 
  |method|route         |bodyreq               |codres         |bodyres                |
  |GET   |/v1/tipo_lugar|./files/req/Vacio.json|200 OK         |./files/res4/Vok1.json |
  |POST  |/v1/tipo_lugar|./files/req/Yt1.json  |201 Created    |./files/res4/Vok2.json |
  |POST  |/v1/tipo_lugar|./files/req/Vacio.json|400 Bad Request|./files/res4/Ierr1.json|
  |POST  |/v1/tipo_lugar|./files/req/Nt1.json  |400 Bad Request|./files/res4/Ierr2.json|
  |POST  |/v1/tipo_lugar|./files/req/Nt2.json  |400 Bad Request|./files/res4/Ierr3.json|
  |POST  |/v1/tipo_lugar|./files/req/Nt3.json  |400 Bad Request|./files/res4/Ierr4.json|
  |POST  |/v1/tipo_lugar|./files/req/Nt4.json  |400 Bad Request|./files/res4/Ierr5.json|
  |POST  |/v1/tipo_lugar|./files/req/Nt5.json  |400 Bad Request|./files/res4/Ierr6.json| 
  |POST  |/v1/tipo_lugar|./files/req/Nt6.json  |400 Bad Request|./files/res4/Ierr7.json| 
  |POST  |/v1/tipo_lugar|./files/req/Yt2.json  |400 Bad Request|./files/res4/Ierr8.json| 
  |PUT   |/v1/tipo_lugar|./files/req/Yt2.json  |200 OK         |./files/res4/Vok2.json |
  |GETID |/v1/tipo_lugar|./files/req/Vacio.json|200 OK         |./files/res4/Vok2.json |
  |DELETE|/v1/tipo_lugar|./files/req/Vacio.json|200 OK         |./files/res4/Ino.json  |
  |DELETE|/v1/tipo_lugar|./files/req/Vacio.json|404 Not Found  |./files/res4/Ierr9.json|
