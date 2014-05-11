package samgo

import(
    //"fmt"
    "reflect"
)

type Router struct{
    RoutMap map[string]reflect.Value
}

func(rout *Router)RegRouter(rout_conf map[string]interface{}){
    rout.RoutMap = make(map[string]reflect.Value)
    for controller_url,controller := range(rout_conf){
        handler := reflect.ValueOf(controller)//type : reflect.Value
        //handlerType := handler.Type()
        //fmt.Println(reflect.Indirect(handler).Type().Name())

        rout.RoutMap[controller_url] = handler
    }
}

func(rout *Router)FindRout(controller_url string)(controller reflect.Value,isExist bool){
    controller,isExist = rout.RoutMap[controller_url]
    return controller,isExist
}
