# gomail
## golang 实现的邮件发送


#### 编译
```
go  get  -u  github.com/Bestfeel/gomail

```

### 运行

```
gomail
  -a string
        mail server address:[smtp.exmail.qq.com:25]
  -b string
        body of mail or  fileHtml
  -c string
        mail send cc 
  -f string
        mail send from 
  -o string
        mail server host:[smtp.exmail.qq.com]
  -p string
        SMTP server password 
  -s string
        subject of the mail
  -t string
        mail send to 
  -u string
        SMTP server username 

```


### 发送文本或者文件

```
gomail  -s="test" -b="/home/feel/README.md"  -a="addr"   -o="host"  -u="username"  -p="password"  -f="from"  -t="to" -c="cc"
```

