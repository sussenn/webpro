package tool

import (
	"bufio"
	"github.com/tedcy/fdfs_client"
	"log"
	"os"
	"strings"
)

func UploadFile(fileName string) string {
	client, err := fdfs_client.NewClientWithConfig("./config/fastdfs.conf")
	defer client.Destory()
	if err != nil {
		log.Println("fdfs_client.NewClientWithConfig() FastDFS文件上传异常. err: ", err)
		return ""
	}
	fileId, err := client.UploadByFilename(fileName)
	if err != nil {
		log.Println("client.UploadByFilename() FastDFS文件上传异常. err: ", err)
		return ""
	}
	return fileId
}

//读取.conf文件配置的IP和端口号信息
func FileServerAddr() string {
	file, err := os.Open("./config/fastdfs.conf")
	if err != nil {
		return ""
	}
	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		line = strings.TrimSuffix(line, "\n")
		str := strings.SplitN(line, "=", 2)
		switch str[0] {
		case "http_server_port":
			return str[1]
		}
		if err != nil {
			return ""
		}
	}
}
