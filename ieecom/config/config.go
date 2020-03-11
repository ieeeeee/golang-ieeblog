/*解析web.config, 即解析一个配置文件规则的XML文件*/

package config
import(
	"fmt"
	"encoding/xml"
	"os"
	"io/ioutil"
)


//Config define a config file struct. main structure
//声明一个配置文件结构体，配置文件的结构
type Config struct{
	XMLName xml.Name `xml:appSettings` //不能省
	WebURL string `xml:"webUrl"`
	WebPort string `xml:"port"`
	//mssql
	Mssql struct{
		XMLName xml.Name `xml:"mssql"`  //不能省，否则整个struct读出来是空的
		DataSource string `xml:"dataSource"`
		Database string `xml:"database"`
		Windows bool `xml:"windows"`
		User string `xml:"user"`
		Password string `xml:"password"`
	}
	LogPath string `xml:"logPath"`
	StaticPath string `xml:"staticPath"`
}

//Conf Define a  global variant
//声明一个全局变量
var Conf *Config

//InitConfig public func
//初始化配置文件
func InitConfig(filename string){
	
	//1.打开配置文件
	configFile,err:=os.Open(filename)
	defer configFile.Close()
	if err!=nil{
		fmt.Printf("error 1 open:%v",err)
		return
	}
	//2.读取文件
	data,err:=ioutil.ReadAll(configFile) 
	if err!=nil{
		fmt.Printf("error 2 readAll:%v",err)
		return
	}

	//3.解析配置xml文件到全局变量Conf
	Conf=&Config{}
	err=xml.Unmarshal(data,Conf)
	if err!=nil{
		fmt.Printf("error 3 xml unmarshal:%v",err)
		return
	}
	//fmt.Println(Conf)
}

//调用例如 Config.Conf.Mssql.DataSource
