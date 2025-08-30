# net包学习记录
### Listen函数 
#### 定义
func Listen(net, laddr string) (Listener, error)
#### 理解
这是一个net包下的函数，可以直接通过net.Listen调用
传的参数为2个字符串，一个是协议类型，一个地址，返回值为：Listener类型和error

### Listener type
#### 定义
type Listener interface {
    // Addr返回该接口的网络地址
    Addr() Addr
    // Accept等待并返回下一个连接到该接口的连接
    Accept() (c Conn, err error)
    // Close关闭该接口，并使任何阻塞的Accept操作都会不再阻塞并返回错误。
    Close() error
}
#### 理解
这是一个接口的类型，接口的名称为:Listener，接口包含的方法有三个Addr() Accept() Close()

#### 思考 
net.Listen返回了一个net.Listener类型,但是net.Listener是一个接口，具体的实现是TCPListener跟UnixListener,
那么返回的Listener到底是那个呢

### Dial函数
#### 定义
func Dial(net,address string) (Conn, error)
net，跟address为string类型的字符串，net主要有tcp tcp4 tcp6 udp udp4
udp6 ip ip4 ip6 unix unixgram unixpacket.
address为网络地址，对于TCP和UDP网络，地址格式为:host:port
返回值为:Conn和error2个类型的返回值，Conn的是一个接口类型，具体介绍见Conn

#### 理解
这是一个net包的函数，使用方式为：net.Dial('tcp','127.0.0.1:8000')

### Conn type
#### 定义
type Conn interface {
    //Read从连接中读取数据
    //Read方法可能会在超过某个固定时间限制后超时返回错误，该错误的Timeout()方法返回真
    Read(b []byte)(n int, err error)
    //Write从连接中写数据
    //Write方法可能会在超过某个固定时间限制时返回错误，该错误的Timeout()方法返回真
    Write(b []byte)(n int, err error)
    //Close方法关闭该连接
    //并会导致任何阻塞中的Read或Write方法不再阻塞并返回错误
    Close() error
    LocalAddr() Addr
    // 返回远端网络地址
    RemoteAddr() Addr
    // 设定该连接的读写deadline，等价于同时调用SetReadDeadline和SetWriteDeadline
    // deadline是一个绝对时间，超过该时间后I/O操作就会直接因超时失败返回而不会阻塞
    // deadline对之后的所有I/O操作都起效，而不仅仅是下一次的读或写操作
    // 参数t为零值表示不设置期限
    SetDeadline(t time.Time) error
    // 设定该连接的读操作deadline，参数t为零值表示不设置期限
    SetReadDeadline(t time.Time) error
    // 设定该连接的写操作deadline，参数t为零值表示不设置期限
    // 即使写入超时，返回值n也可能>0，说明成功写入了部分数据
    SetWriteDeadline(t time.Time) error
}