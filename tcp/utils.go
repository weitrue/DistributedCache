/**
 * Author: Wang P
 * Version: 1.0.0
 * Date: 2020/12/23 下午9:32
 * Description:
 **/

package tcp

import (
	"bufio"
	"fmt"
	"net"
	"strconv"
	"strings"
)

func readLen(r *bufio.Reader) (int, error) {
	tmp, err := r.ReadString(' ')
	if err != nil {
		return 0, err
	}
	// strconv.Atoi(s string)string转为int
	// strings.TrimSpace(s string)会返回一个string类型的slice，并将最前面和最后面的ASCII定义的空格去掉，中间的空格不会去掉，如果遇到了\0等其他字符会认为是非空格
	len, err := strconv.Atoi(strings.TrimSpace(tmp))
	if err != nil {
		return 0, err
	}
	return len, nil
}

func sendResponse(value []byte, err error, conn net.Conn) error {
	if err != nil {
		errString := err.Error()
		tmp := fmt.Sprintf("-%d ", len(errString)) + errString
		_, err := conn.Write([]byte(tmp))
		return err
	}
	vlen := fmt.Sprintf("%d ", len(value))
	_, err = conn.Write(append([]byte(vlen), value...))
	return err
}
