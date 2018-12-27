//go:generate protoc -I . -I $GOPATH/src -I ../../../../vendor/github.com/gogo/protobuf/ --gogofast_out=import_path=github.com/docker/docker/api/types/swarm/runtime:. plugin.proto

package runtime // import "github.com/docker/docker/api/types/swarm/runtime"
