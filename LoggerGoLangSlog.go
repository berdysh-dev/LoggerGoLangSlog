package LoggerGoLangSlog

import (
    "fmt"
    "io"
    "net"
    "net/url"
    "encoding/json"
)

const (
    LOG_KERN        = 0x0
    LOG_USER        = 0x8
    LOG_MAIL        = 0x10
    LOG_DAEMON      = 0x18
    LOG_AUTH        = 0x20
    LOG_SYSLOG      = 0x28
    LOG_LPR         = 0x30
    LOG_NEWS        = 0x38
    LOG_UUCP        = 0x40
    LOG_CRON        = 0x48
    LOG_AUTHPRIV    = 0x50
    LOG_FTP         = 0x58
    LOG_LOCAL0      = 0x80
    LOG_LOCAL1      = 0x88
    LOG_LOCAL2      = 0x90
    LOG_LOCAL3      = 0x98
    LOG_LOCAL4      = 0xa0
    LOG_LOCAL5      = 0xa8
    LOG_LOCAL6      = 0xb0
    LOG_LOCAL7      = 0xb8
) 

const (
    LOG_EMERG       = 0
    LOG_ALERT       = 1
    LOG_CRIT        = 2
    LOG_ERR         = 3
    LOG_WARNING     = 4
    LOG_NOTICE      = 5
    LOG_INFO        = 6
    LOG_DEBUG       = 7

    INTERNAL_NOPRI  = 0x10
) 

const (
    ALERT           = LOG_ALERT
    CRIT            = LOG_CRIT
    DEBUG           = LOG_DEBUG
    EMER            = LOG_EMERG
    ERR             = LOG_ERR
    ERROR           = LOG_ERR
    INFO            = LOG_INFO
    NONE            = INTERNAL_NOPRI
    NOTICE          = LOG_NOTICE
    PANIC           = LOG_EMERG
    WARN            = LOG_WARNING
    WARNING         = LOG_WARNING
) 

const (
    LOG_PID         = 0x01
    LOG_CONS        = 0x02
    LOG_ODELAY      = 0x04
    LOG_NDELAY      = 0x08
    LOG_NOWAIT      = 0x10
    LOG_PERROR      = 0x20
) 

const (
    URL             = "URL"
    FACILITY        = "FACILITY"
) 

type t_my_writer struct {
    facility    int
    url         string

    conn        net.Conn 

    Scheme      string
    Opaque      string
    User        *url.Userinfo
    Host        string
    Path        string
    RawPath     string
    OmitHost    bool
    ForceQuery  bool
    RawQuery    string
    Fragment    string
    RawFragment string
}

func (this *t_my_writer) Setter(k string,v any){
    switch(k){
    case FACILITY:{
            this.facility = v.(int);
        }
    case URL:{
            this.url = v.(string);
        }
    }
}

func (this *t_my_writer) Open(){
    fmt.Printf("URL[%s]\n",this.url) ;

    ui, err := url.Parse(this.url) ;

    if(err != nil){
        fmt.Printf("err[%s]\n",err) ;
    }else{
        this.Scheme         = ui.Scheme ;
        this.Opaque         = ui.Opaque ;
        this.User           = ui.User ;
        this.Host           = ui.Host ;
        this.Path           = ui.Path ;
        this.RawPath        = ui.RawPath ;
        this.OmitHost       = ui.OmitHost ;
        this.ForceQuery     = ui.ForceQuery ;
        this.RawQuery       = ui.RawQuery ;
        this.Fragment       = ui.Fragment ;
        this.RawFragment    = ui.RawFragment ;
    }

    fmt.Printf("Scheme[%s]\n",this.Scheme) ;

    switch(this.Scheme){
    case "unix":{
            conn, err := net.Dial("unix",this.Path) ;
            if(err != nil){
                fmt.Printf("err[%s]\n",err) ;
            }else{
                fmt.Printf("ok[%s]\n",err) ;
                this.conn = conn ;
            }
        }
    }

}

func (this *t_my_writer) Write(p []byte) (n int, err error){
    var j map[string]interface{} ;
    e := json.Unmarshal(p,&j) ;
    if(e != nil){
        fmt.Printf("err[%s]\n",e) ;
    }else{
        fmt.Printf("level[%s]/msg[%s]/facility[0x%x]\n",j["level"],j["msg"],this.facility) ;
    }
    return len(p) , nil ;
}

func NewIoWriter() (io.Writer,*t_my_writer) {

    x := t_my_writer{} ;

    var my_writer io.Writer = &x ;

    return my_writer,&x ;
}









