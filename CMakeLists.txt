cmake_minimum_required(VERSION 3.27.7)
project(learn-by-example C)

set(CMAKE_C_STANDARD 11)

include_directories(
        headers
)

# 无用
add_library(
        test
        ebpf/acl/ping_disable/acl.c
        ebpf/xdp/xdpping/xdp_ping.c
        ebpf/iptables-trace/ebpf/iptables-trace.c
        ebpf/iptables-trace/ebpf/headers/bpf_kprobe_args.h
        ebpf/utils/switch.c
        ebpf/kfunc_ffs/ffs.c
        ebpf/tailcall/tailcall/tcp-connecting.c
        ebpf/tailcall/bpf2bpf/tailcall.c
        ebpf/tailcall/freplace/freplace.c
        ebpf/tailcall/freplace/tcp-connecting.c
        ebpf/tailcall-shared/tcp-connecting.c
)

target_compile_definitions(test PRIVATE -D__clang_major__=8 -D__bpf__)