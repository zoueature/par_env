/*+---------------------------------+
 *| Author: Zoueature               |
 *+---------------------------------+
 *| Email: zoueature@gmail.com      |
 *+---------------------------------+
 *| Date: 2019-11-03 17:40          |
 *+---------------------------------+
 */

package parenv

import (
	"io/ioutil"
	"os"
	"strings"
)

var env = make(map[string]string)
var hasParsed = false

func EnvInit(filePath string) {
	reader, e := os.Open(filePath)
	if e != nil {
		panic("open env file error : " + e.Error())
	}
	content, e := ioutil.ReadAll(reader)
	if e != nil {
		panic("read env info from file error : " + e.Error())
	}
	env = parse(string(content))

}

func parse(content string) map[string]string {
	env := make(map[string]string)
	stringSlice := strings.Split(content, "\n")
	if stringSlice == nil {
		panic("error, parse env string error")
	}
	for _, str :=  range stringSlice {
		envInfo := strings.Split(str, "=")
		if envInfo == nil {
			panic("error")
		}
		if len(envInfo) < 2 {
			panic("format error : " + str)
		}
		env[strings.Trim(envInfo[0], " ")] = strings.Trim(envInfo[1], " ")
	}
	hasParsed = true
	return env
}

func Get(key string, defaultValue string) string {
	value, ok := env[key]
	if !ok {
		return defaultValue
	}
	return value
}

func Set(key string, value string) {
	env[key] = value
}

func All() map[string]string {
	return env
}