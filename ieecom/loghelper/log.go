package loghelper
import (
    "fmt"
    "ieeblog/ieecom/config"
    "os"
    "time"
    "log"
)
/*
1.实现了日志系统的日志分级，默认的级别是Trace，
  用户通过SetLevel可以设置不同的分级
2.
*/
// Log levels to control the logging output.
const ( 
    LogFileFormat="20060102" 
    NewLine="\r\n"   //换行符
    LevelTrace = iota
    LevelDebug
    LevelInfo
    LevelWarning
    LevelError
    LevelCritical
)


// logLevel controls the global log level used by the logger.
var level = LevelTrace

// LogLevel returns the global log level and can be used in
// own implementations of the logger interface.
func LogLevel() int {
    return level
}


// SetLogLevel sets the global log level used by the simple logger.
func SetLogLevel(l int) {
    level = l
}
//StdOutLogger references the used application logger.
var StdOutLogger = log.New(os.Stdout, "", log.Ldate|log.Ltime)

// SetLogger sets a new logger.
func SetLogger(l *log.Logger) {
    StdOutLogger = l
}

//FileLogger references the used application logger.
var FileLogger *log.Logger

//InitFileLogger set fileLogger
func InitFileLogger(){
    filePath:=config.Conf.LogPath+time.Now().Format(LogFileFormat)+".txt"  //Log文件配置取
    //创建文件
    file,err:=os.OpenFile(filePath,os.O_RDWR|os.O_CREATE|os.O_APPEND,0766)
    if err!=nil{
        fmt.Printf("打开日志文件错误=%v \n",err)
        return
    }
    FileLogger=log.New(file,"",log.LstdFlags|log.Lshortfile|log.LUTC)
    fmt.Printf("创建日志文件完成")
    SetLogger(FileLogger)
    return
}

// Trace logs a message at trace level.
func Trace(v ...interface{}) {
    if level <= LevelTrace {
        StdOutLogger.Printf("[T] %v\n", v)
    }
}

// Debug logs a message at debug level.
func Debug(v ...interface{}) {
    if level <= LevelDebug {
        StdOutLogger.Printf("[D] %v\n", v)
    }
}

// Info logs a message at info level.
func Info(v ...interface{}) {
    if level <= LevelInfo {
        StdOutLogger.Printf("[I] %v\n", v)
    }
}

// Warning logs a message at warning level.
func Warning(v ...interface{}) {
    if level <= LevelWarning {
        StdOutLogger.Printf("[W] %v\n", v)
    }
}

// Error logs a message at error level.
func Error(v ...interface{}) {
    if level <= LevelError {
        StdOutLogger.Printf("[E] %v\n", v)
    }
}

// Critical logs a message at critical level.
func Critical(v ...interface{}) {
    if level <= LevelCritical {
        StdOutLogger.Printf("[C] %v\n", v)
    }
}

/*
//os 包

1: os.Stat(name string) (fi FileInfo, err error)       
//返回描述文件的FileInfo信息。如果出错，将是 *PathError类型。

2: os.IsExist(err error) bool 
//返回一个布尔值，它指明err错误是否报告了一个文件或者目录已经存在。它被ErrExist和其它系统调用满足

3: os.MkdirAll(path string, perm FileMode) error
//创建一个新目录，该目录是利用路径（包括绝对路径和相对路径）
  进行创建的，如果需要创建对应的父目录也一起进行创建
  如果已经有了该目录，则不进行新的创建
  当创建一个已经存在的目录时，不会报错.

4: os.Chmod(name string, mode FileMode) error
//更改文件的权限（读写执行，分为三类：all-group-owner）

5: os.OpenFile(name string, flag int, perm FileMode) (file *File, err error)
//指定文件权限和打开方式打开name文件或者create文件，其中flag标志如下:
O_RDONLY：只读模式(read-only)
O_WRONLY：只写模式(write-only)
O_RDWR：读写模式(read-write)
O_APPEND：追加模式(append)
O_CREATE：文件不存在就创建(create a new file if none exists.)
O_EXCL：与 O_CREATE 一起用，构成一个新建文件的功能，它要求文件必须不存在(used with O_CREATE, file must not exist)
O_SYNC：同步方式打开，即不使用缓存，直接写入硬盘
O_TRUNC：打开并清空文件
至于操作权限perm，除非创建文件时才需要指定，不需要创建新文件时可以将其设定为０.虽然go语言给perm权限设定了很多的常量，但是习惯上也可以直接使用数字，如0666(具体含义和Unix系统的一致).

6: io.WriteString(w Writer, s string) (n int, err error)

*/


/*
//Log包

//1.定义logger 文件，前缀字符串，flag标记
func New(out io.Writer,prefix string,flag int) *Logger

//2.设置flag格式
func SetFlags(flag int)
Ldate         = 1 << iota     // the date: 2009/01/23 形如 2009/01/23 的日期
Ltime                         // the time: 01:23:23   形如 01:23:23   的时间
Lmicroseconds                 // microsecond resolution: 01:23:23.123123.  形如01:23:23.123123   的时间
Llongfile                     // full file name and line number: /a/b/c/d.go:23 全路径文件名和行号
Lshortfile                    // final file name element and line number: d.go:23. overrides Llongfile 文件名和行号
LstdFlags     = Ldate | Ltime // 日期和时间

//3.配置log的输出格式
func SetPrefix(prefix string)
*/