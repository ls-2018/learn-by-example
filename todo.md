- bpf-tailcall-tracer                                                                # 
- bpfbak                                                                             # 
- bpfsyscalldist                                                                     # 
- cframe                                                                             # 
- dnsproxy-go                                                                        # 
- drgn-bpf                                                                           # 
- eBPF-TupleMerge                                                                    # 
- ebpf-vm-on-ebpf                                                                    # 
- ebpfdbg                                                                            # 
- ebpfmanager                                                                        # 
- ecapture                                                                           # 
- ethtool                                                                            # 
- ethtoolsnoop                                                                       # 
- functrace                                                                          # 
- go-iproute2                                                                        # 
- go-nfnetlink-example                                                               # 
- go-tc                                                                              # 
- go-tproxy                                                                          # 
- iptables-bpf                                                                       # 
- iptables-trace                                                                     # 
- kcptun                                                                             # 
- kernel-module-fun                                                                  # 
- knetstat                                                                           # 
- l2fwd                                                                              # 
- l2tp-ipsec-vpn-server                                                              # 
- lb-from-scratch                                                                    # 
- libbpf-bootstrap-examples                                                          # 
- ping-latency-injector                                                              # 
- pkt-stucker                                                                        # 
- pwru                                                                               # 
- skbdist                                                                            # 
- skbtracer                                                                          # 
- skbtracer-iptables                                                                 # 
- sockdump                                                                           # 
- socketrace                                                                         # 
- strace                                                                             # 
- syscalldist                                                                        # 
- tc-dump                                                                            # 
- tproxy-experiment                                                                  # 
- xdp_acl                                                                            # ✅与acl一样,只是具体的拦截规则变了
- iptables-in-bpf                                                                    # ✅与acl一样,只是具体的拦截规则变了
- xdpsnoop


sysctl_ipv4 修改ipv4相关参数的进程、网卡


# map
```
├── iptables
│   └── nfqueue                     # Netfilter队列，允许iptables的数据包经过用户空间程序的处理后再决定是否接受、拒绝或修改这些数据包
├── ebpf
│   ├── acl
│   │   └── ping_disable            # 禁止 ping 某个地址
│   ├── bpf2bpf
│   │   └── ebpf
│   ├── bpfprogfuncs
│   ├── bsearch
│   ├── common
│   ├── fentry-bpf2bpf
│   ├── fentry_fexit
│   ├── fentry_fexit-bpf2bpf
│   ├── fentry_fexit-freplace
│   ├── fentry_fexit-freplace-xdp
│   ├── fentry_fexit-kprobe
│   ├── fentry_fexit-tailcall
│   ├── fentry_fexit-tailcall_in_bpf2bpf
│   ├── fentry_fexit-tc
│   ├── fentry_fexit-tracepoint
│   ├── fentry_fexit-xdp
│   ├── fexit_ipv4_sysctl
│   ├── fexit_rpsxps
│   ├── freplace                    # 替换某个函数
│   ├── global-variable
│   ├── headers
│   ├── inject
│   │   ├── cmd
│   │   │   ├── ebpf-inject-global-var
│   │   │   └── ebpf-inject-replace-const
│   │   └── ebpf
│   ├── iptables-bpf
│   ├── iptables-trace
│   │   ├── ebpf
│   │   └── kernel
│   ├── iter
│   ├── kernel-module-fun
│   │   └── custom-netlink
│   ├── kfunc_ffs
│   ├── metadata_xdp2afxdp
│   ├── n-args
│   ├── switch
│   ├── tailcall
│   │   └── ebpf
│   ├── tailcall-in-bpf2bpf
│   ├── tailcall-in-freplace
│   ├── tailcall-in-freplace1
│   ├── tailcall-shared
│   ├── tailcall-stackoverflow
│   ├── tcx
│   ├── timer
│   ├── tracepoint
│   ├── xdp-cpumap
│   ├── xdp
│   │   └── crc                 # 寄存器过度优化导致的问题
│   ├── xdp-traceroute-bpffs
│   ├── xdpmetadata
│   │   └── scripts
│   └── xdpping




```


- https://blog.csdn.net/weixin_40539956/article/details/137938104