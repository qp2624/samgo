package samgo

import(
    "fmt"
    "net/http"
    "reflect"
    "strings"
    "os"
    "time"
)

type Server struct{
    Rout *Router
}

const(
    IMG_DIR = "/home/q/pic"
    IMG_URL = "/img"
)

func(server *Server)Start(){
    fmt.Println("HttpServer starting...")
    //http.ListenAndServe(":9090",server)
    s := &http.Server{
            Addr:           ":8080",
            Handler:        server,
            ReadTimeout:    10 * time.Second,
            WriteTimeout:   10 * time.Second,
            MaxHeaderBytes: 1 << 20,
    }
    s.ListenAndServe()
}

func(server *Server)ServeHTTP(w http.ResponseWriter,r *http.Request){
    r.ParseForm()//解析get、post数据


    if r.URL.Path == "/favicon.ico"{
        http.ServeFile(w,r,r.URL.Path)
        return 
    }

    //处理静态文件请求
    if strings.HasPrefix(r.URL.Path,IMG_URL) && r.URL.Path != IMG_URL && r.URL.Path != IMG_URL+"/"{
        file := IMG_DIR + r.URL.Path[len(IMG_URL):]
        f,err := os.Open(file)
        defer f.Close()

        if err != nil && os.IsNotExist(err){
            file = IMG_DIR + "/default.jpg"
        }
        http.ServeFile(w,r,file)
        return 
    }

    //parse request url
    path_args := strings.Split(r.URL.Path,"/")
    path_num := len(path_args)

    //default controller and action
    var controller string = "default"
    var action string = "Index"

    if path_num == 3{
        controller = strings.ToLower(path_args[1])
        action     = strings.Title(strings.ToLower(path_args[2]))
    }

    handler,isset := server.Rout.FindRout(controller)

    //controller not exist
    if isset == false{
        fmt.Fprint(w,"404")
        return
    }

    t := reflect.Indirect(handler).Type()//handler为一个类型为reflect.value的指针，t将获取这个指针指向资源的类型

    obj := reflect.New(t)//New生成一个新的t类型的value
    init_args := make([]reflect.Value,2)
    init_args[0] = reflect.ValueOf(w)
    init_args[1] = reflect.ValueOf(r)
    method_handler := obj.MethodByName(action)
    obj.MethodByName("Init").Call(init_args)

    //check if the request method exist
    ok := method_handler.IsValid()
    if ok == true{
        arr := make([]reflect.Value,0)
        //string->method
        method_handler.Call(arr)
        //fmt.Fprint(w,res[0]);
    }else{
        fmt.Fprint(w,"404")
    }
}

