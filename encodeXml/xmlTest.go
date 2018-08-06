package main

import (
	"encoding/xml"

	"fmt"
)

type Servers struct {
	XMLName xml.Name `xml:"servers,omitempty"`
	Version string   `xml:"version,attr"`
	Svs     []server `xml:"server,omitempty"`
}

type server struct {
	ServerName string `xml:"serverName,omitempty"`
	ServerIP   ServerIP `xml:"serverIP,omitempty"`
}

type ServerIP struct {
	Id string `xml:"id,attr"`
	Value string `xml:",cdata"'`
}

func main() {
	v := &Servers{Version: "1"}
	v.Svs = append(v.Svs, server{ServerName:"Shanghai_VPN"})
	v.Svs = append(v.Svs, server{ServerName:"Beijing_VPN", ServerIP:ServerIP{Value:"127.0.0.2"}})


	output, err := xml.MarshalIndent(v, "", "	")
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}
	myString := []byte(xml.Header + string(output))
	//将字节流转换成string输出
	fmt.Println(string(myString))

}

type Table struct {
	TableName string `xml:"-"`
	Version   string `xml:"version,attr"`
	Size      int    `xml:"size,attr"`
	Line      []Line `xml:"line"`
}

type Line struct {
	Title   string `xml:"title,attr"`
	Content string `xml:",chardata"`
}


