package LoggerGoLangSlog

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
    alert           = LOG_ALERT
    crit            = LOG_CRIT
    debug           = LOG_DEBUG
    emer            = LOG_EMERG
    err             = LOG_ERR
    error           = LOG_ERR
    info            = LOG_INFO
    none            = INTERNAL_NOPRI
    notice          = LOG_NOTICE
    panic           = LOG_EMERG
    warn            = LOG_WARNING
    warning         = LOG_WARNING
) 

const (
    LOG_PID         = 0x01
    LOG_CONS        = 0x02
    LOG_ODELAY      = 0x04
    LOG_NDELAY      = 0x08
    LOG_NOWAIT      = 0x10
    LOG_PERROR      = 0x20
) 











