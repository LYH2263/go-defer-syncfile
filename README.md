# go-defer-syncfile

defer Sync 错误被吞

internal/save/atomic.go 的 WriteFile：defer 中 f.Sync() 错误未赋给命名返回值
