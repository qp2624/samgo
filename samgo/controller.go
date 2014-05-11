package samgo

import(
    "fmt"
    "net/http"
)

type Controller struct{
    input  *http.Request
    out    http.ResponseWriter
    Params map[string]string
    Server map[string]string
}

func(this *Controller)Init(w http.ResponseWriter,r *http.Request){
    this.out    = w
    this.Params = make(map[string]string)
    this.Server = make(map[string]string)

    for k,v := range(r.Form){
        this.Params[k] = v[len(v) - 1]//同一参数多次赋值以最后一次为准
    }
    this.Server["REQUEST_METHOD"] = r.Method
}

func(this *Controller)Get(params string, default_param ...string)(res string){
    v,ok := this.Params[params]
    if(ok == false){
        if(len(default_param) > 0){
            v = default_param[0]
        }else{
            v = ""
        }
    }
    return v
}

func(this *Controller)Out(res string){
    fmt.Fprint(this.out,res)
}
