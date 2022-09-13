﻿CEA#sh run
CEA#sh running-config 
Building configuration...

Current configuration : 1877 bytes
!
! Last configuration change at 05:57:45 UTC Wed Jul 19 2017
!
version 15.2
service timestamps debug datetime msec
service timestamps log datetime msec
!
hostname CEA
!
boot-start-marker
boot-end-marker
!
!
!
no aaa new-model
!
!
!
!
ip dhcp pool LAN
 network 192.168.1.0 255.255.255.0
 default-router 192.168.1.254 
 dns-server 8.8.8.8 
!
!
!
ip cef
no ipv6 cef
!
!
multilink bundle-name authenticated
!
!
!
!
!
!
!
username cisco privilege 15 password 0 cisco
!
!
!
!
!
!         
!
!
!
!
!
!
!
!
interface FastEthernet0/0
 no ip address
 shutdown
 duplex full
!
interface Ethernet1/0
 description LAN
 ip address 192.168.1.254 255.255.255.0
 ip nat inside
 duplex full
!
interface Ethernet1/1
 description INET
 ip address 58.113.1.1 255.255.255.252
 ip nat outside
 duplex full
!
interface Ethernet1/2
 description MPLS-VPN
 ip address 10.254.100.1 255.255.255.252
 duplex full
!
interface Ethernet1/3
 no ip address
 shutdown
 duplex full
!
interface Ethernet1/4
 no ip address
 shutdown
 duplex full
!
interface Ethernet1/5
 no ip address
 shutdown
 duplex full
!
interface Ethernet1/6
 no ip address
 shutdown
 duplex full
!
interface Ethernet1/7
 no ip address
 shutdown
 duplex full
!
router bgp 65122
 bgp log-neighbor-changes
 neighbor 10.254.100.2 remote-as 65511
 !
 address-family ipv4
  network 192.168.1.0
  neighbor 10.254.100.2 activate
 exit-address-family
!
ip nat inside source route-map NAT interface Ethernet1/1 overload
ip forward-protocol nd
!
!
ip http server
ip http authentication local
no ip http secure-server
ip route 0.0.0.0 0.0.0.0 58.113.1.2
!
ip access-list extended LAN
 permit ip 192.168.1.0 0.0.0.255 any
!
no cdp advertise-v2
!
route-map NAT permit 10
 match ip address LAN
 match interface Ethernet1/1
!
!
!
control-plane
!
!
line con 0
 stopbits 1
line aux 0
 stopbits 1
line vty 0 4
 password cisco
 login local
 transport input all
!
!
end

