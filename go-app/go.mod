module example.com/main

go 1.18

require (
	github.com/justinas/alice v1.2.0
	local.pkg/middleware v0.0.0-00010101000000-000000000000
	local.pkg/sample v0.0.0-00010101000000-000000000000
)

require (
	github.com/cilium/ebpf v0.9.1 // indirect
	github.com/cosiner/argv v0.1.0 // indirect
	github.com/cpuguy83/go-md2man/v2 v2.0.2 // indirect
	github.com/derekparker/trie v0.0.0-20200317170641-1fdf38b7b0e9 // indirect
	github.com/go-delve/delve v1.9.0 // indirect
	github.com/go-delve/liner v1.2.3-0.20220127212407-d32d89dd2a5d // indirect
	github.com/google/go-dap v0.6.0 // indirect
	github.com/hashicorp/golang-lru v0.5.4 // indirect
	github.com/inconshreveable/mousetrap v1.0.0 // indirect
	github.com/konsorten/go-windows-terminal-sequences v1.0.3 // indirect
	github.com/mattn/go-colorable v0.1.12 // indirect
	github.com/mattn/go-isatty v0.0.14 // indirect
	github.com/mattn/go-runewidth v0.0.13 // indirect
	github.com/rivo/uniseg v0.3.1 // indirect
	github.com/russross/blackfriday/v2 v2.1.0 // indirect
	github.com/shurcooL/sanitized_anchor_name v1.0.0 // indirect
	github.com/sirupsen/logrus v1.9.0 // indirect
	github.com/spf13/cobra v1.5.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	go.starlark.net v0.0.0-20220714194419-4cadf0a12139 // indirect
	golang.org/x/arch v0.0.0-20220722155209-00200b7164a7 // indirect
	golang.org/x/sys v0.0.0-20220731174439-a90be440212d // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	local.pkg/api v0.0.0-00010101000000-000000000000 // indirect
)

replace local.pkg/api => ./api

replace local.pkg/middleware => ./middleware

replace local.pkg/sample => ./views
