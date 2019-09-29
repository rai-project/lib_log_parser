//go:generate go get -v github.com/jteeuwen/go-bindata/...
//go:generate go get -v github.com/elazarl/go-bindata-assetfs/...
//go:generate go-bindata -nomemcopy -prefix assets/ -pkg cudnn_log_parser -o assets_static.go -ignore=.DS_Store -ignore=README.md assets/...

package cudnn_log_parser