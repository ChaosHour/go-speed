# go-speed

## Usage

If you want to use go-speed, you need to install it.

```GO
go get -u github.com/mattn/go-speed
```

or use it in your project.

```GO

    package main

    import (
        "fmt"
        "github.com/go-speed/go-speed"
    )

    func main() {
        fmt.Println(go-speed.Get())
    }

```

```GO

To compile  go-speed.go:

    go build -o go-speed

FreeBSD:
env GOOS=freebsd GOARCH=amd64 go build .

On Mac:
env GOOS=darwin GOARCH=amd64 go build .
```

```GO

Output Example:
go-speed

https://www.google.com: 293.620142ms
https://www.cox.com: 1.094964301s
https://www.facebook.com: 279.980286ms
https://www.youtube.com: 308.856589ms
https://www.amazon.com: 810.044073ms
https://www.reddit.com: 1.229588612s
https://www.wikipedia.org: 103.765133ms
https://www.instagram.com: 270.526351ms
https://www.twitter.com: 391.325277ms
https://www.linkedin.com: 333.494767ms
https://www.baidu.com: 723.79939ms
https://www.qq.com: 73.842124ms
https://www.taobao.com: 912.423393ms
https://www.live.com: 438.002218ms
https://www.tmall.com: 859.089674ms
https://www.yahoo.com: 597.189006ms
https://www.sohu.com: 297.150351ms
https://www.sina.com.cn: 1.503742597s
https://www.pinterest.com: 394.630026ms
https://www.microsoft.com: 228.662429ms
https://www.apple.com: 91.357893ms
https://www.ebay.com: 419.924923ms
https://www.paypal.com: 297.218949ms
https://www.microsoft.com/en-us/store/apps: 1.40908225s
https://www.microsoft.com/en-us/windows/microsoft-edge: 251.929013ms
https://www.groupon.com: 1.958095984s

```

### Thank you! [Github Copilot](https://copilot.github.com/)
