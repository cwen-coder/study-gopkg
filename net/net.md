#package net

import "net"

net包提供了可移植的网络I/O接口，包括TCP/IP、UDP、域名解析和Unix域socket。

虽然本包提供了对网络原语的访问，大部分使用者只需要Dial、Listen和Accept函数提供的基本接口；以及相关的Conn和Listener接口。crypto/tls包提供了相同的接口和类似的Dial和Listen函数。

Dial函数和服务端建立连接：
```go
conn, err := net.Dial("tcp", "google.com:80")
if err != nil {
    // handle error
}
fmt.Fprintf(conn, "GET / HTTP/1.0\r\n\r\n")
status, err := bufio.NewReader(conn).ReadString('\n')
// ...
```
Listen函数创建的服务端：
```go
ln, err := net.Listen("tcp", ":8080")
if err != nil {
    // handle error
}
for {
    conn, err := ln.Accept()
    if err != nil {
        // handle error
        continue
    }
    go handleConnection(conn)
}
```
###Index
+ Constants
+ Variables  
+ type ParseError  
    + func (e *ParseError) Error() string  
+ type Error    
+ type InvalidAddrError 
    + func (e InvalidAddrError) Error() string
    + func (e InvalidAddrError) Temporary() bool
    + func (e InvalidAddrError) Timeout() bool
+ type UnknownNetworkError  
    + func (e UnknownNetworkError) Error() string
    + func (e UnknownNetworkError) Temporary() bool
    + func (e UnknownNetworkError) Timeout() bool
+ type DNSConfigError
    + func (e *DNSConfigError) Error() string
    + func (e *DNSConfigError) Temporary() bool
    + func (e *DNSConfigError) Timeout() bool
+ type DNSError
    + func (e *DNSError) Error() string
    + func (e *DNSError) Temporary() bool
    + func (e *DNSError) Timeout() bool
+ type AddrError
    + func (e *AddrError) Error() string
    + func (e *AddrError) Temporary() bool
    + func (e *AddrError) Timeout() bool
+ type OpError
    + func (e *OpError) Error() string
    + func (e *OpError) Temporary() bool
    + func (e *OpError) Timeout() bool
+ func SplitHostPort(hostport string) (host, port string, err error)
+ func JoinHostPort(host, port string) string
+ type HardwareAddr
    + func ParseMAC(s string) (hw HardwareAddr, err error)
    + func (a HardwareAddr) String() string
+ type Flags
    + func (f Flags) String() string
+ type Interface
    + func InterfaceByIndex(index int) (*Interface, error)
    + func InterfaceByName(name string) (*Interface, error)
    + func (ifi *Interface) Addrs() ([]Addr, error)
    + func (ifi *Interface) MulticastAddrs() ([]Addr, error)
+ func Interfaces() ([]Interface, error)
+ func InterfaceAddrs() ([]Addr, error)
+ type IP
    + func IPv4(a, b, c, d byte) IP
    + func ParseIP(s string) IP
    + func (ip IP) IsGlobalUnicast() bool
    + func (ip IP) IsLinkLocalUnicast() bool
    + func (ip IP) IsInterfaceLocalMulticast() bool
    + func (ip IP) IsLinkLocalMulticast() bool
    + func (ip IP) IsMulticast() bool
    + func (ip IP) IsLoopback() bool
    + func (ip IP) IsUnspecified() bool
    + func (ip IP) DefaultMask() IPMask
    + func (ip IP) Equal(x IP) bool
    + func (ip IP) To16() IP
    + func (ip IP) To4() IP
    + func (ip IP) Mask(mask IPMask) IP
    + func (ip IP) String() string
    + func (ip IP) MarshalText() ([]byte, error)
    + func (ip *IP) UnmarshalText(text []byte) error
+ type IPMask
    + func IPv4Mask(a, b, c, d byte) IPMask
    + func CIDRMask(ones, bits int) IPMask
    + func (m IPMask) Size() (ones, bits int)
    + func (m IPMask) String() string
+ type IPNet
    + func ParseCIDR(s string) (IP, *IPNet, error)
    + func (n *IPNet) Contains(ip IP) bool
    + func (n *IPNet) Network() string
    + func (n *IPNet) String() string
+ type Addr
+ type Conn
    + func Dial(network, address string) (Conn, error)
    + func DialTimeout(network, address string, timeout time.Duration) (Conn, error)
    + func Pipe() (Conn, Conn)
+ type PacketConn
    + func ListenPacket(net, laddr string) (PacketConn, error)
+ type Dialer
    + func (d *Dialer) Dial(network, address string) (Conn, error)
+ type Listener
    + func Listen(net, laddr string) (Listener, error)
+ type IPAddr
    + func ResolveIPAddr(net, addr string) (*IPAddr, error)
    + func (a *IPAddr) Network() string
    + func (a *IPAddr) String() string
+ type TCPAddr
    + func ResolveTCPAddr(net, addr string) (*TCPAddr, error)
    + func (a *TCPAddr) Network() string
    + func (a *TCPAddr) String() string
+  type UDPAddr
    + func ResolveUDPAddr(net, addr string) (*UDPAddr, error)
    + func (a *UDPAddr) Network() string
    + func (a *UDPAddr) String() string
+ type UnixAddr
    + func ResolveUnixAddr(net, addr string) (*UnixAddr, error)
    + func (a *UnixAddr) Network() string
    + func (a *UnixAddr) String() string
+ type IPConn
    + func DialIP(netProto string, laddr, raddr *IPAddr) (*IPConn, error)
    + func ListenIP(netProto string, laddr *IPAddr) (*IPConn, error)
    + func (c *IPConn) LocalAddr() Addr
    + func (c *IPConn) RemoteAddr() Addr
    + func (c *IPConn) SetReadBuffer(bytes int) error
    + func (c *IPConn) SetWriteBuffer(bytes int) error
    + func (c *IPConn) SetDeadline(t time.Time) error
    + func (c *IPConn) SetReadDeadline(t time.Time) error
    + func (c *IPConn) SetWriteDeadline(t time.Time) error
    + func (c *IPConn) Read(b []byte) (int, error)
    + func (c *IPConn) ReadFrom(b []byte) (int, Addr, error)
    + func (c *IPConn) ReadFromIP(b []byte) (int, *IPAddr, error)
    + func (c *IPConn) ReadMsgIP(b, oob []byte) (n, oobn, flags int, addr *IPAddr, err error)
    + func (c *IPConn) Write(b []byte) (int, error)
    + func (c *IPConn) WriteTo(b []byte, addr Addr) (int, error)
    + func (c *IPConn) WriteToIP(b []byte, addr *IPAddr) (int, error)
    + func (c *IPConn) WriteMsgIP(b, oob []byte, addr *IPAddr) (n, oobn int, err error)
    + func (c *IPConn) Close() error
    + func (c *IPConn) File() (f *os.File, err error)
+ type TCPConn
    + func DialTCP(net string, laddr, raddr *TCPAddr) (*TCPConn, error)
    + func (c *TCPConn) LocalAddr() Addr
    + func (c *TCPConn) RemoteAddr() Addr
    + func (c *TCPConn) SetReadBuffer(bytes int) error
    + func (c *TCPConn) SetWriteBuffer(bytes int) error
    + func (c *TCPConn) SetDeadline(t time.Time) error
    + func (c *TCPConn) SetReadDeadline(t time.Time) error
    + func (c *TCPConn) SetWriteDeadline(t time.Time) error
    + func (c *TCPConn) SetKeepAlive(keepalive bool) error
    + func (c *TCPConn) SetKeepAlivePeriod(d time.Duration) error
    + func (c *TCPConn) SetLinger(sec int) error
    + func (c *TCPConn) SetNoDelay(noDelay bool) error
    + func (c *TCPConn) Read(b []byte) (int, error)
    + func (c *TCPConn) ReadFrom(r io.Reader) (int64, error)
    + func (c *TCPConn) Write(b []byte) (int, error)
    + func (c *TCPConn) Close() error
    + func (c *TCPConn) CloseRead() error
    + func (c *TCPConn) CloseWrite() error
    + func (c *TCPConn) File() (f *os.File, err error)
+   + type UDPConn
    + func DialUDP(net string, laddr, raddr *UDPAddr) (*UDPConn, error)
    + func ListenMulticastUDP(net string, ifi *Interface, gaddr *UDPAddr) (*UDPConn, error)
    + func ListenUDP(net string, laddr *UDPAddr) (*UDPConn, error)
    + func (c *UDPConn) LocalAddr() Addr
    + func (c *UDPConn) RemoteAddr() Addr
    + func (c *UDPConn) SetReadBuffer(bytes int) error
    + func (c *UDPConn) SetWriteBuffer(bytes int) error
    + func (c *UDPConn) SetDeadline(t time.Time) error
    + func (c *UDPConn) SetReadDeadline(t time.Time) error
    + func (c *UDPConn) SetWriteDeadline(t time.Time) error
    + func (c *UDPConn) Read(b []byte) (int, error)
    + func (c *UDPConn) ReadFrom(b []byte) (int, Addr, error)   + 
    + func (c *UDPConn) ReadFromUDP(b []byte) (n int, addr *UDPAddr, err error)
    + func (c *UDPConn) ReadMsgUDP(b, oob []byte) (n, oobn, flags int, addr *UDPAddr, err error)
    + func (c *UDPConn) Write(b []byte) (int, error)
    + func (c *UDPConn) WriteTo(b []byte, addr Addr) (int, error)
    + func (c *UDPConn) WriteToUDP(b []byte, addr *UDPAddr) (int, error)
    + func (c *UDPConn) WriteMsgUDP(b, oob []byte, addr *UDPAddr) (n, oobn int, err error)
    + func (c *UDPConn) Close() error
    + func (c *UDPConn) File() (f *os.File, err error)
+ type UnixConn
    + func DialUnix(net string, laddr, raddr *UnixAddr) (*UnixConn, error)
    + func ListenUnixgram(net string, laddr *UnixAddr) (*UnixConn, error)
    + func (c *UnixConn) LocalAddr() Addr
    + func (c *UnixConn) RemoteAddr() Addr
    + func (c *UnixConn) SetReadBuffer(bytes int) error
    + func (c *UnixConn) SetWriteBuffer(bytes int) error
    + func (c *UnixConn) SetDeadline(t time.Time) error
    + func (c *UnixConn) SetReadDeadline(t time.Time) error
    + func (c *UnixConn) SetWriteDeadline(t time.Time) error
    + func (c *UnixConn) Read(b []byte) (int, error)
    + func (c *UnixConn) ReadFrom(b []byte) (int, Addr, error)
    + func (c *UnixConn) ReadFromUnix(b []byte) (n int, addr *UnixAddr, err error)
    + func (c *UnixConn)    + ReadMsgUnix(b, oob []byte) (n, oobn, flags int, addr *UnixAddr, err error)
    + func (c *UnixConn) Write(b []byte) (int, error)
    + func (c *UnixConn) WriteTo(b []byte, addr Addr) (n int, err error)
    + func (c *UnixConn) WriteToUnix(b []byte, addr *UnixAddr) (n int, err error)
    + func (c *UnixConn) WriteMsgUnix(b, oob []byte, addr *UnixAddr) (n, oobn int, err error)
    + func (c *UnixConn) Close() error
    + func (c *UnixConn) CloseRead() error
    + func (c *UnixConn) CloseWrite() error
    + func (c *UnixConn) File() (f *os.File, err error)
+ type TCPListener
    + func ListenTCP(net string, laddr *TCPAddr) (*TCPListener, error)
    + func (l *TCPListener) Addr() Addr
    + func (l *TCPListener) SetDeadline(t time.Time) error
    + func (l *TCPListener) Accept() (Conn, error)
    + func (l *TCPListener) AcceptTCP() (*TCPConn, error)
    + func (l *TCPListener) Close() error
    + func (l *TCPListener) File() (f *os.File, err error)
+ type UnixListener
    + func ListenUnix(net string, laddr *UnixAddr) (*UnixListener, error)
    + func (l *UnixListener) Addr() Addr
    + func (l *UnixListener) SetDeadline(t time.Time) (err error)
    + func (l *UnixListener) Accept() (c Conn, err error)
    + func (l *UnixListener) AcceptUnix() (*UnixConn, error)
   +  + func (l *UnixListener) Close() error
    + func (l *UnixListener) File() (f *os.File, err error)
    + func FileConn(f *os.File) (c Conn, err error)
    + func FilePacketConn(f *os.File) (c PacketConn, err error)
    + func FileListener(f *os.File) (l Listener, err error)
+ type MX
+ type NS
+ type SRV
+ func LookupPort(network, service string) (port int, err error)
+ func LookupCNAME(name string) (cname string, err error)
+ func LookupHost(host string) (addrs []string, err error)
+ func LookupIP(host string) (addrs []IP, err error)
+ func LookupAddr(addr string) (name []string, err error)
+ func LookupMX(name string) (mx []*MX, err error)
+ func LookupNS(name string) (ns []*NS, err error)
+ func LookupSRV(service, proto, name string) (cname string, addrs []*SRV, err error)
+ func LookupTXT(name string) (txt []string, err error)

#### Examples

+ Listener
#### Constants
```GO
const ( 
    IPv4len = 4 
    IPv6len = 16
)
```
IP address lengths (bytes).

#### Variables
```GO
var ( 
    IPv4bcast = IPv4(255, 255, 255, 255) // 广播地址 
    IPv4allsys = IPv4(224, 0, 0, 1) // 所有主机和路由器 
    IPv4allrouter = IPv4(224, 0, 0, 2) // 所有路由器 
    IPv4zero = IPv4(0, 0, 0, 0) // 本地地址，只能作为源地址（曾用作广播地址） )
```
常用的IPv4地址。

```go
var ( 
    IPv6zero = IP{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
    IPv6unspecified = IP{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0} 
    IPv6loopback = IP{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1} 
    IPv6interfacelocalallnodes = IP{0xff, 0x01, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0x01} 
    IPv6linklocalallnodes = IP{0xff, 0x02, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0x01} 
    IPv6linklocalallrouters = IP{0xff, 0x02, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0x02}
)
```
常用的IPv6地址。
```go
var ( 
    ErrWriteToConnected = errors.New("use of WriteTo with pre-connected connection")
)
```
很多OpError类型的错误会包含本错误。

##### type ParseError
```go
type ParseError struct { 
    Type string Text string 
}
```
ParseError代表一个格式错误的字符串，Type为期望的格式。


##### func (*ParseError) Error
```go
func (e *ParseError) Error() string
```
##### type Error
```go
type Error interface { 
    error Timeout() bool // 错误是否为超时？ 
    Temporary() bool // 错误是否是临时的？ 
}
```
Error代表一个网络错误。

##### type UnknownNetworkError
```go
type UnknownNetworkError string
```
##### func (UnknownNetworkError) Error
```go
func (e UnknownNetworkError) Error() string
```
##### func (UnknownNetworkError) Temporary
```go
func (e UnknownNetworkError) Temporary() bool
```
##### func (UnknownNetworkError) Timeout
```go
func (e UnknownNetworkError) Timeout() bool
```
##### type InvalidAddrError
```go
type InvalidAddrError string
```
##### func (InvalidAddrError) Error
```go
func (e InvalidAddrError) Error() string
```
##### func (InvalidAddrError) Temporary
```go
func (e InvalidAddrError) Temporary() bool
```
##### func (InvalidAddrError) Timeout
```go
func (e InvalidAddrError) Timeout() bool
```
##### type DNSConfigError
```go
type DNSConfigError struct { Err error }
```
DNSConfigError代表读取主机DNS配置时出现的错误。

##### func (*DNSConfigError) Error
```go
func (e *DNSConfigError) Error() string
```
##### func (*DNSConfigError) Temporary
```go
func (e *DNSConfigError) Temporary() bool
```
##### func (*DNSConfigError) Timeout
```go
func (e *DNSConfigError) Timeout() bool
```
##### type DNSError
```go
type DNSError struct { 
    Err string // 错误的描述 
    Name string // 查询的名称 
    Server string // 使用的服务器 
    IsTimeout bool 
}
```
DNSError代表DNS查询的错误。

##### func (*DNSError) Error
```go
func (e *DNSError) Error() string
```
##### func (*DNSError) Temporary
```go
func (e *DNSError) Temporary() bool
```
##### func (*DNSError) Timeout
```go
func (e *DNSError) Timeout() bool
```
##### type AddrError
```go
type AddrError struct { Err string Addr string }
```
##### func (*AddrError) Error
```go
func (e *AddrError) Error() string
```
##### func (*AddrError) Temporary
```go
func (e *AddrError) Temporary() bool
```
##### func (*AddrError) Timeout
```go
func (e *AddrError) Timeout() bool
```
##### type OpError
```go
type OpError struct { 
    // Op是出现错误的操作，如"read"或"write" Op string 
    // Net是错误所在的网络类型，如"tcp"或"udp6" Net string 
    // Addr是出现错误的网络地址 Addr Addr 
    // Err是操作中出现的错误 Err error }
```
OpError是经常被net包的函数返回的错误类型。它描述了该错误的操作、网络类型和网络地址。

##### func (*OpError) Error
```go
func (e *OpError) Error() string
```
##### func (*OpError) Temporary
```go
func (e *OpError) Temporary() bool
```
##### func (*OpError) Timeout
```go
func (e *OpError) Timeout() bool
```
##### func SplitHostPort
```go
func SplitHostPort(hostport string) (host, port string, err error)
```
函数将格式为"host:port"、"[host]:port"或"[ipv6-host%zone]:port"的网络地址分割为host或ipv6-host%zone和port两个部分。Ipv6的文字地址或者主机名必须用方括号括起来，如"[::1]:80"、"[ipv6-host]:http"、"[ipv6-host%zone]:80"。

##### func JoinHostPort
```go
func JoinHostPort(host, port string) string
```
函数将host和port合并为一个网络地址。一般格式为"host:port"；如果host含有冒号或百分号，格式为"[host]:port"。

##### type HardwareAddr
```go
type HardwareAddr []byte
```
HardwareAddr类型代表一个硬件地址（MAC地址）。

##### func ParseMAC
```go
func ParseMAC(s string) (hw HardwareAddr, err error)
```
ParseMAC函数使用如下格式解析一个IEEE 802 MAC-48、EUI-48或EUI-64硬件地址：
```go
01:23:45:67:89:ab
01:23:45:67:89:ab:cd:ef
01-23-45-67-89-ab
01-23-45-67-89-ab-cd-ef
0123.4567.89ab
0123.4567.89ab.cdef
```
##### func (HardwareAddr) String
```go
func (a HardwareAddr) String() string
```
##### type Flags
```go
type Flags uint
const ( 
    FlagUp Flags = 1 << iota // 接口在活动状态 
    FlagBroadcast // 接口支持广播 
    FlagLoopback // 接口是环回的 
    FlagPointToPoint // 接口是点对点的 
    FlagMulticast // 接口支持组播 
)
```
##### func (Flags) String
```go
func (f Flags) String() string
```
##### type Interface
```go
type Interface struct { 
    Index int // 索引，>=1的整数 
    MTU int // 最大传输单元 
    Name string // 接口名，例如"en0"、"lo0"、"eth0.100" 
    HardwareAddr HardwareAddr // 硬件地址，IEEE MAC-48、EUI-48或EUI-64格式 
    Flags Flags // 接口的属性，例如FlagUp、FlagLoopback、FlagMulticast 
}
```
Interface类型代表一个网络接口（系统与网络的一个接点）。包含接口索引到名字的映射，也包含接口的设备信息。

##### func InterfaceByIndex
```go
func InterfaceByIndex(index int) (*Interface, error)
```
InterfaceByIndex返回指定索引的网络接口。

##### func InterfaceByName
```go
func InterfaceByName(name string) (*Interface, error)
```
InterfaceByName返回指定名字的网络接口。

##### func (*Interface) Addrs
```go
func (ifi *Interface) Addrs() ([]Addr, error)
```
Addrs方法返回网络接口ifi的一或多个接口地址。

##### func (*Interface) MulticastAddrs
```go
func (ifi *Interface) MulticastAddrs() ([]Addr, error)
```
MulticastAddrs返回网络接口ifi加入的多播组地址。

##### func Interfaces
```go
func Interfaces() ([]Interface, error)
```
Interfaces返回该系统的网络接口列表。

##### func InterfaceAddrs
```go
func InterfaceAddrs() ([]Addr, error)
```
InterfaceAddrs返回该系统的网络接口的地址列表。

##### type IP
```go
type IP []byte
```
IP类型是代表单个IP地址的[]byte切片。本包的函数都可以接受4字节（IPv4）和16字节（IPv6）的切片作为输入。

注意，IP地址是IPv4地址还是IPv6地址是语义上的属性，而不取决于切片的长度：16字节的切片也可以是IPv4地址。

##### func IPv4
```go
func IPv4(a, b, c, d byte) IP
```
IPv4返回包含一个IPv4地址a.b.c.d的IP地址（16字节格式）。

##### func ParseIP
```go
func ParseIP(s string) IP
```
ParseIP将s解析为IP地址，并返回该地址。如果s不是合法的IP地址文本表示，ParseIP会返回nil。

字符串可以是小数点分隔的IPv4格式（如"74.125.19.99"）或IPv6格式（如"2001:4860:0:2001::68"）格式。

##### func (IP) IsGlobalUnicast
```go
func (ip IP) IsGlobalUnicast() bool
```
如果ip是全局单播地址，则返回真。

##### func (IP) IsLinkLocalUnicast
```go
func (ip IP) IsLinkLocalUnicast() bool
```
如果ip是链路本地单播地址，则返回真。

##### func (IP) IsInterfaceLocalMulticast
```go
func (ip IP) IsInterfaceLocalMulticast() bool
```
如果ip是接口本地组播地址，则返回真。

##### func (IP) IsLinkLocalMulticast
```go
func (ip IP) IsLinkLocalMulticast() bool
```
如果ip是链路本地组播地址，则返回真。

##### func (IP) IsMulticast
```go
func (ip IP) IsMulticast() bool
```
如果ip是组播地址，则返回真。

##### func (IP) IsLoopback
```
func (ip IP) IsLoopback() bool
```
如果ip是环回地址，则返回真。

#####func (IP) IsUnspecified
```go
func (ip IP) IsUnspecified() bool
```
如果ip是未指定地址，则返回真。

##### func (IP) DefaultMask
```go
func (ip IP) DefaultMask() IPMask
```
函数返回IP地址ip的默认子网掩码。只有IPv4有默认子网掩码；如果ip不是合法的IPv4地址，会返回nil。

##### func (IP) Equal
```go
func (ip IP) Equal(x IP) bool
```
如果ip和x代表同一个IP地址，Equal会返回真。代表同一地址的IPv4地址和IPv6地址也被认为是相等的。

##### func (IP) To16
```go
func (ip IP) To16() IP
```
To16将一个IP地址转换为16字节表示。如果ip不是一个IP地址（长度错误），To16会返回nil。

##### func (IP) To4
```go
func (ip IP) To4() IP
```
To4将一个IPv4地址转换为4字节表示。如果ip不是IPv4地址，To4会返回nil。

#####func (IP) Mask
```go
func (ip IP) Mask(mask IPMask) IP
```
Mask方法认为mask为ip的子网掩码，返回ip的网络地址部分的ip。（主机地址部分都置0）

##### func (IP) String
```go
func (ip IP) String() string
```
String返回IP地址ip的字符串表示。如果ip是IPv4地址，返回值的格式为点分隔的，如"74.125.19.99"；否则表示为IPv6格式，如"2001:4860:0:2001::68"。

##### func (IP) MarshalText
```go
func (ip IP) MarshalText() ([]byte, error)
```
MarshalText实现了encoding.TextMarshaler接口，返回值和String方法一样。

##### func (*IP) UnmarshalText
```go
func (ip *IP) UnmarshalText(text []byte) error
```
UnmarshalText实现了encoding.TextUnmarshaler接口。IP地址字符串应该是ParseIP函数可以接受的格式。

##### type IPMask
```go
type IPMask []byte
```
IPMask代表一个IP地址的掩码。

##### func IPv4Mask
```go
func IPv4Mask(a, b, c, d byte) IPMask
```
IPv4Mask返回一个4字节格式的IPv4掩码a.b.c.d。

##### func CIDRMask
```go
func CIDRMask(ones, bits int) IPMask
```
CIDRMask返回一个IPMask类型值，该返回值总共有bits个字位，其中前ones个字位都是1，其余字位都是0。

##### func (IPMask) Size
```go
func (m IPMask) Size() (ones, bits int)
```
Size返回m的前导的1字位数和总字位数。如果m不是规范的子网掩码（字位：/^1+0+$/），将返会(0, 0)。

##### func (IPMask) String
```go
func (m IPMask) String() string
```
String返回m的十六进制格式，没有标点。

##### type IPNet
```go
type IPNet struct { 
    IP IP // 网络地址 
    Mask IPMask // 子网掩码 
}
```
IPNet表示一个IP网络。

##### func ParseCIDR
```go
func ParseCIDR(s string) (IP, *IPNet, error)
```
ParseCIDR将s作为一个CIDR（无类型域间路由）的IP地址和掩码字符串，如"192.168.100.1/24"或"2001:DB8::/48"，解析并返回IP地址和IP网络，参见RFC 4632和RFC 4291。

本函数会返回IP地址和该IP所在的网络和掩码。例如，ParseCIDR("192.168.100.1/16")会返回IP地址192.168.100.1和IP网络192.168.0.0/16。

##### func (*IPNet) Contains
```go
func (n *IPNet) Contains(ip IP) bool
```
Contains报告该网络是否包含地址ip。

##### func (*IPNet) Network
```go
func (n *IPNet) Network() string
```
Network返回网络类型名："ip+net"，注意该类型名是不合法的。

##### func (*IPNet) String
```go
func (n *IPNet) String() string
```
String返回n的CIDR表示，如"192.168.100.1/24"或"2001:DB8::/48"，参见RFC 4632和RFC 4291。如果n的Mask字段不是规范格式，它会返回一个包含n.IP.String()、斜线、n.Mask.String()（此时表示为无标点十六进制格式）的字符串，如"192.168.100.1/c000ff00"。

##### type Addr
```go 
type Addr interface { 
    Network() string // 网络名 
    String() string // 字符串格式的地址 
}
```
Addr代表一个网络终端地址。

##### type Conn
```go
type Conn interface { 
    // Read从连接中读取数据 
    // Read方法可能会在超过某个固定时间限制后超时返回错误，该错误的Timeout()方法返回真 
    Read(b []byte) (n int, err error) 
    // Write从连接中写入数据 // Write方法可能会在超过某个固定时间限制后超时返回错误，该错误的Timeout()方法返回真 
    Write(b []byte) (n int, err error) 
    // Close方法关闭该连接 
    // 并会导致任何阻塞中的Read或Write方法不再阻塞并返回错误 
    Close() error 
    // 返回本地网络地址 
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
```
Conn接口代表通用的面向流的网络连接。多个线程可能会同时调用同一个Conn的方法。

##### func Dial
```go
func Dial(network, address string) (Conn, error)
```
在网络network上连接地址address，并返回一个Conn接口。可用的网络类型有：

"tcp"、"tcp4"、"tcp6"、"udp"、"udp4"、"udp6"、"ip"、"ip4"、"ip6"、"unix"、"unixgram"、"unixpacket"

对TCP和UDP网络，地址格式是host:port或[host]:port，参见函数JoinHostPort和SplitHostPort。
```go
Dial("tcp", "12.34.56.78:80")
Dial("tcp", "google.com:http")
Dial("tcp", "[2001:db8::1]:http")
Dial("tcp", "[fe80::1%lo0]:80")
```
对IP网络，network必须是"ip"、"ip4"、"ip6"后跟冒号和协议号或者协议名，地址必须是IP地址字面值。
```go
Dial("ip4:1", "127.0.0.1")
Dial("ip6:ospf", "::1")
```
对Unix网络，地址必须是文件系统路径。

##### func DialTimeout
```go
func DialTimeout(network, address string, timeout time.Duration) (Conn, error)
```
DialTimeout类似Dial但采用了超时。timeout参数如果必要可包含名称解析。

##### func Pipe
```go
func Pipe() (Conn, Conn)
```
Pipe创建一个内存中的同步、全双工网络连接。连接的两端都实现了Conn接口。一端的读取对应另一端的写入，直接将数据在两端之间作拷贝；没有内部缓冲。

##### type PacketConn
```go
type PacketConn interface { 
    // ReadFrom方法从连接读取一个数据包，并将有效信息写入b 
    // ReadFrom方法可能会在超过某个固定时间限制后超时返回错误，该错误的Timeout()方法返回真 
    // 返回写入的字节数和该数据包的来源地址 
    ReadFrom(b []byte) (n int, addr Addr, err error) 
    // WriteTo方法将有效数据b写入一个数据包发送给addr 
    // WriteTo方法可能会在超过某个固定时间限制后超时返回错误，该错误的Timeout()方法返回真 
    // 在面向数据包的连接中，写入超时非常罕见 
    WriteTo(b []byte, addr Addr) (n int, err error) 
    // Close方法关闭该连接 
    // 会导致任何阻塞中的ReadFrom或WriteTo方法不再阻塞并返回错误 
    Close() error 
    // 返回本地网络地址 
    LocalAddr() Addr 
    // 设定该连接的读写deadline 
    SetDeadline(t time.Time) error 
    // 设定该连接的读操作deadline，参数t为零值表示不设置期限 
    // 如果时间到达deadline，读操作就会直接因超时失败返回而不会阻塞 
    SetReadDeadline(t time.Time) error 
    // 设定该连接的写操作deadline，参数t为零值表示不设置期限 
    // 如果时间到达deadline，写操作就会直接因超时失败返回而不会阻塞
    // 即使写入超时，返回值n也可能>0，说明成功写入了部分数据 
    SetWriteDeadline(t time.Time) error 
}
```
PacketConn接口代表通用的面向数据包的网络连接。多个线程可能会同时调用同一个Conn的方法。

##### func ListenPacket
```go
func ListenPacket(net, laddr string) (PacketConn, error)
```
ListenPacket函数监听本地网络地址laddr。网络类型net必须是面向数据包的网络类型：

"ip"、"ip4"、"ip6"、"udp"、"udp4"、"udp6"、或"unixgram"。laddr的格式参见Dial函数。

##### type Dialer
```go
type Dialer struct { 
    // Timeout是dial操作等待连接建立的最大时长，默认值代表没有超时。 
    // 如果Deadline字段也被设置了，dial操作也可能更早失败。 
    // 不管有没有设置超时，操作系统都可能强制执行它的超时设置。 
    // 例如，TCP（系统）超时一般在3分钟左右。 
    Timeout time.Duration 
    // Deadline是一个具体的时间点期限，超过该期限后，dial操作就会失败。 
    // 如果Timeout字段也被设置了，dial操作也可能更早失败。 
    // 零值表示没有期限，即遵守操作系统的超时设置。 
    Deadline time.Time 
    // LocalAddr是dial一个地址时使用的本地地址。 
    // 该地址必须是与dial的网络相容的类型。 
    // 如果为nil，将会自动选择一个本地地址。 
    LocalAddr Addr 
    // DualStack允许单次dial操作在网络类型为"tcp"， 
    // 且目的地是一个主机名的DNS记录具有多个地址时， 
    // 尝试建立多个IPv4和IPv6连接，并返回第一个建立的连接。 
    DualStack bool 
    // KeepAlive指定一个活动的网络连接的生命周期；如果为0，会禁止keep-alive。 
    // 不支持keep-alive的网络连接会忽略本字段。 
    KeepAlive time.Duration 
}
```
Dialer类型包含与某个地址建立连接时的参数。

每一个字段的零值都等价于没有该字段。因此调用Dialer零值的Dial方法等价于调用Dial函数。

##### func (*Dialer) Dial
```go
func (d *Dialer) Dial(network, address string) (Conn, error)
```
Dial在指定的网络上连接指定的地址。参见Dial函数获取网络和地址参数的描述。

##### type Listener
```go
type Listener interface { 
    // Addr返回该接口的网络地址 
    Addr() Addr 
    // Accept等待并返回下一个连接到该接口的连接 
    Accept() (c Conn, err error) // Close关闭该接口，并使任何阻塞的Accept操作都会不再阻塞并返回错误。 
    Close() error 
}
```
Listener是一个用于面向流的网络协议的公用的网络监听器接口。多个线程可能会同时调用一个Listener的方法。

Example
##### func Listen
```go
func Listen(net, laddr string) (Listener, error)
```
返回在一个本地网络地址laddr上监听的Listener。网络类型参数net必须是面向流的网络：

"tcp"、"tcp4"、"tcp6"、"unix"或"unixpacket"。参见Dial函数获取laddr的语法。

##### type IPAddr
```go
type IPAddr struct { 
    IP IP Zone string // IPv6范围寻址域 
}
```
IPAddr代表一个IP终端的地址。

##### func ResolveIPAddr
```go
func ResolveIPAddr(net, addr string) (*IPAddr, error)
```
ResolveIPAddr将addr作为一个格式为"host"或"ipv6-host%zone"的IP地址来解析。 函数会在参数net指定的网络类型上解析，net必须是"ip"、"ip4"或"ip6"。

##### func (*IPAddr) Network
```go
func (a *IPAddr) Network() string
```
Network返回地址的网络类型："ip"。

##### func (*IPAddr) String
```go
func (a *IPAddr) String() string
```
##### type TCPAddr
```go
type TCPAddr struct { 
    IP IP Port int Zone string // IPv6范围寻址域 
}
```
TCPAddr代表一个TCP终端地址。

##### func ResolveTCPAddr
```go
func ResolveTCPAddr(net, addr string) (*TCPAddr, error)
```
ResolveTCPAddr将addr作为TCP地址解析并返回。参数addr格式为"host:port"或"[ipv6-host%zone]:port"，解析得到网络名和端口名；net必须是"tcp"、"tcp4"或"tcp6"。

IPv6地址字面值/名称必须用方括号包起来，如"[::1]:80"、"[ipv6-host]:http"或"[ipv6-host%zone]:80"。

##### func (*TCPAddr) Network
```go
func (a *TCPAddr) Network() string
```
返回地址的网络类型，"tcp"。

##### func (*TCPAddr) String
```go
func (a *TCPAddr) String() string
```
##### type UDPAddr
```go
type UDPAddr struct { 
    IP IP Port int Zone string // IPv6范围寻址域 
}
```
UDPAddr代表一个UDP终端地址。

##### func ResolveUDPAddr
```go
func ResolveUDPAddr(net, addr string) (*UDPAddr, error)
```
ResolveTCPAddr将addr作为TCP地址解析并返回。参数addr格式为"host:port"或"[ipv6-host%zone]:port"，解析得到网络名和端口名；net必须是"udp"、"udp4"或"udp6"。

IPv6地址字面值/名称必须用方括号包起来，如"[::1]:80"、"[ipv6-host]:http"或"[ipv6-host%zone]:80"。

##### func (*UDPAddr) Network
```go
func (a *UDPAddr) Network() string
```
返回地址的网络类型，"udp"。

##### func (*UDPAddr) String
```go
func (a *UDPAddr) String() string
```
##### type UnixAddr
```go
type UnixAddr struct { 
    Name string Net string 
}
```
UnixAddr代表一个Unix域socket终端地址。

##### func ResolveUnixAddr
```go
func ResolveUnixAddr(net, addr string) (*UnixAddr, error)
```
ResolveUnixAddr将addr作为Unix域socket地址解析，参数net指定网络类型："unix"、"unixgram"或"unixpacket"。

##### func (*UnixAddr) Network
```go
func (a *UnixAddr) Network() string
```
返回地址的网络类型，"unix"，"unixgram"或"unixpacket"。

##### func (*UnixAddr) String
```go
func (a *UnixAddr) String() string
```
##### type IPConn
```go
type IPConn struct { 
    // 内含隐藏或非导出字段 
}
```
IPConn类型代表IP网络连接，实现了Conn和PacketConn接口。

##### func DialIP
```go
func DialIP(netProto string, laddr, raddr *IPAddr) (*IPConn, error)
```
DialIP在网络协议netProto上连接本地地址laddr和远端地址raddr，netProto必须是"ip"、"ip4"或"ip6"后跟冒号和协议名或协议号。

##### func ListenIP
```go
func ListenIP(netProto string, laddr *IPAddr) (*IPConn, error)
```
ListenIP创建一个接收目的地是本地地址laddr的IP数据包的网络连接，返回的*IPConn的ReadFrom和WriteTo方法可以用来发送和接收IP数据包。（每个包都可获取来源址或者设置目标地址）

##### func (*IPConn) LocalAddr
```go
func (c *IPConn) LocalAddr() Addr
```
LocalAddr返回本地网络地址

##### func (*IPConn) RemoteAddr
```go
func (c *IPConn) RemoteAddr() Addr
```
RemoteAddr返回远端网络地址

##### func (*IPConn) SetReadBuffer
```go
func (c *IPConn) SetReadBuffer(bytes int) error
```
SetReadBuffer设置该连接的系统接收缓冲

##### func (*IPConn) SetWriteBuffer
```go
func (c *IPConn) SetWriteBuffer(bytes int) error
```
SetWriteBuffer设置该连接的系统发送缓冲

##### func (*IPConn) SetDeadline
```go
func (c *IPConn) SetDeadline(t time.Time) error
```
SetDeadline设置读写操作绝对期限，实现了Conn接口的SetDeadline方法

##### func (*IPConn) SetReadDeadline
```go
func (c *IPConn) SetReadDeadline(t time.Time) error
```
SetReadDeadline设置读操作绝对期限，实现了Conn接口的SetReadDeadline方法

##### func (*IPConn) SetWriteDeadline
```go
func (c *IPConn) SetWriteDeadline(t time.Time) error
```
SetWriteDeadline设置写操作绝对期限，实现了Conn接口的SetWriteDeadline方法

##### func (*IPConn) Read
```go
func (c *IPConn) Read(b []byte) (int, error)
```
Read实现Conn接口Read方法

##### func (*IPConn) ReadFrom
```go
func (c *IPConn) ReadFrom(b []byte) (int, Addr, error)
```
ReadFrom实现PacketConn接口ReadFrom方法。注意本方法有bug，应避免使用。

##### func (*IPConn) ReadFromIP
```go
func (c *IPConn) ReadFromIP(b []byte) (int, *IPAddr, error)
```
ReadFromIP从c读取一个IP数据包，将有效负载拷贝到b，返回拷贝字节数和数据包来源地址。

ReadFromIP方法会在超过一个固定的时间点之后超时，并返回一个错误。注意本方法有bug，应避免使用。

##### func (*IPConn) ReadMsgIP
```go
func (c *IPConn) ReadMsgIP(b, oob []byte) (n, oobn, flags int, addr *IPAddr, err error)
```
ReadMsgIP从c读取一个数据包，将有效负载拷贝进b，相关的带外数据拷贝进oob，返回拷贝进b的字节数，拷贝进oob的字节数，数据包的flag，数据包来源地址和可能的错误。

##### func (*IPConn) Write
```go
func (c *IPConn) Write(b []byte) (int, error)
```
Write实现Conn接口Write方法

##### func (*IPConn) WriteTo
```go
func (c *IPConn) WriteTo(b []byte, addr Addr) (int, error)
```
WriteTo实现PacketConn接口WriteTo方法

##### func (*IPConn) WriteToIP
```go
func (c *IPConn) WriteToIP(b []byte, addr *IPAddr) (int, error)
```
WriteToIP通过c向地址addr发送一个数据包，b为包的有效负载，返回写入的字节。

WriteToIP方法会在超过一个固定的时间点之后超时，并返回一个错误。在面向数据包的连接上，写入超时是十分罕见的。

##### func (*IPConn) WriteMsgIP
```go
func (c *IPConn) WriteMsgIP(b, oob []byte, addr *IPAddr) (n, oobn int, err error)
```
WriteMsgIP通过c向地址addr发送一个数据包，b和oob分别为包有效负载和对应的带外数据，返回写入的字节数（包数据、带外数据）和可能的错误。

##### func (*IPConn) Close
```go
func (c *IPConn) Close() error
``
Close关闭连接

##### func (*IPConn) File
```go
func (c *IPConn) File() (f *os.File, err error)
```
File方法设置下层的os.File为阻塞模式并返回其副本。

使用者有责任在用完后关闭f。关闭c不影响f，关闭f也不影响c。返回的os.File类型文件描述符和原本的网络连接是不同的。试图使用该副本修改本体的属性可能会（也可能不会）得到期望的效果。

##### type TCPConn
```go
type TCPConn struct { 
    // 内含隐藏或非导出字段 
}
```
TCPConn代表一个TCP网络连接，实现了Conn接口。

##### func DialTCP
```go
func DialTCP(net string, laddr, raddr *TCPAddr) (*TCPConn, error)
```
DialTCP在网络协议net上连接本地地址laddr和远端地址raddr。net必须是"tcp"、"tcp4"、"tcp6"；如果laddr不是nil，将使用它作为本地地址，否则自动选择一个本地地址。

##### func (*TCPConn) LocalAddr
```go
func (c *TCPConn) LocalAddr() Addr
```
LocalAddr返回本地网络地址

##### func (*TCPConn) RemoteAddr
```go
func (c *TCPConn) RemoteAddr() Addr
```
RemoteAddr返回远端网络地址

##### func (*TCPConn) SetReadBuffer
```go
func (c *TCPConn) SetReadBuffer(bytes int) error
```
SetReadBuffer设置该连接的系统接收缓冲

##### func (*TCPConn) SetWriteBuffer
```go
func (c *TCPConn) SetWriteBuffer(bytes int) error
```
SetWriteBuffer设置该连接的系统发送缓冲

##### func (*TCPConn) SetDeadline
```go
func (c *TCPConn) SetDeadline(t time.Time) error
```
SetDeadline设置读写操作期限，实现了Conn接口的SetDeadline方法

##### func (*TCPConn) SetReadDeadline
```go
func (c *TCPConn) SetReadDeadline(t time.Time) error
```
SetReadDeadline设置读操作期限，实现了Conn接口的SetReadDeadline方法

##### func (*TCPConn) SetWriteDeadline
```go
func (c *TCPConn) SetWriteDeadline(t time.Time) error
```
SetWriteDeadline设置写操作期限，实现了Conn接口的SetWriteDeadline方法

##### func (*TCPConn) SetKeepAlive
```go
func (c *TCPConn) SetKeepAlive(keepalive bool) error
```
SetKeepAlive设置操作系统是否应该在该连接中发送keepalive信息

##### func (*TCPConn) SetKeepAlivePeriod
```go
func (c *TCPConn) SetKeepAlivePeriod(d time.Duration) error
```
SetKeepAlivePeriod设置keepalive的周期，超出会断开

##### func (*TCPConn) SetLinger
```go
func (c *TCPConn) SetLinger(sec int) error
```
SetLinger设定当连接中仍有数据等待发送或接受时的Close方法的行为。

如果sec < 0（默认），Close方法立即返回，操作系统停止后台数据发送；如果 sec == 0，Close立刻返回，操作系统丢弃任何未发送或未接收的数据；如果sec > 0，Close方法阻塞最多sec秒，等待数据发送或者接收，在一些操作系统中，在超时后，任何未发送的数据会被丢弃。

##### func (*TCPConn) SetNoDelay
```go
func (c *TCPConn) SetNoDelay(noDelay bool) error
```
SetNoDelay设定操作系统是否应该延迟数据包传递，以便发送更少的数据包（Nagle's算法）。默认为真，即数据应该在Write方法后立刻发送。

##### func (*TCPConn) Read
```go
func (c *TCPConn) Read(b []byte) (int, error)
```
Read实现了Conn接口Read方法

##### func (*TCPConn) Write
```go
func (c *TCPConn) Write(b []byte) (int, error)
```
Write实现了Conn接口Write方法

##### func (*TCPConn) ReadFrom
```go
func (c *TCPConn) ReadFrom(r io.Reader) (int64, error)
```
ReadFrom实现了io.ReaderFrom接口的ReadFrom方法

##### func (*TCPConn) Close
```go
func (c *TCPConn) Close() error
```
Close关闭连接

##### func (*TCPConn) CloseRead
```go
func (c *TCPConn) CloseRead() error
```
CloseRead关闭TCP连接的读取侧（以后不能读取），应尽量使用Close方法。

##### func (*TCPConn) CloseWrite
```go
func (c *TCPConn) CloseWrite() error
```
CloseWrite关闭TCP连接的写入侧（以后不能写入），应尽量使用Close方法。

##### func (*TCPConn) File
```go
func (c *TCPConn) File() (f *os.File, err error)
```
File方法设置下层的os.File为阻塞模式并返回其副本。

使用者有责任在用完后关闭f。关闭c不影响f，关闭f也不影响c。返回的os.File类型文件描述符和原本的网络连接是不同的。试图使用该副本修改本体的属性可能会（也可能不会）得到期望的效果。

##### type UDPConn
```go
type UDPConn struct { 
    // 内含隐藏或非导出字段 
}
```
UDPConn代表一个UDP网络连接，实现了Conn和PacketConn接口。

##### func DialUDP
```go
func DialUDP(net string, laddr, raddr *UDPAddr) (*UDPConn, error)
```
DialTCP在网络协议net上连接本地地址laddr和远端地址raddr。net必须是"udp"、"udp4"、"udp6"；如果laddr不是nil，将使用它作为本地地址，否则自动选择一个本地地址。

##### func ListenUDP
```go
func ListenUDP(net string, laddr *UDPAddr) (*UDPConn, error)
```
ListenUDP创建一个接收目的地是本地地址laddr的UDP数据包的网络连接。net必须是"udp"、"udp4"、"udp6"；如果laddr端口为0，函数将选择一个当前可用的端口，可以用Listener的Addr方法获得该端口。返回的*UDPConn的ReadFrom和WriteTo方法可以用来发送和接收UDP数据包（每个包都可获得来源地址或设置目标地址）。

##### func ListenMulticastUDP
```go
func ListenMulticastUDP(net string, ifi *Interface, gaddr *UDPAddr) (*UDPConn, error)
```
ListenMulticastUDP接收目的地是ifi接口上的组地址gaddr的UDP数据包。它指定了使用的接口，如果ifi是nil，将使用默认接口。

##### func (*UDPConn) LocalAddr
```go
func (c *UDPConn) LocalAddr() Addr
```
LocalAddr返回本地网络地址

##### func (*UDPConn) RemoteAddr
```go
func (c *UDPConn) RemoteAddr() Addr
```
RemoteAddr返回远端网络地址

##### func (*UDPConn) SetReadBuffer
```go
func (c *UDPConn) SetReadBuffer(bytes int) error
```
SetReadBuffer设置该连接的系统接收缓冲

##### func (*UDPConn) SetWriteBuffer
```go
func (c *UDPConn) SetWriteBuffer(bytes int) error
```
SetWriteBuffer设置该连接的系统发送缓冲

##### func (*UDPConn) SetDeadline
```go
func (c *UDPConn) SetDeadline(t time.Time) error
```
SetDeadline设置读写操作期限，实现了Conn接口的SetDeadline方法

##### func (*UDPConn) SetReadDeadline
```go
func (c *UDPConn) SetReadDeadline(t time.Time) error
```
SetReadDeadline设置读操作期限，实现了Conn接口的SetReadDeadline方法

##### func (*UDPConn) SetWriteDeadline
```go
func (c *UDPConn) SetWriteDeadline(t time.Time) error
```
SetWriteDeadline设置写操作期限，实现了Conn接口的SetWriteDeadline方法

##### func (*UDPConn) Read
```go
func (c *UDPConn) Read(b []byte) (int, error)
```
Read实现Conn接口Read方法

##### func (*UDPConn) ReadFrom
```go
func (c *UDPConn) ReadFrom(b []byte) (int, Addr, error)
```
ReadFrom实现PacketConn接口ReadFrom方法

##### func (*UDPConn) ReadFromUDP
```go
func (c *UDPConn) ReadFromUDP(b []byte) (n int, addr *UDPAddr, err error)
```
ReadFromUDP从c读取一个UDP数据包，将有效负载拷贝到b，返回拷贝字节数和数据包来源地址。

ReadFromUDP方法会在超过一个固定的时间点之后超时，并返回一个错误。

##### func (*UDPConn) ReadMsgUDP
```go
func (c *UDPConn) ReadMsgUDP(b, oob []byte) (n, oobn, flags int, addr *UDPAddr, err error)
```
ReadMsgUDP从c读取一个数据包，将有效负载拷贝进b，相关的带外数据拷贝进oob，返回拷贝进b的字节数，拷贝进oob的字节数，数据包的flag，数据包来源地址和可能的错误。

##### func (*UDPConn) Write
```go
func (c *UDPConn) Write(b []byte) (int, error)
```
Write实现Conn接口Write方法

##### func (*UDPConn) WriteTo
```go
func (c *UDPConn) WriteTo(b []byte, addr Addr) (int, error)
```
WriteTo实现PacketConn接口WriteTo方法

##### func (*UDPConn) WriteToUDP
```go
func (c *UDPConn) WriteToUDP(b []byte, addr *UDPAddr) (int, error)
```
WriteToUDP通过c向地址addr发送一个数据包，b为包的有效负载，返回写入的字节。

WriteToUDP方法会在超过一个固定的时间点之后超时，并返回一个错误。在面向数据包的连接上，写入超时是十分罕见的。

##### func (*UDPConn) WriteMsgUDP
```go
func (c *UDPConn) WriteMsgUDP(b, oob []byte, addr *UDPAddr) (n, oobn int, err error)
```
WriteMsgUDP通过c向地址addr发送一个数据包，b和oob分别为包有效负载和对应的带外数据，返回写入的字节数（包数据、带外数据）和可能的错误。

##### func (*UDPConn) Close
```go
func (c *UDPConn) Close() error
```
Close关闭连接

##### func (*UDPConn) File
```go
func (c *UDPConn) File() (f *os.File, err error)
```
File方法设置下层的os.File为阻塞模式并返回其副本。

使用者有责任在用完后关闭f。关闭c不影响f，关闭f也不影响c。返回的os.File类型文件描述符和原本的网络连接是不同的。试图使用该副本修改本体的属性可能会（也可能不会）得到期望的效果。

##### type UnixConn
```go
type UnixConn struct { 
    // 内含隐藏或非导出字段 
}
```
UnixConn代表Unix域socket连接，实现了Conn和PacketConn接口。

##### func DialUnix
```go
func DialUnix(net string, laddr, raddr *UnixAddr) (*UnixConn, error)
```
DialUnix在网络协议net上连接本地地址laddr和远端地址raddr。net必须是"unix"、"unixgram"、"unixpacket"，如果laddr不是nil将使用它作为本地地址，否则自动选择一个本地地址。

##### func ListenUnixgram
```go
func ListenUnixgram(net string, laddr *UnixAddr) (*UnixConn, error)
```
ListenUnixgram接收目的地是本地地址laddr的Unix datagram网络连接。net必须是"unixgram"，返回的*UnixConn的ReadFrom和WriteTo方法可以用来发送和接收数据包（每个包都可获取来源址或者设置目标地址）。

##### func (*UnixConn) LocalAddr
```go
func (c *UnixConn) LocalAddr() Addr
```
LocalAddr返回本地网络地址

##### func (*UnixConn) RemoteAddr
```go
func (c *UnixConn) RemoteAddr() Addr
```
RemoteAddr返回远端网络地址

##### func (*UnixConn) SetReadBuffer
```go
func (c *UnixConn) SetReadBuffer(bytes int) error
```
SetReadBuffer设置该连接的系统接收缓冲

##### func (*UnixConn) SetWriteBuffer
```go
func (c *UnixConn) SetWriteBuffer(bytes int) error
```
SetWriteBuffer设置该连接的系统发送缓冲

##### func (*UnixConn) SetDeadline
```go
func (c *UnixConn) SetDeadline(t time.Time) error
```
SetDeadline设置读写操作期限，实现了Conn接口的SetDeadline方法

##### func (*UnixConn) SetReadDeadline
```go
func (c *UnixConn) SetReadDeadline(t time.Time) error
```
SetReadDeadline设置读操作期限，实现了Conn接口的SetReadDeadline方法

##### func (*UnixConn) SetWriteDeadline
```go
func (c *UnixConn) SetWriteDeadline(t time.Time) error
```
SetWriteDeadline设置写操作期限，实现了Conn接口的SetWriteDeadline方法

##### func (*UnixConn) Read
```go
func (c *UnixConn) Read(b []byte) (int, error)
```
Read实现了Conn接口Read方法

##### func (*UnixConn) ReadFrom
```go
func (c *UnixConn) ReadFrom(b []byte) (int, Addr, error)
```
ReadFrom实现PacketConn接口ReadFrom方法

##### func (*UnixConn) ReadFromUnix
```go
func (c *UnixConn) ReadFromUnix(b []byte) (n int, addr *UnixAddr, err error)
```
ReadFromUnix从c读取一个UDP数据包，将有效负载拷贝到b，返回拷贝字节数和数据包来源地址。

ReadFromUnix方法会在超过一个固定的时间点之后超时，并返回一个错误。

##### func (*UnixConn) ReadMsgUnix
```go
func (c *UnixConn) ReadMsgUnix(b, oob []byte) (n, oobn, flags int, addr *UnixAddr, err error)
```
ReadMsgUnix从c读取一个数据包，将有效负载拷贝进b，相关的带外数据拷贝进oob，返回拷贝进b的字节数，拷贝进oob的字节数，数据包的flag，数据包来源地址和可能的错误。

##### func (*UnixConn) Write
```go
func (c *UnixConn) Write(b []byte) (int, error)
```
Write实现了Conn接口Write方法

##### func (*UnixConn) WriteTo
```go
func (c *UnixConn) WriteTo(b []byte, addr Addr) (n int, err error)
```
WriteTo实现PacketConn接口WriteTo方法

##### func (*UnixConn) WriteToUnix
```go
func (c *UnixConn) WriteToUnix(b []byte, addr *UnixAddr) (n int, err error)
```
WriteToUnix通过c向地址addr发送一个数据包，b为包的有效负载，返回写入的字节。

WriteToUnix方法会在超过一个固定的时间点之后超时，并返回一个错误。在面向数据包的连接上，写入超时是十分罕见的。

##### func (*UnixConn) WriteMsgUnix
```go
func (c *UnixConn) WriteMsgUnix(b, oob []byte, addr *UnixAddr) (n, oobn int, err error)
```
WriteMsgUnix通过c向地址addr发送一个数据包，b和oob分别为包有效负载和对应的带外数据，返回写入的字节数（包数据、带外数据）和可能的错误。

##### func (*UnixConn) Close
```go
func (c *UnixConn) Close() error
```
Close关闭连接

##### func (*UnixConn) CloseRead
```go
func (c *UnixConn) CloseRead() error
```
CloseRead关闭TCP连接的读取侧（以后不能读取），应尽量使用Close方法

##### func (*UnixConn) CloseWrite
```go
func (c *UnixConn) CloseWrite() error
```
CloseWrite关闭TCP连接的写入侧（以后不能写入），应尽量使用Close方法

##### func (*UnixConn) File
```go 
func (c *UnixConn) File() (f *os.File, err error)
```
File方法设置下层的os.File为阻塞模式并返回其副本。

使用者有责任在用完后关闭f。关闭c不影响f，关闭f也不影响c。返回的os.File类型文件描述符和原本的网络连接是不同的。试图使用该副本修改本体的属性可能会（也可能不会）得到期望的效果。

##### type TCPListener
```go
type TCPListener struct { 
    // 内含隐藏或非导出字段 
}
```
TCPListener代表一个TCP网络的监听者。使用者应尽量使用Listener接口而不是假设（网络连接为）TCP。

##### func ListenTCP
```go
func ListenTCP(net string, laddr *TCPAddr) (*TCPListener, error)
```
ListenTCP在本地TCP地址laddr上声明并返回一个*TCPListener，net参数必须是"tcp"、"tcp4"、"tcp6"，如果laddr的端口字段为0，函数将选择一个当前可用的端口，可以用Listener的Addr方法获得该端口。

##### func (*TCPListener) Addr
```go
func (l *TCPListener) Addr() Addr
```
Addr返回l监听的的网络地址，一个*TCPAddr。

##### func (*TCPListener) SetDeadline
```go
func (l *TCPListener) SetDeadline(t time.Time) error
```
设置监听器执行的期限，t为Time零值则会关闭期限限制。

##### func (*TCPListener) Accept
```go
func (l *TCPListener) Accept() (Conn, error)
```
Accept用于实现Listener接口的Accept方法；他会等待下一个呼叫，并返回一个该呼叫的Conn接口。

##### func (*TCPListener) AcceptTCP
```go
func (l *TCPListener) AcceptTCP() (*TCPConn, error)
```
AcceptTCP接收下一个呼叫，并返回一个新的*TCPConn。

##### func (*TCPListener) Close
```go
func (l *TCPListener) Close() error
```
Close停止监听TCP地址，已经接收的连接不受影响。

##### func (*TCPListener) File
```go 
func (l *TCPListener) File() (f *os.File, err error)
```
File方法返回下层的os.File的副本，并将该副本设置为阻塞模式。

使用者有责任在用完后关闭f。关闭c不影响f，关闭f也不影响c。返回的os.File类型文件描述符和原本的网络连接是不同的。试图使用该副本修改本体的属性可能会（也可能不会）得到期望的效果。

##### type UnixListener
```go
type UnixListener struct { // 内含隐藏或非导出字段 }
```
UnixListener代表一个Unix域scoket的监听者。使用者应尽量使用Listener接口而不是假设（网络连接为）Unix域scoket。

##### func ListenUnix
```go 
func ListenUnix(net string, laddr *UnixAddr) (*UnixListener, error)
```
ListenTCP在Unix域scoket地址laddr上声明并返回一个*UnixListener，net参数必须是"unix"或"unixpacket"。

##### func (*UnixListener) Addr
```go 
func (l *UnixListener) Addr() Addr
```
Addr返回l的监听的Unix域socket地址

##### func (*UnixListener) SetDeadline
```go
func (l *UnixListener) SetDeadline(t time.Time) (err error)
```
设置监听器执行的期限，t为Time零值则会关闭期限限制

##### func (*UnixListener) Accept
```go
func (l *UnixListener) Accept() (c Conn, err error)
```
Accept用于实现Listener接口的Accept方法；他会等待下一个呼叫，并返回一个该呼叫的Conn接口。

##### func (*UnixListener) AcceptUnix
```go
func (l *UnixListener) AcceptUnix() (*UnixConn, error)
```
AcceptUnix接收下一个呼叫，并返回一个新的*UnixConn。

##### func (*UnixListener) Close
```go 
func (l *UnixListener) Close() error
```
Close停止监听Unix域socket地址，已经接收的连接不受影响。

##### func (*UnixListener) File
```go 
func (l *UnixListener) File() (f *os.File, err error)
```
File方法返回下层的os.File的副本，并将该副本设置为阻塞模式。

使用者有责任在用完后关闭f。关闭c不影响f，关闭f也不影响c。返回的os.File类型文件描述符和原本的网络连接是不同的。试图使用该副本修改本体的属性可能会（也可能不会）得到期望的效果。

##### func FileConn
```go 
func FileConn(f *os.File) (c Conn, err error)
```
FileConn返回一个下层为文件f的网络连接的拷贝。调用者有责任在结束程序前关闭f。关闭c不会影响f，关闭f也不会影响c。本函数与各种实现了Conn接口的类型的File方法是对应的。

##### func FilePacketConn
```go
func FilePacketConn(f *os.File) (c PacketConn, err error)
```
FilePacketConn函数返回一个下层为文件f的数据包网络连接的拷贝。调用者有责任在结束程序前关闭f。关闭c不会影响f，关闭f也不会影响c。本函数与各种实现了PacketConn接口的类型的File方法是对应的。

##### func FileListener
```go 
func FileListener(f *os.File) (l Listener, err error)
```
FileListener返回一个下层为文件f的网络监听器的拷贝。调用者有责任在使用结束后改变l。关闭l不会影响f，关闭f也不会影响l。本函数与各种实现了Listener接口的类型的File方法是对应的。

##### type MX
```go
type MX struct { Host string Pref uint16 }
```
MX代表一条DNS MX记录（邮件交换记录），根据收信人的地址后缀来定位邮件服务器。

##### type NS
```go
type NS struct { Host string }
```
NS代表一条DNS NS记录（域名服务器记录），指定该域名由哪个DNS服务器来进行解析。

##### type SRV
```go
type SRV struct { Target string Port uint16 Priority uint16 Weight uint16 }
```
SRV代表一条DNS SRV记录（资源记录），记录某个服务由哪台计算机提供。

##### func LookupPort
```go 
func LookupPort(network, service string) (port int, err error)
```
LookupPort函数查询指定网络和服务的（默认）端口。

##### func LookupCNAME
```go
func LookupCNAME(name string) (cname string, err error)
```
LookupCNAME函数查询name的规范DNS名（但该域名未必可以访问）。如果调用者不关心规范名可以直接调用LookupHost或者LookupIP；这两个函数都会在查询时考虑到规范名。

##### func LookupHost
```go 
func LookupHost(host string) (addrs []string, err error)
```
LookupHost函数查询主机的网络地址序列。

##### func LookupIP
```go 
func LookupIP(host string) (addrs []IP, err error)
```
LookupIP函数查询主机的ipv4和ipv6地址序列。

##### func LookupAddr
```go
func LookupAddr(addr string) (name []string, err error)
```
LookupAddr查询某个地址，返回映射到该地址的主机名序列，本函数和LookupHost不互为反函数。

##### func LookupMX
```go 
func LookupMX(name string) (mx []*MX, err error)
```
LookupMX函数返回指定主机的按Pref字段排好序的DNS MX记录。

##### func LookupNS
```go 
func LookupNS(name string) (ns []*NS, err error)
```
LookupNS函数返回指定主机的DNS NS记录。

##### func LookupSRV
```go
func LookupSRV(service, proto, name string) (cname string, addrs []*SRV, err error)
```
LookupSRV函数尝试执行指定服务、协议、主机的SRV查询。协议proto为"tcp" 或"udp"。返回的记录按Priority字段排序，同一优先度按Weight字段随机排序。

LookupSRV函数按照RFC 2782的规定构建用于查询的DNS名。也就是说，它会查询_service._proto.name。为了适应将服务的SRV记录发布在非规范名下的情况，如果service和proto参数都是空字符串，函数会直接查询name。

##### func LookupTXT
```go
func LookupTXT(name string) (txt []string, err error)
```
LookupTXT函数返回指定主机的DNS TXT记录。

Bugs

☞在任何POSIX平台上，从"ip4"网络使用ReadFrom或ReadFromIP方法读取数据时，即使有足够的空间，都可能不会返回完整的IPv4数据包，包括数据包的头域。即使Read或ReadMsgIP方法可以返回完整的数据包，也有可能出现这种情况。因为对go 1的兼容性要求，这个情况无法被修正。因此，当必须获取完整数据包时，建议你不要使用这两个方法，请使用Read或ReadMsgIP代替。

☞在OpenBSD系统中，在"tcp"网络监听时不会同时监听IPv4和IPv6连接。 因为该系统中IPv4通信不会导入IPv6套接字中。请使用两个独立的监听，如果有必要的话。