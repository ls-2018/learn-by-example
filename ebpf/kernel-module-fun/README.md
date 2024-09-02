Kernel module fun
=================

## Motivation

I didn't know at all how kernel modules worked. This is me learning
how. This is all tested using the `4.19.0-9` kernel.

## Contents

**`hello.c`**: a simple "hello world" module

**`who-connect-me.c`**: 基于 TCP SYN packet

**`add-arp-records.c`**: a custom netfilter hook to add arp records to global arp table

**`check-tcp-syncookies.c`**: a custom netfilter hook to check mss from tcp syncookies

**`custom-netlink.c`**: a custom netlink to communicate with Go

**`kprobe_tcp_conn_request`**: a custom kprobe to learn getting argument from kprobing function

**`run-bpf-prog`**: run bpf prog in kernel module, based on [github.com/Asphaltt/iptables-bpf](https://github.com/Asphaltt/iptables-bpf)



## Inserting into your kernel (at your own risk!)

sudo insmod hello.ko
dmesg | tail
sudo rmmod hello.ko

insmod ./run-bpf-prog.ko bpf_path=/sys/fs/bpf/iptbpf
ping -c4 114.114.114.114
rmmod `run_bpf_prog`
dmesg | tail



should display the "hello world" message

