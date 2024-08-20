# 使用 Go 对接 iptables-nfqueue 的例子

## 例子的效果

使用 `iptables NFQUEUE` 监听新建 tcp 连接(作为服务端)：

```bash
./nfqueue-example
tcp connect: 113.81.xx.yy:19885 -> 10.7.xxx.yyy:22
tcp connect: 157.245.xx.yy:46131 -> 10.7.xxx.yyy:2060
tcp connect: 113.81.xx.yy:19907 -> 10.7.xxx.yyy:8080
tcp connect: 113.81.xx.yy:19918 -> 10.7.xxx.yyy:443
tcp connect: 113.81.xx.yy:19918 -> 10.7.xxx.yyy:443
tcp connect: 113.81.xx.yy:19918 -> 10.7.xxx.yyy:443
tcp connect: 46.101.xx.yy:48780 -> 10.7.xxx.yyy:5537
tcp connect: 89.248.xx.yy:54067 -> 10.7.xxx.yyy:309
```


```
Netfilter队列（nfqueue）是一种机制，允许iptables的数据包经过用户空间程序的处理后再决定是否接受、拒绝或修改这些数据包。
iptables -t raw -I PREROUTING -p tcp --syn -j NFQUEUE --queue-num=1 --queue-bypass   
iptables -t raw -D PREROUTING -p tcp --syn -j NFQUEUE --queue-num=1 --queue-bypass
cat  /proc/net/netfilter/nfnetlink_queue


iptables: 这是 Linux 中用于配置网络传输相关规则的工具，特别是包过滤防火墙。

-t raw: 指定要操作的表是 raw 表。raw 表通常用于配置与原始数据包相关的规则，比如在数据包的早期处理阶段进行规则匹配。

-I PREROUTING: -I 是插入（Insert）的缩写，表示将一条新规则插入到指定的链中。PREROUTING 是 raw 表中的一个链，用于处理到达本机、在路由决策之前的数据包。

-p tcp: 指定要匹配的数据包协议是 TCP。

--syn: 指定只匹配 TCP 握手中的 SYN 包，即初始连接请求包。这通常用于识别新的入站 TCP 连接尝试。

-j NFQUEUE: -j 指定规则匹配后执行的动作。NFQUEUE 是一个动作，它将匹配的数据包发送到用户空间的 NFQUEUE 队列中，可以由用户空间程序处理。

--queue-num=1: 指定 NFQUEUE 的编号，这里是 1。每个 NFQUEUE 都有一个编号，可以同时运行多个 NFQUEUE 以处理不同类型的数据包。

--queue-bypass: 这个选项指定当 NFQUEUE 队列满时的行为。--queue-bypass 表示如果队列满了，数据包将不会被丢弃，而是绕过 NFQUEUE 继续正常处理。这可以防止因为用户空间程序处理不过来而导致的数据包丢失。

综上所述，这条命令的作用是：将所有到达本机、使用 TCP 协议、并且携带 SYN 标志的入站数据包（即新的 TCP 连接请求）发送到编号为 1 的 NFQUEUE 队列中，
如果队列满了，这些数据包将绕过 NFQUEUE 继续正常处理。这通常用于在用户空间对新的 TCP 连接请求进行进一步的检查或处理。

```

