package tlog

import (
	"github.com/TwiN/go-color"
	"io"
	"log"
	"os"
)

type Loggers struct {
	Debug *log.Logger
	Log   *log.Logger
	Con   *log.Logger
	Err   *log.Logger
}

var Debug = log.New(io.Discard, "", log.Lmsgprefix+log.LstdFlags)
var Log = log.New(io.Discard, "", log.Lmsgprefix+log.LstdFlags)
var Con = log.New(io.Discard, "", log.Lmsgprefix+log.LstdFlags)
var Err = log.New(io.Discard, "", log.Lmsgprefix+log.LstdFlags)

var DefLoggers = Loggers{
	Debug: Debug,
	Log:   Log,
	Con:   Con,
	Err:   Err,
}

func InitDebug(prefix string, outs ...io.Writer) {
	Debug.SetPrefix(prefix)
	Debug.SetOutput(getOut(color.Blue, outs...))
}
func InitLog(prefix string, outs ...io.Writer) {
	Log.SetPrefix(prefix)
	Log.SetOutput(getOut(color.Gray, outs...))
}
func InitCon(prefix string, outs ...io.Writer) {
	Con.SetPrefix(prefix)
	Con.SetOutput(getOut(color.Green, outs...))
}
func InitErr(prefix string, outs ...io.Writer) {
	Err.SetPrefix(prefix)
	Err.SetOutput(getOut(color.Red, outs...))
}

func init() {
	InitErr("", os.Stdout)
	InitCon("", os.Stdout)
}
