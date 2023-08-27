package main  

var routes=[]string{
	"/",
	"/login",
	"/logout",
	"/register",
	"/activate",
	"/plans",
	"/subscribe",
}

func Test_Routes_Exist(t *testing.T){
	testRoutes:=testApp.routes()
	chiRoutes:=testRoutes.(chi.Router)
	for _,route:=range routes {
		routeExists(t,chiRoutes,route)
	}
}
func routeExists(t *testing.T,routes chi.Router,route string){
	found:=false 
	_=chi.Walk(routes,func(method string,foundRoute string,handler http.Handler,middleware ...func(http.Handler) http.Handler)error{
if route==foundRoute{
	found=true
}
return nil
	})
	if !found{
		t.Errorf("did not find %s in registered routes,",route)
	}
		
	}
