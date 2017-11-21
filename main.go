package main

import (
	"net/smtp"
	"flag"
	"fmt"
	"strings"
	"bytes"
	"mime"
	"os"
	"log"
	"io/ioutil"
)

type MailClient struct {
	addr string // smtp.example.com:25
	auth smtp.Auth
	from string
}

//Message 邮件发送数据
type Message struct {
	Subject string   // 标题
	Content []byte   // 支持html的消息主体
	To      []string // 邮箱地址,发送给谁,多方发送
	Cc      []string // 抄送
}

type Sender interface {
	Send(msg *Message)
}

func (this *MailClient) GetBody(msg *Message) []byte {

	buffer := new(bytes.Buffer)
	fmt.Fprintf(buffer, "%s: %s\r\n", "To", strings.Join(msg.To, ","))
	fmt.Fprintf(buffer, "%s: %s\r\n", "Cc", strings.Join(msg.Cc, ","))
	fmt.Fprintf(buffer, "%s: %s\r\n", "From", this.from)
	fmt.Fprintf(buffer, "%s: %s\r\n", "subject", mime.QEncoding.Encode("utf-8", msg.Subject))
	fmt.Fprintf(buffer, "%s: %s\r\n", "Content-Type", "text/html;charset=utf-8")
	fmt.Fprintf(buffer, "%s: %s\r\n", "Mime-Version", "1.0")
	fmt.Fprintf(buffer, "%s: %s\r\n", "Content-Transfer-Encoding", "Quoted-Printable")
	buffer.WriteString("\r\n")
	buffer.Write(msg.Content)
	buffer.WriteString("\r\n")

	return buffer.Bytes()
}

/**
 实现邮件发送接口
 */
func (this *MailClient) Send(msg *Message) error {

	return smtp.SendMail(this.addr, this.auth, this.from, msg.To, this.GetBody(msg))
}


func exist(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil || os.IsExist(err)
}

var (
	subject  = flag.String("s", "", "subject of the mail")
	content  = flag.String("b", "", "body of mail or  fileHtml")
	addr     = flag.String("a", "", "mail server address:[smtp.exmail.qq.com:25]")
	host     = flag.String("o", "", "mail server host:[smtp.exmail.qq.com]")
	username = flag.String("u", "", "SMTP server username ")
	password = flag.String("p", "", "SMTP server password ")
	from     = flag.String("f", "", "mail send from ")
	to       = flag.String("t", "", "mail send to ")
	cc       = flag.String("c", "", "mail send cc ")
)

func main() {
	args := os.Args[1:]
	if len(args) < 8 {
		flag.PrintDefaults()
		return
	}

	flag.Parse()
	var body []byte
	value := *content
	if exist(value) {

		data, err := ioutil.ReadFile(value)
		if err != nil {
			return
		} else {
			body = data
		}

	} else {
		body = []byte(value)
	}

	message := Message{*subject, body, strings.Split(*to, ","), strings.Split(*cc, ",")}
	client := MailClient{*addr, smtp.PlainAuth("", *username, *password, *host), *from}
	err := client.Send(&message)
	if err != nil {
		log.Fatal(err)
	}

}
