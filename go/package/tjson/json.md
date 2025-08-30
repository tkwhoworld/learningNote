## 概述
json包实现了json对象的编解码，参见RFC 4627。Json对象和go类型的映射关系请参见Marshal和Unmarshal函数的文档。

### 使用方式
import "encoding/json"

#### type Number 介绍
概述:
type Number 这个是"encoding/json"包下的一个类型。此number的定义为:
type Number string //就是扩展了string这个类型。给这个类型添加了几个方法
此类型提供的方法
func (Number) Int64
func (n Number) Int64() (int64, error)
将该数字作为int64类型返回。

func (Number) Float64()
func (n Number) Float64() (float64, error)
将该数字作为float64类型返回。
func (n Number) String() string
返回该数字的字面值文本表示。

#### type RawMesage
概述:
RawMessage类型是一个保持原本编码的json对象。本类型实现了Marshaler和Unmarshaler接口，用于延迟json的解码或者预计算json的编码。
定义:type RawMesage []byte

func (*RawMessage) MarshalJSON
func (m *RawMessage) MarshalJSON() ([]byte, error)

func (*RawMessage) UnmarshalJSON
func (m *RawMessage) UnmarshalJSON(data []byte) error

#### type Marshaler
这个是"encoding/json"包下的一个接口类型。接口类型定义了一个方法
type Marshaler interface {
    MarshalJSON() ([]byte, error)
}
实现了Marshaler接口的类型可以将自身序列化为合法的json描述。


#### type Unmarshaler
这个是"encoding/json"包下的一个接口类型。接口类型定义了一个方法
type Unmarshaler interface {
    UnmarshalJSON([]byte) error
}
实现了Unmarshaler接口的对象可以将自身的json描述反序列化。该方法可以认为输入是合法的json字符串。如果要在方法返回后保存自身的json数据，必须进行拷贝。

