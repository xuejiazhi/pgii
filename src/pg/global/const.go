package global

const (
	ZeroCMDLength = iota
	OneCMDLength
	TwoCMDLength
	ThreeCMDLength
	FourCMDLength
	FiveCMDLength
)

const (
	NoneStyle = iota
	DatabaseStyle
	TableStyle
	IndexStyle
	ViewStyle
	SelectStyle
	SchemaStyle
	TriggerStyle
	VersionStyle
	ConnectionStyle
	ProcessStyle
	TableSpaceStyle //表空间
)

const (
	MaxConnections               = iota //最大连接数
	SuperuserReservedConnections        //超级用户保留的连接数
	RemainingConnections                //剩余连接数
	InUseConnections                    //正在使用的链接数
)

const (
	DDL = iota
	DUMP
)
