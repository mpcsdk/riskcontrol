// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/shopspring/decimal"
)

// Contractrule is the golang structure for table contractrule.
type Contractrule struct {
	Id               int             `json:"id"               ` // ID
	CreateTime       *gtime.Time     `json:"createTime"       ` // 创建时间
	UpdateTime       *gtime.Time     `json:"updateTime"       ` // 更新时间
	SceneNo          string          `json:"sceneNo"          ` // 场景号
	ContractAddress  string          `json:"contractAddress"  ` // 合约地址
	ContractName     string          `json:"contractName"     ` // 合约名
	MethodName       string          `json:"methodName"       ` // 方法名
	MethodSignature  string          `json:"methodSignature"  ` // 方法签名
	MethodFromField  string          `json:"methodFromField"  ` // 方法from字段名
	MethodToField    string          `json:"methodToField"    ` // 方法to字段名
	MethodValueField string          `json:"methodValueField" ` // 方法value字段名
	EventName        string          `json:"eventName"        ` // 事件名
	EventSignature   string          `json:"eventSignature"   ` // 事件签名
	EventFromField   string          `json:"eventFromField"   ` // 事件from字段名
	EventToField     string          `json:"eventToField"     ` // 事件to字段名
	EventValueField  string          `json:"eventValueField"  ` // 事件value字段名
	ControlOperation string          `json:"controlOperation" ` // 控制操作
	Threshold        decimal.Decimal `json:"threshold"        ` // 阈值
	ContractKind     string          `json:"contractKind"     ` // 合约类型
	EventTopic       string          `json:"eventTopic"       ` // 事件topic
	WhiteAddrList    string          `json:"whiteAddrList"    ` // to地址白名单
	ChainId          string          `json:"chainId"          ` // 链id
}
