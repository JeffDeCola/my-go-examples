package tcp_formdata

import (
	"bytes"
	"crypto/sha1"
	"errors"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strconv"
	"time"
)

var conType map[string]string = map[string]string{
	"png": "image/png",
	"jpg": "image/jpeg",
	"svg": "text/xml",
	"txt": "text/plain",
	"zip": "application/x-zip-compressed",
	"mp4": "video/mpeg4",
	"pdf": "application/pdf",
	"avi": "video/avi",
	"mp3": "audio/mp3"}

type FormData struct {
	Appid  string
	Secret string
}
type QueryItem struct {
	Key   string
	Value string
}
type Query struct {
	Items map[string]*QueryItem
}

func (q *Query) Get(key string) *QueryItem {
	da, ok := q.Items[key]
	if ok {
		return da
	} else {
		return &QueryItem{}
	}
}
func (q *Query) Set(key, value string) {
	q.Items[key] = &QueryItem{
		Key:   key,
		Value: value,
	}
}
func Start(yourselves string) {
	client(yourselves)
}
func client(yourselves string) {
	l, err := net.Listen("tcp", ":7777")
	if err != nil {
		log.Println(err)
	}
	for {
		con, err := l.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		go handle(con, yourselves)
	}
}
func handle(conn net.Conn, yourselves string) {
	// ********** MAX READ FILE SIZE ************
	cache := make([]byte, 2048000)
	// ---------- MAX READ FILE SIZE ------------
	n, err := conn.Read(cache)
	fmt.Println(string(cache[:n]))
	if err != nil && err != io.EOF {
		log.Println(err, 1)
	}
	responsePage(conn, cache[:n], "/ ")
	if conn == nil {
		return
	}
	// use tcp
	// TCP protocol , this is pretty faster , fist byte must is FF(255)
	// 1-129 is filename area,its a built-in protocol,use 0 to end
	// true length 129
	//fmt.Println(123)
	if cache[0] == 0xFF {
		filename := cache[:n][1:129]
		if len(filename) < 5 {
			_, _ = conn.Write([]byte("error"))
			_ = conn.Close()
			return
		}
		for k, v := range filename {
			if v == 0 {
				filename = cache[:k]
				break
			}
		}
		contuse, _ := SplitString(filename, []byte("."))
		filer := sha(cache[129:n])
		fs, err := os.OpenFile(filer+"."+string(contuse[len(contuse)-1]), os.O_CREATE|os.O_WRONLY, 666)
		if err != nil {
			log.Println(err)
		}
		_, _ = fs.Write(cache[129:n])
		_ = fs.Close()
		_, _ = conn.Write([]byte(yourselves + filer))
		goto end
	}
	// upload
	if checkUploadOrDownload(cache[5:12]) {
		// start upload Verification
		//form := parseQuery(cache[:100])
		//if form.Get("appid").Value != "admin" && form.Get("secret").Value != "123456" {
		//	toErrorJson(conn, "appid or secret is error")
		//	return
		//}
		name, _, err := formatFile(cache[:n])
		if err != nil {
			toErrorJson(conn, err.Error())
			return
		}
		toSuccessJson(conn, yourselves+string(name))
		return
	}
	// download
	for k, v := range cache[5:] {
		if v == 32 {
			fs, err := os.Open(string(cache[5 : k+5]))
			if err != nil {
				toErrorJson(conn, err.Error())
				return
			}
			Faye, _ := SplitString(cache[5:k+5], []byte("."))
			conn.Write([]byte("HTTP/1.1 200 OK\r\n"))
			conn.Write([]byte("Server: FileServer\r\n"))
			conn.Write([]byte("Date: " + time.Now().String() + "\r\n"))
			conn.Write([]byte("Content-Type: " + conType[string(Faye[len(Faye)-1])] + "\r\n\r\n"))
			out := make([]byte, 10240)
			for {
				m, err := fs.Read(out)
				if err == io.EOF || m == 0 {
					break
				}
				_, _ = conn.Write(out[:m])
			}
			break
		}
	}
end:
	if conn != nil {
		_ = conn.Close()
	}
}
func responsePage(conn net.Conn, body []byte, prefix string) {
	switch body[0] {

	// response GET
	case 'G':
		if Equal(body[4:6], []byte(prefix)) {
			conn.Write([]byte("HTTP/1.1 200 OK\r\n"))
			conn.Write([]byte("Server: FileServer\r\n"))
			conn.Write([]byte("Date: " + time.Now().String() + "\r\n"))
			conn.Write([]byte("Content-Type: text/html\r\n\r\n"))
			conn.Write([]byte(`<!DOCTYPE html><html lang="en"><head><meta charset="UTF-8"><meta name="viewport" content="width=device-width, initial-scale=1.0"><meta http-equiv="X-UA-Compatible" content="ie=edge"><title>Welcome</title></head><body><h3>This is FileSystemServer</h3> </body></html>`))
			_ = conn.Close()
		}
	// response POST
	case 'P':
		if Equal(body[5:7], []byte(prefix)) {
			conn.Write([]byte("HTTP/1.1 200 OK\r\n"))
			conn.Write([]byte("Server: FileServer\r\n"))
			conn.Write([]byte("Date: " + time.Now().String() + "\r\n"))
			conn.Write([]byte("Content-Type: text/html\r\n\r\n"))
			conn.Write([]byte(`<!DOCTYPE html><html lang="en"><head><meta charset="UTF-8"><meta name="viewport" content="width=device-width, initial-scale=1.0"><meta http-equiv="X-UA-Compatible" content="ie=edge"><title>Welcome</title></head><body><h3>This is FileSystemServer</h3> </body></html>`))
			_ = conn.Close()
		}
	// response other
	default:
	}
}
func checkUploadOrDownload(body []byte) bool {
	if Equal(body, []byte{47, 117, 112, 108, 111, 97, 100}) {
		return true
	}
	return false
}
func toSuccessJson(conn net.Conn, body string) {
	if conn != nil {
		conn.Write([]byte("HTTP/1.1 200 OK\r\n"))
		conn.Write([]byte("Server: FileServer\r\n"))
		conn.Write([]byte("Date: " + time.Now().String() + "\r\n"))
		conn.Write([]byte("Content-Type: Application/json\r\n\r\n"))
		conn.Write([]byte(`{"t":1,"ok"":"yes","msg":"success","data":"` + body + `"}`))
		conn.Close()
	}
	return
}
func toErrorJson(conn net.Conn, msg string) {
	if conn != nil {
		conn.Write([]byte("HTTP/1.1 503 Service Unavailable\r\n"))
		conn.Write([]byte("Server: FileServer\r\n"))
		conn.Write([]byte("Date: " + time.Now().String() + "\r\n"))
		conn.Write([]byte("Content-Type: Application/json\r\n\r\n"))
		conn.Write([]byte(`{"t":0,"ok"":"no","msg":"` + msg + `","data":""}`))
		conn.Close()
	}
	return
}
func Equal(one []byte, two []byte) bool {
	if len(one) != len(two) {
		return false
	}
	for k, v := range one {
		if v != two[k] {
			return false
		}
	}
	return true
}
func SplitString(str []byte, p []byte) ([][]byte, []int) {
	group := make([][]byte, 0)
	portion := make([]int, 0)
	ps := 0
	for i := 0; i < len(str); i++ {
		if str[i] == p[0] && i < len(str)-len(p) {
			if len(p) == 1 {
				group = append(group, str[ps:i])
				portion = append(portion, ps)
				ps = i + len(p)
			} else {
				for j := 1; j < len(p); j++ {
					if str[i+j] != p[j] || j != len(p)-1 {
						continue
					} else {
						group = append(group, str[ps:i])
						portion = append(portion, ps)
						ps = i + len(p)
					}
				}
			}
		} else {
			continue
		}
	}
	group = append(group, str[ps:])
	portion = append(portion, ps)
	return group, portion
}
func sha(data []byte) string {
	t := sha1.New()
	t.Write(data)
	return fmt.Sprintf("%x", t.Sum(nil))
}
func formatFile(data []byte) ([]byte, int, error) {
	// boundary name
	var boundary []byte
	var boundaryPoint = false
	// file size
	var fileSize int
	var fileSizeSilce = make([]byte, 0)
	// file name
	var fileName []byte = make([]byte, 0)
	var fileNamePoint bool
	var fileSizePoint bool
	var seek = make([]int, 0)
	var trueData int = 0
	for k, v := range data {
		// find boundary name, out => boundary name
		if !boundaryPoint {
			if v == 'b' {
				if data[k+1] == 'o' && data[k+2] == 'u' && data[k+3] == 'n' {
					for _, c := range data[k+9:] {
						if c == '\r' {
							break
						}
						boundary = append(boundary, c)
					}
					boundaryPoint = true
				}
			}
		}
		// find file size,out => fileSize
		if !fileSizePoint {
			if v == 'C' {
				if data[k+8] == 'L' && data[k+9] == 'e' && data[k+10] == 'n' {
					for _, c := range data[k+16:] {
						if c == '\r' {
							break
						}
						fileSizeSilce = append(fileSizeSilce, c)
					}
					fileSizePoint = true
					fileSize, _ = strconv.Atoi(string(fileSizeSilce))
				}
			}
		}
		if boundaryPoint {
			if len(seek) >= 3 {
				break
			}
			if bytes.Equal(boundary, data[k:k+len(boundary)]) {
				seek = append(seek, k)
			}
		}
	}
	if len(seek) < 3 {
		return nil, 0, errors.New("FormData parsing error : Seek length is less than 3")
	}
	for i := 0; i < len(data[seek[1]:seek[2]]); i++ {
		if i > 300 {
			break
		}
		if data[seek[1]:seek[2]][i] == 'f' {
			if !fileNamePoint && data[seek[1]:seek[2]][i+4] == 'n' && data[seek[1]:seek[2]][i+5] == 'a' {
				for _, v := range data[seek[1]:seek[2]][i+10:] {
					if v == 13 {
						break
					}
					fileName = append(fileName, v)
				}
				fileNamePoint = true
			}
		}
		if data[seek[1]:seek[2]][i] == 13 && data[seek[1]:seek[2]][i+1] == 13 {
			trueData = i + 2
			break
		}
	}
	if trueData == 0 {
		for i := 0; i < len(data[seek[1]:seek[2]]); i++ {
			if i > 300 {
				break
			}
			if data[seek[1]:seek[2]][i] == 'f' {
				if !fileNamePoint && data[seek[1]:seek[2]][i+4] == 'n' && data[seek[1]:seek[2]][i+5] == 'a' {
					for _, v := range data[seek[1]:seek[2]][i+10:] {
						if v == 13 {
							break
						}
						fileName = append(fileName, v)
					}
					fileNamePoint = true
				}
			}
			if data[seek[1]:seek[2]][i] == 13 && data[seek[1]:seek[2]][i+1] == 10 && data[seek[1]:seek[2]][i+2] == 13 && data[seek[1]:seek[2]][i+3] == 10 {
				trueData = i + 4
				break
			}
		}
	}
	if trueData == 0 {
		return nil, 0, errors.New("FormData parsing error : trueData is 0")
	}
	files, _ := SplitString(fileName[:len(fileName)-1], []byte{'.'})
	outName := sha(data[seek[2]/2 : seek[2]])
	fmt.Println("boundary name =>", string(boundary))
	//fmt.Println(trueData)
	//fmt.Println(fileSize)
	//fmt.Println(string(fileName))
	//fmt.Println(seek)
	// 239 191 189 239 191
	//fmt.Println((data[seek[1]:seek[2]][trueData : seek[2]-175-len(boundary)]))
	//fmt.Println(outName + "." + string(files[1]))
	//log.Println("upload :", string(fileName[:len(fileName)-1]))
	fs, err := os.OpenFile(outName+"."+string(files[1]), os.O_CREATE|os.O_WRONLY, 666)
	if err != nil {
		return nil, 0, err
	}
	_, err = fs.Write(data[seek[1]:seek[2]][trueData : seek[2]-175-len(boundary)])
	if err != nil {
		return nil, 0, err
	}
	_ = fs.Close()
	return []byte(outName + "." + string(files[1])), fileSize, nil
}
func parseQuery(body []byte) *Query {
	q := new(Query)
	q.Items = make(map[string]*QueryItem)
	for k, v := range body {
		if v == 63 {
			for j, c := range body[k+1:] {
				if c == 32 {
					d, _ := SplitString(body[k+1:j+k+1], []byte{38})
					if len(d) == 0 {
						return nil
					}
					for _, v := range d {
						c, _ := SplitString(v, []byte{61})
						if len(c) == 0 {
							continue
						}
						q.Set(string(c[0]), string(c[1]))
					}
					break
				}
			}
			break
		}
	}
	return q
}
