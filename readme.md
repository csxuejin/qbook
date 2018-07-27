### 使用七牛存储空间创建 gitbook 静态站点

- 创建一个 bucket， 在空间设置中打开 “默认首页设置”
- `go get github.com/csxuejin/kodo`
- `go get github.com/csxuejin/qbook`
- 编辑 `$GOPATH/src/github.com/csxuejin/qbook/kodo.json` 文件，填写以下几个字段

``` go
"book_dir": gitbook 生成的 _book 目录路径
"access_key": 七牛账户的 ak
"secret_key": 七牛账户的 sk
"bucket": bucket 名称
```
