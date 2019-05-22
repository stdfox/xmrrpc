# xmrrpc
Golang client for Monero (XMR) RPC API.

Support daemon's RPC methods. Support digest authentication.

Tested on: Monero 'Boron Butterfly' (v0.14.0.2-release), stagenet.

## Documentation
Full API Documentation can be found at https://web.getmonero.org/resources/developer-guides/

Current API version: network height of 1,562,465.

### JSON RPC Methods:
- get_block_count
- on_get_block_hash
- get_block_template
- submit_block
- get_last_block_header
- get_block_header_by_hash
- get_block_header_by_height
- get_block_headers_range
- get_block
- get_connections
- get_info
- hard_fork_info
- set_bans
- get_bans
- flush_txpool
- get_output_histogram
- get_version
- get_coinbase_tx_sum
- get_fee_estimate
- get_alternate_chains
- relay_tx
- sync_info
- get_txpool_backlog

### RPC Methods:
- get_height
- get_transactions
- get_alt_blocks_hashes
- is_key_image_spent
- send_raw_transaction
- start_mining
- stop_mining
- mining_status
- save_bc
- get_peer_list
- set_log_hash_rate
- set_log_level
- set_log_categories
- get_transaction_pool_stats
- update

## Installation
```bash
$ go get github.com/stdfox/xmrrpc
```

## Importing
```go
import (
    "github.com/stdfox/xmrrpc"
)
```

## Example
```go
package main

import (
	"fmt"

	"github.com/stdfox/xmrrpc"
)

func main() {
	daemonClient := xmrrpc.NewDaemonClient("http://127.0.0.1:38081", "username", "password")

	res1, err := daemonClient.GetLastBlockHeader()
	if err != nil {
		panic(err)
	}

	fmt.Printf("[Status: %s] Hash: %s\n", res1.Status, res1.BlockHeader.Hash)

	res2, err := daemonClient.GetHeight()
	if err != nil {
		panic(err)
	}

	fmt.Printf("[Status: %s] Height: %d\n", res2.Status, res2.Height)
}
```
```shell
$ go run main.go
[Status: OK] Hash: 0a5389d9521ea4a049375aca17875d5178fa1978b6e681c9c8f968cc12e3a501
[Status: OK] Height: 321885
```

## License
Licensed under [MIT License](https://github.com/stdfox/xmrrpc/blob/master/LICENSE.md).

## Donation ❤
If this project help you, you can give me a cup of coffee :)

XMR: `49Nz5mVA9sQjKrT65TgdEneiZo1oCp3n8bCtjA3qaCoa5cuPKxqWBcZfD1f1iv6ASjCQUK55m3r4iho7ivMcNvsLDnP3sqX`

BTC: `1Kyxw87175msSRkXez7iKrzZLBaPb7rp1H`
