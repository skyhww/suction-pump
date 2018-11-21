package version

const (
	ON = iota
	OFF
)

//每个模块应可以独立的装载，卸载已经重新加载
type Module interface {
	reload() error
	install() error
	unInstall() error
	getStatus() int
	//提供更细粒度的锁
	Lock()
	unLock()
}
