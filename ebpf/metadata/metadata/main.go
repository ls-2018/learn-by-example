package main

import (
	"github.com/cilium/ebpf"
	"github.com/cilium/ebpf/link"
	"github.com/cilium/ebpf/rlimit"
	"github.com/iovisor/gobpf/pkg/bpffs"
	"github.com/spf13/cobra"
	"github.com/vishvananda/netlink"
	"log"
	"net"
	"os"
	"path/filepath"
)

const bpffsPath = "bpffs"

//go:generate bpf2go -cc clang xdptailcall ./xdp_tailcall.c -- -D__TARGET_ARCH_x86 -I../../../headers -Wall
//go:generate bpf2go -cc clang tcmd ./tc_metadata.c -- -D__TARGET_ARCH_x86 -I../../../headers -Wall

var flags struct {
	device string
}

var rootCmd = cobra.Command{
	Use: "xdpmetadata",
}

func init() {
	rootCmd.Run = func(cmd *cobra.Command, args []string) {
		runXDPtailcall()
	}

	flag := rootCmd.PersistentFlags()
	flag.StringVar(&flags.device, "dev", "", "device to run XDP")
}

func runXDPtailcall() {
	if err := rlimit.RemoveMemlock(); err != nil {
		log.Fatalf("Failed to remove rlimit memlock: %v", err)
	}

	ifiDev, err := net.InterfaceByName(flags.device)
	if err != nil {
		log.Fatalf("Failed to fetch device info of %s: %v", flags.device, err)
	}
	var obj xdptailcallObjects
	if err := loadXdptailcallObjects(&obj, nil); err != nil {
		log.Fatalf("Failed to load xdp_tailcall bpf obj: %v", err)
	}
	defer obj.Close()

	if err := obj.XdpProgs.Put(uint32(0), obj.XdpFn); err != nil {
		log.Fatalf("Failed to save xdp_fn to xdp_progs: %v", err)
	}

	mapPinPath := filepath.Join(bpffsPath, "xdp_progs")
	if err := obj.XdpProgs.Pin(mapPinPath); err != nil {
		log.Fatalf("Failed to pin xdp_pros to %s: %v", "xdp_progs", err)
	}

	xdp, err := link.AttachXDP(link.XDPOptions{
		Program:   obj.XdpTailcall,
		Interface: ifiDev.Index,
		Flags:     link.XDPGenericMode,
	})
	if err != nil {
		log.Fatalf("Failed to attach xdp_tailcall to %s: %v", flags.device, err)
	}
	defer xdp.Close()

	devPinPath := filepath.Join(bpffsPath, flags.device)
	if err := xdp.Pin(devPinPath); err != nil {
		log.Fatalf("Failed to pin xdp_tailcall to %s: %v", flags.device, err)
	}

	log.Printf("xdp_tailcall is running on %s\n", flags.device)

	var _obj tcmdObjects
	if err := loadTcmdObjects(&_obj, nil); err != nil {
		log.Fatalf("Failed to load xdp_tailcall bpf obj: %v", err)
	}
	defer _obj.Close()

	ifi, err := netlink.LinkByName(flags.device)
	if err != nil {
		log.Fatalf("Failed to fetch link info of %s: %v", flags.device, err)
	}

	_tc, err := link.AttachTCX(link.TCXOptions{
		Interface: ifi.Attrs().Index,
		Program:   _obj.TcMetadata,
		Attach:    ebpf.AttachTCXIngress,
	})
	if err != nil {
		log.Fatalf("Failed to attach xdp_tailcall to %s: %v", flags.device, err)
	}
	defer _tc.Close()

	tcPinPath := filepath.Join(bpffsPath, "tc")
	if err := _tc.Pin(tcPinPath); err != nil {
		log.Fatalf("Failed to pin xdp_tailcall to %s: %v", flags.device, err)
	}

	log.Printf("tc is running on %s\n", flags.device)

}

func checkBpffs() {
	_ = os.Mkdir(bpffsPath, 0o700)
	mounted, _ := bpffs.IsMountedAt(bpffsPath)
	if mounted {
		return
	}

	err := bpffs.MountAt(bpffsPath)
	if err != nil {
		log.Fatalf("Failed to mount -t bpf %s: %v", bpffsPath, err)
	}
}

func main() {
	checkBpffs()
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
