package main

import (
	"github.com/balwaninitu/nitu-s_bookstore_api/app"
)

/*microservice will work like : controllers -> service -> Providers(includes Db) -> Domain
application will start with controllers then it will call service abd service will going to
call external REST providers/ and databases/any other service which may have and from services also
calling domain and from domain respond back with given request.

when start to write code correct approach is to always start from domain then develop service and then
develop controllers and finally can put http framework like gin gonic etc, always start from outward which is
domain and then go inward which is controller
domain is core of application, service has business logic like users starting with active status and password are
protected this is all business logic

domain driven development: it is modular approach where your domain/core(entities) is inside the centre of your
application then have use-cases(services)(here there is business logic) then you have controlles, presenter and gateways
that can access use-cases to make them available to web, devices,db,UI and external interfaces(fig ref to medium
article bookmarked)

dependancy workflow: dependency from the device to the controller to the use-case to the entities i.e from outside
to inside. In our application gingonic accessing controller then use-case(services) and then accessing domains
but the domain nothing has to do with services and services dont know any controller and controllers dont know which
framework.

Data workflow:data flow from one external API eg. request accessing from REST API to controller and to domain and then response
from domain/entity process to service and return from controller to your REST response.const



*/

func main() {
	//once we execute it will go to start app map urls and run
	app.StartApplication()
}
