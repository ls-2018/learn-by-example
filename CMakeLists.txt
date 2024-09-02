cmake_minimum_required(VERSION 3.27.7)
project(learn-by-example C)

set(CMAKE_C_STANDARD 11)

include_directories(
        headers
)

# 无用
add_library(
        test
        fentry_fexit/tailcall_in_bpf2bpf/tcp-connecting.c
        fentry_fexit/tailcall_in_bpf2bpf/fentry_fexit.c
        fentry_fexit/xdp/xdp.c
        fentry_fexit/tracepoint/tracepoint.c
        fentry_fexit/tracepoint/fentry_fexit.c
        ebpf/fentry_fexit-tailcall/tcp-connecting.c
        ebpf/fentry_fexit-tailcall/fentry_fexit.c
        ebpf/iptables-trace/ebpf/iptables-trace.c
        ebpf/iptables-trace/kernel/iptables-trace.c
        ebpf/fexit/sysctl_ipv4/sysctl.c
        ebpf/fentry_fexit-bpf2bpf/tcp-connecting.c
        ebpf/fentry_fexit-bpf2bpf/fentry_fexit.c
        ebpf/inject/ebpf/inject_replace_const.c
        ebpf/inject/ebpf/inject_global_var.c
        ebpf/fentry_fexit-freplace-xdp/freplace.c
        ebpf/fentry_fexit-freplace-xdp/fentry_fexit.c
        ebpf/fentry_fexit-freplace-xdp/xdp.c
        ebpf/timer/tcp-connecting.c
        ebpf/hack/global-variable/tcp-connecting.c
        ebpf/xdp/xdpping/xdp_ping.c
        ebpf/xdp/crc/xdp.c
        ebpf/xdp/traceroute-bpffs/traceroute.c
        ebpf/xdp/cpumap/xdp.c
        ebpf/fentry_fexit/fentry_fexit/tcp-connecting.c
        ebpf/fentry_fexit/tc/tc.c
        ebpf/fentry-bpf2bpf/tcp-connecting.c
        ebpf/fentry-bpf2bpf/fentry.c
        ebpf/utils/switch.c
        ebpf/tailcall/tailcall/tcp-connecting.c
        ebpf/tailcall/shared/tcp-connecting.c
        ebpf/tailcall/shared/xdp.c
        ebpf/tailcall/stackoverflow/tcp-connecting.c
        ebpf/tailcall/stackoverflow/fexit.c
        ebpf/tailcall/stackoverflow/fentry.c
        ebpf/tailcall/freplace/tcp-connecting.c
        ebpf/tailcall/freplace/freplace.c
        ebpf/tailcall/call_static/tailcall.c
        ebpf/tc/ingress/tcx.c
        ebpf/fentry_fexit-freplace/tcp-connecting.c
        ebpf/fentry_fexit-freplace/freplace.c
        ebpf/fentry_fexit-freplace/fentry_fexit.c
        ebpf/fentry_fexit-kprobe/tcp-connecting.c
        ebpf/fentry_fexit-kprobe/fentry_fexit.c
        ebpf/tracepoint/tracepoint.c
        ebpf/iptables/iptables-bpf/iptables-bpf.c
        ebpf/kernel-module-fun/add-arp-records.c
        ebpf/kernel-module-fun/hello.c
        ebpf/kernel-module-fun/custom-netlink.c
        ebpf/kernel-module-fun/excited_virus.c
        ebpf/kernel-module-fun/who-connect-me.c
        ebpf/kernel-module-fun/kprobe_tcp_conn_request.c
        ebpf/kernel-module-fun/run-bpf-prog.c
        ebpf/kernel-module-fun/check-tcp-syncookies.c
        ebpf/fexit_rpsxps/rpsxps.c
        ebpf/metadata/xdp2afxdp/xdp.c
        ebpf/metadata/metadata/xdp_tailcall.c
        ebpf/metadata/metadata/tc_metadata.c
        ebpf/acl/ping_disable/acl.c

)

target_compile_definitions(test PRIVATE -D__clang_major__=8 -D__bpf__)