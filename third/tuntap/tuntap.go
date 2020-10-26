package tuntap

import (
	"encoding/json"
	"fmt"
	"github.com/songgao/water"
	"log"
	"os"
	"os/exec"
	"os/signal"
	"runtime"
	"strconv"
	"strings"
	"syscall"
)
//ref:https://juejin.im/post/6844904119837147149#heading-1

//创建虚拟网卡并监听该文件描述符(读取流经该网卡的ip报文段)
func RcvTun()  {
	sig := make(chan os.Signal)

	go func() {
		/**
		http://tuntaposx.sourceforge.net/
		once an application opens the character device, say /dev/tap0, a virtual network interface is created in the system,
		which will be named accordingly, i.e. tap0. The network interface can be assigned addresses just like any other network interfaces
		一旦应用程序打开了tunX类的虚拟网卡,该设备会自动创建虚拟网卡接口
		*/

		tun_fd,err := os.OpenFile("/dev/tun11",os.O_RDONLY,0666)
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		shell := [][]string{
			//启动tun11虚拟网卡,该网卡的路由规则为: gateway:192.168.7.1,destination: 192.168.7.2
			//解释:目的地为192.168.7.2 地址会经过网卡tun11 并走向网关192.168.7.1
			{"/sbin/ifconfig","tun11 192.168.7.1 192.168.7.2 up"},

			//添加路由规则,destination: 180.101.49.0/24 Gateway: 192.168.7.2
			{"/sbin/route","-n add -net 180.101.49.0 -netmask 255.255.255.0 192.168.7.2"},
		}
		//经过上述网卡以及路由规则 ipv4=180.101.49/24 网段地址流向为 180.101.49/24 -> tun11 -> 192.168.7.2 -> tun11 -> 192.168.7.1(这里也就是应用程序读取/dev/tun11的位置)
		for _,c := range shell {
			cmd        := exec.Command(c[0],strings.Split(c[1]," ")...)
			output,err := cmd.CombinedOutput()
			if err != nil {
				panic(err.Error())
			}
			fmt.Println(string(output),"..ok!")
		}

		for {
			b     := make([]byte,2048)
			n,err := tun_fd.Read(b)
			if err != nil {
				sig <- syscall.SIGTERM
				return
			}

			//十进制ip+icmp报文段
			fmt.Println("got ipv4+icmp(ten):",b[:n])

			//二进制ip+icmp
			ipIcmp := fmt.Sprintf("%08b",b[:n])
			fmt.Println("got ipv4+icmp(binary):",ipIcmp)

			//二进制ip header报文段(20 bytes)
			binStr := ""
			for _, v1 := range b[:20] {
				//byte
				binStr += fmt.Sprintf("%08b",v1)
			}
			v,_ := json.Marshal(IPv4BinToStr(binStr))
			fmt.Println("got ipv4Header struct(20bytes):",string(v))

			//读取到的ip+icmp报文段
			//[01000101 00000000 00000000 01010100 01111011 01000100 00000000 00000000 01000000 00000001 01010010 01010110 11000000 10101000 00000111 00000001 10110100 01100101 00110001 00000000 00001000 00000000 11111101 10110000 10100010 11101101 00000000 00000000 01011111 10001001 01001001 11101010 00000000 00000110 11000010 11100100 00001000 00001001 00001010 00001011 00001100 00001101 00001110 00001111 00010000 00010001 00010010 00010011 00010100 00010101 00010110 00010111 00011000 00011001 00011010 00011011 00011100 00011101 00011110 00011111 00100000 00100001 00100010 00100011 00100100 00100101 00100110 00100111 00101000 00101001 00101010 00101011 00101100 00101101 00101110 00101111 00110000 00110001 00110010 00110011 00110100 00110101 00110110 00110111]


			//往tun虚拟网卡描述符中写入
			/*n,err = tun_fd.Write([]byte{})
			if err != nil {
				sig <- syscall.SIGTERM
				return
			}
			fmt.Println("write",n,"bytes success!")*/
		}
	}()


	//kill or ctrl + c
	signal.Notify(sig,syscall.SIGTERM,syscall.SIGINT)
	s := <-sig
	fmt.Println("get signal:",s," exit!")
}

func RcvTunFromWater()  {
	waterCfg := water.Config{
		DeviceType:water.TUN,
	}
	ifce,err := water.New(waterCfg)
	if err != nil {
		log.Fatal("tun err:",err)
	}
	log.Printf("Interface name: %s , OS:%s, \n",ifce.Name(),runtime.GOOS)

	packet := make([]byte,2000)
	for {
		n,err := ifce.Read(packet)
		if err != nil {
			log.Fatal(err)
		}
		log.Println("Packet Received:",packet[:n])
	}
}

//ipv4报文段转ip结构体
func IpTest()  {
	ipv4 := "0100010100000000000000000101010001111011010001000000000000000000010000000000000101010010010101101100000010101000000001110000000110110100011001010011000100000000"
	v,_  := json.Marshal(IPv4BinToStr(ipv4))
	fmt.Println(string(v))
	fmt.Println(len(ipv4))
}

//tcpdump 抓取的ipv4报文段(16进制)读取到ip结构体中
func Ipv4Test()  {
	IPv4Binary16Str := "4510003ca5da4000400696cf7f0000017f000001"
	Ipv4BinaryStr   := IPv416To2(IPv4Binary16Str)

	//src
	ip := IPv4BinToStr(Ipv4BinaryStr)
	v,_ := json.Marshal(ip)
	fmt.Println(string(v))
}

//16进制str 转 2进制str
func IPv416To2(IPv4Binary16Str string) string {
	IPv4BinaryStr := ""
	start         := 0
	end           := 7
	for {

		if start >= 24 {

		}
		//32bit
		str := IPv4Binary16Str[start:end+1]
		for i := 0; i < len(IPv4Binary16Str[start:end+1])/2 ; i++ {
			v,_ := strconv.ParseUint(str[i*2:i*2+2],16,10)
			IPv4BinaryStr  += fmt.Sprintf("%08b",v)
		}

		start = end + 1
		end   = start + 8 - 1
		if start >= 40 {
			break
		}
	}

	return IPv4BinaryStr
}

//ip数据结构
type Ip struct {
	Ver            int
	HeadLen        int
	Tos            int
	TotalLen       int
	Identification int
	Flag           string
	FlagOffset     int
	TTL            int
	Protocol       int
	CRC            string
	Src            string
	Dst            string
}

//Ipv4 2进制str数据转 struct Ip
func IPv4BinToStr(IPv4BinaryStr string) *Ip {
	if len(IPv4BinaryStr) < 5*32 {
		return nil
	}


	ip := Ip{}

	for i := 0; i < len(IPv4BinaryStr)/32; i++ {
		currentLeve32Bit := IPv4BinaryStr[i*32:i*32+32]
		switch i {
		case 0:
			tmp,_      := strconv.ParseInt(currentLeve32Bit[0:4],2,10)
			ip.Ver     = int(tmp)
			tmp,_      = strconv.ParseInt(currentLeve32Bit[4:8],2,10)
			ip.HeadLen = int(tmp)
			tmp,_      = strconv.ParseInt(currentLeve32Bit[8:16],2,10)
			ip.Tos     = int(tmp)
			tmp,_      = strconv.ParseInt(currentLeve32Bit[16:32],2,10)
			ip.TotalLen= int(tmp)
		case 1:
			tmp,_        := strconv.ParseInt(currentLeve32Bit[0:16],2,10)
			ip.Identification= int(tmp)
			ip.Flag       = currentLeve32Bit[16:19]
			tmp,_         = strconv.ParseInt(currentLeve32Bit[19:32],2,10)
			ip.FlagOffset = int(tmp)
		case 2:
			tmp,_       := strconv.ParseInt(currentLeve32Bit[0:8],2,10)
			ip.TTL       = int(tmp)
			tmp,_        = strconv.ParseInt(currentLeve32Bit[8:16],2,10)
			ip.Protocol  = int(tmp)
			ip.CRC       = currentLeve32Bit[16:32]
		case 3,4:
			Str := []string{}
			for j := 0; j < len(currentLeve32Bit)/8; j++ {
				v,_ := strconv.ParseInt(currentLeve32Bit[j*8:j*8 + 8],2,10)
				Str = append(Str,strconv.Itoa(int(v)))
			}
			if i == 3 {
				ip.Src = strings.Join(Str,".")
			} else {
				ip.Dst = strings.Join(Str,".")
			}
		}
	}


	return &ip
}