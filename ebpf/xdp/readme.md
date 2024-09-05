#### 通过iproute2 ip加载
```bash
sudo ip link set dev lo xdpgeneric obj xdp_pass_kern.o sec xdp
sudo ip link show dev lo # 如果没有 sudo 的情况下运行它,将获得更少的信息
sudo ip link set dev lo xdpgeneric off
```

#### 使用 xdp-loader 加载

sudo xdp-loader load -m skb lo xdp_pass_kern.o
sudo xdp-loader status lo


