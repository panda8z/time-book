# 新建

## 1. 过程

```bash
mkdir time-book
go mod init github.com/panda8z/time-book
git init 
git remote add origin git@github.com:panda8z/time-book.git
git add -A
git commit -m "init"
git push --set-upstream origin master
```

## 2. 目录结构

```bash

```

## 3. 加入 fiber 运行 hello world

### 3.1 installation

`go get -u github.com/gofiber/fiber`

### 3.2 hello world

```go
package main

import "github.com/gofiber/fiber"

func main() {
  app := fiber.New()

  app.Get("/", func(c *fiber.Ctx) {
    c.Send("Hello, World!")
  })

  app.Listen(3000)
}
```


```bash
go run server.go
```

浏览器查看: [http://localhost:3000](http://localhost:3000)

