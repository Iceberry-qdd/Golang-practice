/**
 * Created by visual studio code
 * @author Iceberry
 * @date 2022-1-5
 * @version 1.0
 * 单元测试
 * */
package test

import "strings"

func Split(s, sep string) (result []string) {
	i := strings.Index(s, sep)
	for i > -1 {
		result = append(result, s[:i])
		s = s[i+len(sep):]
		i = strings.Index(s, sep)
	}
	result = append(result, s)
	return
}
