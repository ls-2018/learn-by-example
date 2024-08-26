//go:build ignore

#include "bpf_all.h"

#define CSUM_SIZE 64

SEC("xdp")
int xdp_ping(struct xdp_md *ctx) {
    void *data = ctx_ptr(ctx, data);
    void *data_end = ctx_ptr(ctx, data_end);

    struct ethhdr *eth;
    eth = (typeof(eth))data;
    if ((void *)(eth + 1) > data_end)
        return XDP_PASS;

    if (eth->h_proto != bpf_htons(ETH_P_IP))
        return XDP_PASS;

    struct iphdr *iph;
    iph = (typeof(iph))(eth + 1);
    if ((void *)(iph + 1) > data_end)
        return XDP_PASS;

    if (iph->protocol != IPPROTO_ICMP)
        return XDP_PASS;

    int csum_size = CSUM_SIZE;

    struct icmphdr *icmph;
    // icmph = (typeof(icmph))((void *)iph + (iph->ihl * 4)); // FAILED: R3 offset is outside of the packet
    icmph = (typeof(icmph))(iph + 1);
    if ((void *)(icmph) + csum_size > data_end)
        return XDP_PASS;

    if (icmph->type != ICMP_ECHO)
        return XDP_PASS;

    // FAILED: int csum_size = iph->tot_len - sizeof(*iph); // R3 offset is outside of the packet
    // FAILED: int csum_size = data_end - (void *)icmph;    // R4 unbounded memory access, use 'var &= const' or 'if (var < const)'

    // 正确的 icmp hdr
    icmph->type = ICMP_ECHOREPLY;
    icmph->checksum = 0; // Note: reset and then checksum
    icmph->checksum = ipv4_csum(icmph, csum_size);

    // 正确的 ip hdr
    __be32 daddr = iph->daddr;
    iph->daddr = iph->saddr;
    iph->saddr = daddr;
    iph->ttl = 64;
    iph->check = 0; // Note: reset and then checksum
    iph->check = ipv4_csum(iph, csum_size + sizeof(*iph));

    // correct eth hdr
    char dmac[ETH_ALEN];
    __builtin_memcpy(dmac, eth->h_dest, ETH_ALEN);
    __builtin_memcpy(eth->h_dest, eth->h_source, ETH_ALEN);
    __builtin_memcpy(eth->h_source, dmac, ETH_ALEN);

    bpf_printk("xdpping replay icmp echo reply\n");

    return XDP_TX;
}