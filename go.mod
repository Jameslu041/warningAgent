module warningAgent

go 1.12

require (
	github.com/Sirupsen/logrus v1.4.0
	github.com/fsnotify/fsnotify v1.4.7
	github.com/golang/glog v0.0.0-20160126235308-23def4e6c14b
	github.com/hyperchain/gosdk v1.2.12
	github.com/mindprince/gonvml v0.0.0-20190828220739-9ebdce4bb989
	github.com/op/go-logging v0.0.0-20160315200505-970db520ece7 // indirect
	github.com/open-falcon/falcon-plus v0.2.2
	github.com/pelletier/go-toml v1.4.0 // indirect
	github.com/pkg/errors v0.8.1 // indirect
	github.com/spf13/afero v1.2.2 // indirect
	github.com/spf13/jwalterweatherman v1.1.0 // indirect
	github.com/spf13/viper v1.6.2
	github.com/stretchr/testify v1.4.0 // indirect
	github.com/terasum/viper v1.0.0 // indirect
	github.com/toolkits/file v0.0.0-20160325033739-a5b3c5147e07
	github.com/toolkits/net v0.0.0-20160910085801-3f39ab6fe3ce // indirect
	github.com/toolkits/nux v0.0.0-20191107142017-8ddcb501004c
	github.com/toolkits/slice v0.0.0-20141116085117-e44a80af2484
	github.com/toolkits/sys v0.0.0-20170615103026-1f33b217ffaf
	github.com/toolkits/time v0.0.0-20160524122720-c274716e8d7f // indirect
	golang.org/x/sys v0.0.0-20190813064441-fde4db37ae7a // indirect
)

replace github.com/hyperchain/gosdk v1.2.12 => git.hyperchain.cn/hyperchain/gosdk.git v1.2.12

replace github.com/mholt/archiver v3.1.0+incompatible => github.com/mholt/archiver v2.0.0+incompatible

replace github.com/Sirupsen/logrus v1.4.0 => github.com/sirupsen/logrus v1.3.0
