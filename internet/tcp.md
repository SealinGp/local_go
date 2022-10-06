https://draveness.me/whys-the-design-tcp-three-way-handshake/

1. 定义 TCP 连接
- socket: ip+port
- 序列号(SEQ): 跟踪包的传递
- 窗口大小: 流控制

2. 历史连接
   网络较差的情况下, A -> B 发起创建连接的请求
   A 发 3 个创建连接的请求
   B 收 3 个创建连接的请求
   由于网络不稳定,3 个连接在 B 这边并不是按顺序收到的,因此 B 无法判断哪个是最新的连接,其他的都是历史连接
   因此 B 收到请求后,将连接的 SEQ+1 返回给接收方,由接收方判断
   A收到第二次握手后,将历史连接通过RST标志位中止连接,将非历史连接通过 ACK 回应第3次握手,创建连接成功
   3次握手+RST标志位控制连接的生命周期,将控制权交给发送方,发送方有足够的上下文来判断连接的最终创建

3. 4次挥手
FIN_WAIT1 -> CLOSE_WAIT
FIN_WAIT2 <- ACK

TIME_WAIT <- FIN LAST_ACK
    ACK   -> CLOSED

2MSL
CLOSED

4. 网络协议
// application  HTTP、HTTPS、FTP、Telnet、SSH、SMTP、POP3
// transport   tcp/udp
// network      ip
// data link layer
// physical 

//应用
//表现
//会话
//传输      tcp/udp
//网络      IP
//数据链路 以太网
//物理    网卡接口

5. TCP粘包问题
