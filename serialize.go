package exception

// Serialize 序列化实现，可以在例外信息补充中进行序列化显示
type Serialize interface {
	// Serialize 序列化
	Serialize() string
}
