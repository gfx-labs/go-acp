# go-acp

Go implementation of the [Agent Client Protocol
(ACP)](https://agentclientprotocol.com): JSON-RPC 2.0 over
newline-delimited messages, with the protocol's types generated from the
pinned schema in `schema/` (1.31.0).

```sh
go get github.com/gfx-labs/go-acp
```

```go
import acp "github.com/gfx-labs/go-acp"
```

- `NewAgentSideConnection(agent, w, r, logger)`: the agent's end. It calls
  your `acp.Agent` for each client request and its methods call the client.
- `NewClientSideConnection(client, w, r, logger)`: the client's end. It
  calls your `acp.Client` for agent requests and `session/update`, and its
  methods (`Initialize`, `NewSession`, `LoadSession`, `Prompt`, `Cancel`,
  ...) call the agent.
- `Serve(ctx)` reads until the reader ends. Requests run on their own
  goroutine. Notifications and responses are handled in order, so
  `session/update` arrives in order. `Arrival(ctx)` gives a message's
  position in the read order, for ordering a `session/cancel` against
  prompts.
- `acptest` drives an agent from a test with raw JSON lines.

Requires Go 1.27 (`encoding/json/v2`). Dependencies: `valyala/bytebufferpool`,
and `stretchr/testify` in tests.

## Generated code

The `*_gen.go` files come from `internal/acpgen`, a modified copy of the
generator in [coder/acp-go-sdk](https://github.com/coder/acp-go-sdk)
(Apache-2.0). Its README lists the changes. Regenerate after changing the
schema or the generator:

```sh
go generate .
```

CI fails when the committed files differ from what the generator writes.

## Origin

Extracted from `pkg/acp` in towa, with history. Design notes are in
`docs/notes/`.

## Check

```sh
go build ./... && go vet ./... && go test ./... && golangci-lint run ./...
(cd internal/acpgen && go vet ./... && go test ./...)
```

## License

Apache-2.0.
