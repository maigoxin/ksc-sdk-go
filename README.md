这是一个非官方的金山云openapi的golang签名包，在go version go1.8.3 darwin/amd64环境下测试通过。

获取方式

```shell
go get github.com/maigoxin/ksc-sdk-go
```



Demo1

```go
func main() {
    host := "http://iam.api.ksyun.com"
    client := new(http.Client)

    v := url.Values{}
    v.Add("Action", "ListUsers")
    v.Add("Version", "2015-11-01")
    req, _ := http.NewRequest("POST", host, strings.NewReader(v.Encode()))

    ksc.Sign(req, ksc.Credentials{
        AccessKeyID:     "this is your ak",
        SecretAccessKey: "this is your sk",
        Service:         "iam",
        Region:          "cn-beijing-6",
    })  

    resp, _ := client.Do(req)

    responseData, _ := ioutil.ReadAll(resp.Body)
    fmt.Println(string(responseData))
}
```



Demo2

```go
func main() {
    host := "http://iam.api.ksyun.com/?Action=ListUsers&Version=2015-11-01"
    client := new(http.Client)

    req, _ := http.NewRequest("GET", host, nil)

    ksc.Sign(req, ksc.Credentials{
        AccessKeyID:     "this is your ak",
        SecretAccessKey: "this is your sk",
        Service:         "iam",
        Region:          "cn-beijing-6",
    })

    resp, _ := client.Do(req)

    responseData, _ := ioutil.ReadAll(resp.Body)
    fmt.Println(string(responseData))
}
```



> 如果有使用的问题，可以邮件联系