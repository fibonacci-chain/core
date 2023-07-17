module github.com/fibonacci-chain/core

go 1.18

require (
	cosmossdk.io/api v0.6.1
	cosmossdk.io/core v0.9.0
	cosmossdk.io/depinject v1.0.0-alpha.3
	cosmossdk.io/errors v1.0.0-beta.7
	cosmossdk.io/log v1.1.0
	cosmossdk.io/math v1.0.1
	cosmossdk.io/store v1.0.0-alpha.1
	github.com/CosmWasm/wasmd v0.27.0
	github.com/CosmWasm/wasmvm v1.0.1
	github.com/armon/go-metrics v0.4.0
	github.com/btcsuite/btcd v0.22.1
	github.com/btcsuite/btcd/btcec/v2 v2.3.2
	github.com/btcsuite/btcd/chaincfg/chainhash v1.0.1
	github.com/cometbft/cometbft v0.38.0-rc2
	github.com/cosmos/cosmos-db v1.0.0
	github.com/cosmos/cosmos-proto v1.0.0-beta.3
	github.com/cosmos/cosmos-sdk v0.45.10
	github.com/cosmos/go-bip39 v1.0.0
	github.com/cosmos/gogoproto v1.4.10
	github.com/cosmos/iavl v1.0.0-beta.2
	github.com/cosmos/ibc-go/v3 v3.0.0
	github.com/davecgh/go-spew v1.1.1
	github.com/decred/dcrd/dcrec/secp256k1/v4 v4.2.0
	github.com/ethereum/go-ethereum v1.10.26
	github.com/gogo/protobuf v1.3.3
	github.com/golang/protobuf v1.5.3
	github.com/gorilla/mux v1.8.0
	github.com/gorilla/websocket v1.5.0
	github.com/grpc-ecosystem/grpc-gateway v1.16.0
	github.com/holiman/uint256 v1.2.0
	github.com/improbable-eng/grpc-web v0.14.1
	github.com/k0kubun/pp/v3 v3.2.0
	github.com/manifoldco/promptui v0.9.0
	github.com/onsi/ginkgo/v2 v2.1.3
	github.com/onsi/gomega v1.19.0
	github.com/pkg/errors v0.9.1
	github.com/rs/cors v1.8.3
	github.com/rs/zerolog v1.29.1
	github.com/spf13/cast v1.5.1
	github.com/spf13/cobra v1.6.1
	github.com/spf13/pflag v1.0.5
	github.com/spf13/viper v1.15.0
	github.com/status-im/keycard-go v0.0.0-20190316090335-8537d3370df4
	github.com/stretchr/testify v1.8.4
	github.com/syndtr/goleveldb v1.0.1-0.20220721030215-126854af5e6d
	github.com/tendermint/tendermint v0.37.0-dev
	github.com/tendermint/tm-db v0.6.8-0.20220519162814-e24b96538a12
	github.com/tidwall/gjson v1.6.7
	github.com/tidwall/sjson v1.1.4
	github.com/tyler-smith/go-bip39 v1.0.2
	go.opentelemetry.io/otel v1.14.0
	go.opentelemetry.io/otel/sdk v1.14.0
	go.opentelemetry.io/otel/trace v1.14.0
	golang.org/x/crypto v0.11.0
	golang.org/x/exp v0.0.0-20230522175609-2e198f4a06a1
	golang.org/x/net v0.12.0
	golang.org/x/text v0.11.0
	google.golang.org/genproto/googleapis/api v0.0.0-20230530153820-e85fd2cbaebc
	google.golang.org/grpc v1.56.2
	google.golang.org/protobuf v1.31.0
	gopkg.in/yaml.v2 v2.4.0
	pgregory.net/rapid v0.5.5
	sigs.k8s.io/yaml v1.3.0
)

require (
	filippo.io/edwards25519 v1.0.0-beta.2 // indirect
	github.com/99designs/keyring v1.1.6 // indirect
	github.com/DataDog/zstd v1.5.5 // indirect
	github.com/StackExchange/wmi v0.0.0-20180116203802-5d049714c4a6 // indirect
	github.com/VictoriaMetrics/fastcache v1.6.0 // indirect
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/bgentry/speakeasy v0.1.0 // indirect
	github.com/cespare/xxhash v1.1.0 // indirect
	github.com/cespare/xxhash/v2 v2.2.0 // indirect
	github.com/chzyer/readline v0.0.0-20180603132655-2972be24d48e // indirect
	github.com/cockroachdb/errors v1.10.0 // indirect
	github.com/cockroachdb/logtags v0.0.0-20230118201751-21c54148d20b // indirect
	github.com/cockroachdb/pebble v0.0.0-20230525220056-bb4fc9527b3b // indirect
	github.com/cockroachdb/redact v1.1.5 // indirect
	github.com/coinbase/rosetta-sdk-go v0.7.0 // indirect
	github.com/confio/ics23/go v0.7.0 // indirect
	github.com/cosmos/btcutil v1.0.4 // indirect
	github.com/cosmos/gorocksdb v1.2.0 // indirect
	github.com/cosmos/ics23/go v0.10.0 // indirect
	github.com/cosmos/ledger-cosmos-go v0.11.1 // indirect
	github.com/cosmos/ledger-go v0.9.2 // indirect
	github.com/creachadair/taskgroup v0.3.2 // indirect
	github.com/danieljoos/wincred v1.1.0 // indirect
	github.com/deckarep/golang-set v1.8.0 // indirect
	github.com/desertbit/timer v0.0.0-20180107155436-c41aec40b27f // indirect
	github.com/dgraph-io/badger/v3 v3.2103.2 // indirect
	github.com/dgraph-io/ristretto v0.1.1 // indirect
	github.com/dustin/go-humanize v1.0.1 // indirect
	github.com/dvsekhvalnov/jose2go v0.0.0-20200901110807-248326c1351b // indirect
	github.com/edsrzf/mmap-go v1.0.0 // indirect
	github.com/felixge/httpsnoop v1.0.2 // indirect
	github.com/fsnotify/fsnotify v1.6.0 // indirect
	github.com/gballet/go-libpcsclite v0.0.0-20190607065134-2772fd86a8ff // indirect
	github.com/getsentry/sentry-go v0.21.0 // indirect
	github.com/go-kit/kit v0.12.0 // indirect
	github.com/go-kit/log v0.2.1 // indirect
	github.com/go-logfmt/logfmt v0.6.0 // indirect
	github.com/go-logr/logr v1.2.3 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/go-ole/go-ole v1.2.6 // indirect
	github.com/go-stack/stack v1.8.0 // indirect
	github.com/godbus/dbus v0.0.0-20190726142602-4481cbc300e2 // indirect
	github.com/gogo/gateway v1.1.0 // indirect
	github.com/golang/glog v1.1.0 // indirect
	github.com/golang/groupcache v0.0.0-20210331224755-41bb18bfe9da // indirect
	github.com/golang/snappy v0.0.4 // indirect
	github.com/google/btree v1.1.2 // indirect
	github.com/google/flatbuffers v1.12.1 // indirect
	github.com/google/go-cmp v0.5.9 // indirect
	github.com/google/gofuzz v1.2.0 // indirect
	github.com/google/orderedcode v0.0.1 // indirect
	github.com/google/uuid v1.3.0 // indirect
	github.com/gorilla/handlers v1.5.1 // indirect
	github.com/grpc-ecosystem/go-grpc-middleware v1.3.0 // indirect
	github.com/grpc-ecosystem/go-grpc-prometheus v1.2.0 // indirect
	github.com/gsterjov/go-libsecret v0.0.0-20161001094733-a6f4afe4910c // indirect
	github.com/hashicorp/go-immutable-radix v1.3.1 // indirect
	github.com/hashicorp/go-metrics v0.5.1 // indirect
	github.com/hashicorp/golang-lru v0.5.5-0.20210104140557-80c98217689d // indirect
	github.com/hashicorp/golang-lru/v2 v2.0.1 // indirect
	github.com/hashicorp/hcl v1.0.0 // indirect
	github.com/hdevalence/ed25519consensus v0.0.0-20210204194344-59a8610d2b87 // indirect
	github.com/holiman/bloomfilter/v2 v2.0.3 // indirect
	github.com/huin/goupnp v1.0.3 // indirect
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/jackpal/go-nat-pmp v1.0.2 // indirect
	github.com/jmhodges/levigo v1.0.0 // indirect
	github.com/keybase/go-keychain v0.0.0-20190712205309-48d3d31d256d // indirect
	github.com/klauspost/compress v1.16.5 // indirect
	github.com/kr/pretty v0.3.1 // indirect
	github.com/kr/text v0.2.0 // indirect
	github.com/lib/pq v1.10.7 // indirect
	github.com/libp2p/go-buffer-pool v0.1.0 // indirect
	github.com/linxGnu/grocksdb v1.7.16 // indirect
	github.com/magiconair/properties v1.8.7 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.19 // indirect
	github.com/mattn/go-runewidth v0.0.9 // indirect
	github.com/matttproud/golang_protobuf_extensions v1.0.4 // indirect
	github.com/mitchellh/go-homedir v1.1.0 // indirect
	github.com/mitchellh/mapstructure v1.5.0 // indirect
	github.com/mtibben/percent v0.2.1 // indirect
	github.com/oasisprotocol/curve25519-voi v0.0.0-20220708102147-0a8a51822cae // indirect
	github.com/olekukonko/tablewriter v0.0.5 // indirect
	github.com/opencontainers/runc v1.1.5 // indirect
	github.com/pelletier/go-toml/v2 v2.0.6 // indirect
	github.com/petermattis/goid v0.0.0-20221215004737-a150e88a970d // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/prometheus/client_golang v1.16.0 // indirect
	github.com/prometheus/client_model v0.4.0 // indirect
	github.com/prometheus/common v0.44.0 // indirect
	github.com/prometheus/procfs v0.10.1 // indirect
	github.com/prometheus/tsdb v0.7.1 // indirect
	github.com/rakyll/statik v0.1.7 // indirect
	github.com/regen-network/cosmos-proto v0.3.1 // indirect
	github.com/rjeczalik/notify v0.9.1 // indirect
	github.com/rogpeppe/go-internal v1.10.0 // indirect
	github.com/sasha-s/go-deadlock v0.3.1 // indirect
	github.com/savaki/jq v0.0.0-20161209013833-0e6baecebbf8 // indirect
	github.com/sei-protocol/sei-tm-db v0.0.5 // indirect
	github.com/shirou/gopsutil v3.21.4-0.20210419000835-c7a38de76ee5+incompatible // indirect
	github.com/spf13/afero v1.9.3 // indirect
	github.com/spf13/jwalterweatherman v1.1.0 // indirect
	github.com/stretchr/objx v0.5.0 // indirect
	github.com/subosito/gotenv v1.4.2 // indirect
	github.com/tendermint/btcd v0.1.1 // indirect
	github.com/tendermint/crypto v0.0.0-20191022145703-50d29ede1e15 // indirect
	github.com/tendermint/go-amino v0.16.0 // indirect
	github.com/tidwall/btree v1.6.0 // indirect
	github.com/tidwall/match v1.0.3 // indirect
	github.com/tidwall/pretty v1.0.2 // indirect
	github.com/tklauser/go-sysconf v0.3.10 // indirect
	github.com/tklauser/numcpus v0.4.0 // indirect
	github.com/yourbasic/graph v0.0.0-20210606180040-8ecfec1c2869 // indirect
	github.com/zondax/hid v0.9.0 // indirect
	go.etcd.io/bbolt v1.3.6 // indirect
	go.opencensus.io v0.24.0 // indirect
	go.opentelemetry.io/otel/exporters/jaeger v1.9.0 // indirect
	golang.org/x/sync v0.2.0 // indirect
	golang.org/x/sys v0.10.0 // indirect
	golang.org/x/term v0.10.0 // indirect
	google.golang.org/genproto v0.0.0-20230629202037-9506855d4529 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20230706204954-ccb25ca9f130 // indirect
	gopkg.in/ini.v1 v1.67.0 // indirect
	gopkg.in/natefinch/npipe.v2 v2.0.0-20160621034901-c1b8fa8bdcce // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
	nhooyr.io/websocket v1.8.6 // indirect
)

replace (
	github.com/CosmWasm/wasmd => github.com/fibonacci-chain/fbc-wasmd v0.0.2
	github.com/confio/ics23/go => github.com/cosmos/cosmos-sdk/ics23/go v0.8.0
	github.com/cosmos/cosmos-sdk => github.com/fibonacci-chain/fbc-cosmos v0.2.48
	github.com/cosmos/iavl => github.com/fibonacci-chain/fbc-iavl v0.1.4
	github.com/cosmos/ibc-go/v3 => github.com/fibonacci-chain/fbc-ibc-go/v3 v3.1.0
	github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.3-alpha.regen.1
	github.com/keybase/go-keychain => github.com/99designs/go-keychain v0.0.0-20191008050251-8e49817e8af4
	github.com/tendermint/tendermint => github.com/sei-protocol/sei-tendermint v0.2.24
	github.com/tendermint/tm-db => github.com/fibonacci-chain/fbc-db v0.0.4
	google.golang.org/grpc => google.golang.org/grpc v1.33.2
)
