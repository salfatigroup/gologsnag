<div align="center">
	<img src="https://logsnag.com/og-image.png" alt="LogSnag"/>
	<br>
    <h1>LogSnag</h1>
	<p>Get notifications and track your project events. This package is maintained by the <a href="https://salfati.group">salfati.group</a> and <a href="https://twitter.com/elonsalfati">@elonsalfati</a></p>
	<a href="https://discord.gg/dY3pRxgWua"><img src="https://img.shields.io/discord/922560704454750245?color=%237289DA&label=Discord" alt="Discord"></a>
	<a href="https://docs.logsnag.com"><img src="https://img.shields.io/badge/Docs-LogSnag" alt="Documentation"></a>
	<br>
	<br>
</div>


## Installation

```sh
go get -u github.com/salfatigroup/gologsnag
```

## Usage

### Import Library

```go
import "github.com/salfatigroup/gologsnag"
```

### Initialize Client

```go
logsnag := gologsnag.NewLogSnag("7f568d735724351757637b1dbf108e5", "my-saas")
```

### Publish Event

```go
logsnag.Publish(
    context.Background(),
    &gologsnag.PublishOptions{
        Channel: "waitlist",
        Event: "User Joined",
        Description: "Email: john@doe.com",
        Icon: "🎉",
        Tags: &gologsnag.Tags{
            "email": "john@doe.com",
            "user-id": "uid-12",
        },
        Notify: true,
    },
)
```
