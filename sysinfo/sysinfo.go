package sysinfo

import (
	"fmt"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/net"
	"time"
)

func Info() {
	v, _ := mem.VirtualMemory()
	c, _ := cpu.Info()
	//cpu占用百分比
	cc, _ := cpu.Percent(time.Second, false)
	d, _ := disk.Usage("/")
	n, _ := host.Info()
	nv, _ := net.IOCounters(true)
	fmt.Printf("CPU: %v * %v Cores\t Used: %.2f%%\n", c[0].ModelName, len(c), cc[0])
	fmt.Printf("Mem: %vGB \t Free: %vGB \t Used: %vMB \t Usage: %.2f%%\n", v.Total/1024/1000/1000, v.Available/1024/1000/1000, v.Used/1024/1000, v.UsedPercent)
	fmt.Printf("Disk: %vGB \t Free: %vGB \t Usage:%.2f%%\n", d.Total/1024/1024/1024, d.Free/1024/1024/1024, d.UsedPercent)
	fmt.Printf("Network: %vbytes / %vbytes ( %vMB / %vMB )\n", nv[0].BytesRecv, nv[0].BytesSent, nv[0].BytesRecv/1024/1024, nv[0].BytesSent/1024/1024)
	fmt.Printf("Operating System: %v(%v) %v\n", n.Platform, n.PlatformFamily, n.PlatformVersion)
	fmt.Printf("Hostname: %v\n", n.Hostname)
}
