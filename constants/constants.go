package constants

import (
	"os"
)

const UNDEFINED_COMMAND = "Undefined command/"
const ERROR_COMMAND = "Syntax error/"
const STOP_COMMAND = "Command stop/"
const UNDEFINED_WORD_KEY = "Undefined word key/"

var mode = "prod"
var Root = ""

func INIT_ROOT() {
	if mode == "dev" {
		Root = "./private"
	} else {
		os.Mkdir(os.TempDir()+`\holoproject`, 0755)
		Root = os.TempDir() + `\holoproject`
	}
}

const LOG_PATH = "/log.asc"
const COMMAND_PATH = "/command.asc"
const DATA_PATH = "/data.asc"

const PROJECT_NAME = "Holo"

var PROJECT_INIT = `
       )                  )    (             )                            
    ( /(  ( /(   )\ )  ( /(    )\ ) )\ )  ( /(                (        ) 
    )\( ) )\()) (()/(  )\())  (()/((()/(  )\())    (   (      )\   )  /( 
   ((_)\ ((_)\   /(_))((_)\    /(_))/(_))((_)\     )\  )\   (((_)  ( )(_))
    _((_)  ((_) (_))    ((_)  (_)) (_))    ((_)   ((_)((_)  )\___ (_(_()) 
   | || | / _ \ | |    / _ \  | _ \| _ \  / _ \  _ | || __|((/ __||_   _| 
   | __ || (_) || |__ | (_) | |  _/|   / | (_) || || || _|  | (__   | |   
   |_||_| \___/ |____| \___/  |_|  |_|_\  \___/  \__/ |___|  \___|  |_|   
`

var HelpMessage = `
 Holo Project

 -- cripto
 -- ecripto

 -- stop
 -- drop
 -- master

 -- g:key
 -- c:log

 -- run
 -- run:m
 -- place
 -- commands
 -- comm
 -- rm:c

`
