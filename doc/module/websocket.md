# WebSocket

This is a module for [Project Forge](https://projectforge.dev). It provides a WebSocket API

https://github.com/kyleu/projectforge/tree/main/module/websocket

### License

Licensed under [CC0](https://creativecommons.org/publicdomain/zero/1.0)

### Usage

To create a WebSocket connection, you'll need to create two controller actions and routes.

#### Controller
```go
func Example(w http.ResponseWriter, r *http.Request) {
	Act("example", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		ps.SetTitleAndData("Example Test", "TODO")
		return Render(r, as, &views.Example{}, ps)
	})
}

func ExampleSocket(w http.ResponseWriter, r *http.Request) {
	Act("example.socket", w, r, func(as *app.State, ps *cutil.PageState) (string, error) {
		channel := r.URL.Query().Get("ch")
		if channel == "" {
			channel = util.RandomString(16)
		}
		id, err := as.Services.Socket.Upgrade(...)
		if err != nil {
			return "", err
		}
		return "", as.Services.Socket.ReadLoop(ps.Context, id, ps.Logger)
	})
}

```

#### Routes
```go
makeRoute(r, http.MethodGet, "/example", controller.Example)
makeRoute(r, http.MethodGet, "/example/socket", controller.ExampleSocket)
```

Then, in TypeScript or your template, you'll need to include the following script:

```js
const ch = "example";
let sock;

function open() {
  console.log("[socket]: open");
}

function recv(m) {
  const list = document.getElementById("messages");
  if (list) {
    const pre = document.createElement("pre");
    pre.innerText = JSON.stringify(m, null, 2);
    list.append(pre);
  }
  handle(svc, m);
}

function err(e) {
  console.log(`[socket error]: ${e}`);
}

function send(cmd, param) {
  sock.send({channel: ch, cmd: cmd, param: param});
}

document.addEventListener("DOMContentLoaded", function() {
  sock = new projectforge.Socket(true, open, recv, err, "/example/connect");
  console.log("loaded socket connection");
});

```
