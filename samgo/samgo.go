package samgo
var Test string = "hello"
func RegRout(rout map[string]interface{})*Server{
    router := &Router{}
    router.RegRouter(rout)
    server := &Server{}
    server.Rout = router
    return server
}

