package common

import (
	"crypto/md5"
	"database/sql"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"strings"
)

func Md5(text string) string {
	hashMd5 := md5.New()
	io.WriteString(hashMd5, text)
	return fmt.Sprintf("%x", hashMd5.Sum(nil))
}

/*
获取IP
*/
func GetIp(r *http.Request) string {
	ip := net.ParseIP(strings.Split(r.RemoteAddr, ":")[0]).String()
	if ip == "<nil>" {
		ip = "127.0.0.1"
	}
	return ip
}

/*
组合数据 原转JSON(已修正不需要转JSON)
*/
//FIXME 此处方法名需要重新命名 否则会产生干扰
func SaveMapData(key string, data interface{}) interface{} {
	dataMap := make(map[string]interface{})
	dataMap[key] = data
	return dataMap
}

/*
把查询数据库的结果集转换成map
*/
func ResToMap(rows *sql.Rows) map[string]string {
	data := make(map[string]string)
	columns, err := rows.Columns()
	if err != nil {
		log.Println("获取结果集中列名数组错误:", err)
	}
	values := make([]sql.RawBytes, len(columns))
	scanArgs := make([]interface{}, len(values))
	for i := range values {
		scanArgs[i] = &values[i]
	}
	for rows.Next() {
		err = rows.Scan(scanArgs...)
		if err != nil {
			log.Println("扫描结果集中参数值错误:", err)
		}
		var value string
		for i, col := range values {
			if col == nil {
				value = "NULL"
			} else {
				value = string(col)
			}
			data[columns[i]] = value
		}

	}
	return data
}
