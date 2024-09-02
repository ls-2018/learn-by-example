/**
 * Copyright 2023 Leon Hwang.
 * SPDX-License-Identifier: MIT
 */

//go:build ignore

#include "bpf_all.h"

#include "lib_tp_msg.h"

struct netlink_extack_error_ctx {
    unsigned long unused;

    /*
     * bpf does not support tracepoint __data_loc directly.
     * 实际上，这个字段是一个32位整数，它的值编码了在哪里可以找到实际数据的信息。前2个字节是数据的大小。最后2个字节是从数据开始的跟踪点结构开始的偏移量。
     * -- https://github.com/iovisor/bpftrace/pull/1542
     */
    __u32 msg; // __data_loc char[] msg;
};

//cat /sys/kernel/debug/tracing/events/netlink/netlink_extack/format

SEC("fentry/netlink_extack")
int BPF_PROG(fentry_netlink_extack, struct netlink_extack_error_ctx *nl_ctx) {
    bpf_printk("tcpconn, fentry_netlink_extack\n");

    /*
     * BPF_CORE_READ() is not dedicated/专用 to user-defined struct.
     */

    __u32 msg;
    bpf_probe_read(&msg, sizeof(msg), &nl_ctx->msg);
    char *c = (void *)(__u64)((void *)nl_ctx + (__u64)(msg & 0xFFFF));

    __output_msg(ctx, c, PROBE_TYPE_FENTRY, 0);

    return 0;
}

SEC("fexit/netlink_extack")
int BPF_PROG(fexit_netlink_extack, struct netlink_extack_error_ctx *nl_ctx, int retval) {
    bpf_printk("tcpconn, fexit_netlink_extack\n");

    __u32 msg;
    bpf_probe_read(&msg, sizeof(msg), &nl_ctx->msg);
    char *c = (void *)(__u64)((void *)nl_ctx + (__u64)(msg & 0xFFFF));

    __output_msg(ctx, c, PROBE_TYPE_FEXIT, retval);

    return 0;
}