package main

import (
	"fmt"
	"github.com/sparrc/go-ping"
	"math/rand"
)

var SERVERS = map[string]string{
	"US-East (Virginia)":         "dynamodb.us-east-1.amazonaws.com",
	"US East (Ohio)":             "dynamodb.us-east-2.amazonaws.com",
	"US-West (California)":       "dynamodb.us-west-1.amazonaws.com",
	"US-West (Oregon)":           "dynamodb.us-west-2.amazonaws.com",
	"Canada (Central)":           "dynamodb.ca-central-1.amazonaws.com",
	"Europe (Ireland)":           "dynamodb.eu-west-1.amazonaws.com",
	"Europe (London)":            "dynamodb.eu-west-2.amazonaws.com",
	"Europe (Frankfurt)":         "dynamodb.eu-central-1.amazonaws.com",
	"Europe (Paris)":             "dynamodb.eu-west-3.amazonaws.com",
	"Europe (Stockholm)":         "dynamodb.eu-north-1.amazonaws.com",
	"Asia Pacific (Mumbai)":      "dynamodb.ap-south-1.amazonaws.com",
	"Asia Pacific (Osaka-Local)": "dynamodb.ap-northeast-3.amazonaws.com",
	"Asia Pacific (Seoul)":       "dynamodb.ap-northeast-2.amazonaws.com",
	"Asia Pacific (Singapore)":   "dynamodb.ap-southeast-1.amazonaws.com",
	"Asia Pacific (Sydney)":      "dynamodb.ap-southeast-2.amazonaws.com",
	"Asia Pacific (Tokyo)":       "dynamodb.ap-northeast-1.amazonaws.com",
	"South America (São Paulo)":  "dynamodb.sa-east-1.amazonaws.com",
	"China (Beijing)":            "dynamodb.cn-north-1.amazonaws.com.cn",
	//"China (Ningxia)":            "dynamodb.cn-northwest-1.amazonaws.com.cn",
	"AWS GovCloud (US-East)":     "dynamodb.us-gov-east-1.amazonaws.com",
	"AWS GovCloud (US)":          "dynamodb.us-gov-west-1.amazonaws.com",
}

const SEED = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func main() {
	//var randomStr string

	for region,addr := range SERVERS {
		//randomStr = GetRandomString()
		pinger,err := ping.NewPinger(addr)
		if err != nil {
			fmt.Println("ping error:",err)
		}
		pinger.Count = 10
		//stats := pinger.Statistics()

		pinger.OnRecv = func(pkt *ping.Packet) {
			fmt.Printf("%d bytes from %s: icmp_seq=%d time=%v\n",
				pkt.Nbytes, pkt.IPAddr, pkt.Seq, pkt.Rtt)
		}

		pinger.OnFinish = func(stats *ping.Statistics) {
			fmt.Printf("--- %s ping 统计 ---\n", region)
			fmt.Printf("发送数量：%d , 接收数量：%d, 丢包数量：%v%%\n",
				stats.PacketsSent, stats.PacketsRecv, stats.PacketLoss)
			fmt.Printf("时长统计： min/avg/max/stddev = %v/%v/%v/%v\n\n",
				stats.MinRtt, stats.AvgRtt, stats.MaxRtt, stats.StdDevRtt)
		}

		fmt.Printf("%v: %s (%s):\n",region, pinger.Addr(), pinger.IPAddr())
		pinger.Run()
	}
}

func GetRandomString() string {
	random := make([]byte,6)
	max := len(SEED)
	for i := range random {
		random[i] = SEED[rand.Intn(max)]
	}
	return string(random)
}