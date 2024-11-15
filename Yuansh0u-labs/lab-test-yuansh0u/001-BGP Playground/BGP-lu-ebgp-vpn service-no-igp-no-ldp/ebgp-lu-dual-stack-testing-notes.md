```bash
R5#ping 2001::1 sou lo0
Type escape sequence to abort.
Sending 5, 100-byte ICMP Echos to 2001::1, timeout is 2 seconds:
Packet sent with a source address of 2001::5
!!!!!
Success rate is 100 percent (5/5), round-trip min/avg/max = 44/50/60 ms
R5#ping vrf a fc00::1 sou lo1
Type escape sequence to abort.
Sending 5, 100-byte ICMP Echos to FC00::1, timeout is 2 seconds:
Packet sent with a source address of FC00::5%a
!!!!!
Success rate is 100 percent (5/5), round-trip min/avg/max = 36/48/56 ms

R5#ping vrf a 10.0.0.1 sou lo1
Type escape sequence to abort.
Sending 5, 100-byte ICMP Echos to 10.0.0.1, timeout is 2 seconds:
Packet sent with a source address of 10.0.0.5 
!!!!!
Success rate is 100 percent (5/5), round-trip min/avg/max = 36/63/80 ms
R5#
R5#traceroute vrf a 10.0.0.1 sou lo1
Type escape sequence to abort.
Tracing the route to 10.0.0.1
VRF info: (vrf in name/id, vrf out name/id)
  1 35.1.1.3 [MPLS: Labels 21/20 Exp 0] 84 msec 32 msec 52 msec
  2 34.1.1.4 [MPLS: Label 20 Exp 0] 52 msec 48 msec 52 msec
  3 10.0.0.1 80 msec 76 msec 48 msec
R5#
```

```bash

R5#show bgp label
   Network          Next Hop      In label/Out label
   1.1.1.1/32       35.1.1.3        nolabel/20
   2.2.2.2/32       35.1.1.3        nolabel/18   <v6 service>
   3.3.3.3/32       35.1.1.3        nolabel/imp-null
   4.4.4.4/32       35.1.1.3        nolabel/21   <v4 service>

   5.5.5.5/32       0.0.0.0         imp-null/nolabel

R3#show bgp label
   Network          Next Hop      In label/Out label
   1.1.1.1/32       23.1.1.2        20/19
                    34.1.1.4        20/17
   2.2.2.2/32       23.1.1.2        18/imp-null
                    34.1.1.4        18/18
   3.3.3.3/32       0.0.0.0         imp-null/nolabel
   4.4.4.4/32       23.1.1.2        21/27
                    34.1.1.4        21/imp-null
   5.5.5.5/32       35.1.1.5        22/imp-null

 R4#show bgp label
   Network          Next Hop      In label/Out label
   1.1.1.1/32       14.1.1.1        17/imp-null
   2.2.2.2/32       34.1.1.3        18/18
                    14.1.1.1        18/18
   3.3.3.3/32       34.1.1.3        19/imp-null
   4.4.4.4/32       0.0.0.0         imp-null/nolabel
   5.5.5.5/32       34.1.1.3        23/22

R4#  

R1#show bgp label
   Network          Next Hop      In label/Out label
   1.1.1.1/32       0.0.0.0         imp-null/nolabel
   2.2.2.2/32       12.1.1.2        18/imp-null
   3.3.3.3/32       12.1.1.2        21/20
                    14.1.1.4        21/19
   4.4.4.4/32       14.1.1.4        20/imp-null
   5.5.5.5/32       12.1.1.2        22/26
                    14.1.1.4        22/23

R1#

R2#show bgp label
   Network          Next Hop      In label/Out label
   1.1.1.1/32       23.1.1.3        19/20
                    12.1.1.1        19/imp-null
   2.2.2.2/32       0.0.0.0         imp-null/nolabel
   3.3.3.3/32       23.1.1.3        20/imp-null
                    12.1.1.1        20/21
   4.4.4.4/32       23.1.1.3        27/21
                    12.1.1.1        27/20
   5.5.5.5/32       23.1.1.3        26/22
                    12.1.1.1        26/22

R2#
```
```bash
 R5#show bgp vpnv4 un all 
BGP table version is 13, local router ID is 5.5.5.5
Status codes: s suppressed, d damped, h history, * valid, > best, i - internal, 
              r RIB-failure, S Stale, m multipath, b backup-path, f RT-Filter, 
              x best-external, a additional-path, c RIB-compressed, 
Origin codes: i - IGP, e - EGP, ? - incomplete
RPKI validation codes: V valid, I invalid, N Not found

     Network          Next Hop            Metric LocPrf Weight Path
Route Distinguisher: 1:1 (default for vrf a)
 *>  10.0.0.1/32      4.4.4.4                  0             0 1000 ?
 *>  10.0.0.2/32      4.4.4.4                  0             0 1000 ?
 *>  10.0.0.3/32      4.4.4.4                  0             0 1000 ?
 *>  10.0.0.4/32      4.4.4.4                  0             0 1000 4 ?
 *>  10.0.0.5/32      0.0.0.0                  0         32768 ?
R5#show bgp vpnv6 un all 
BGP table version is 15, local router ID is 5.5.5.5
Status codes: s suppressed, d damped, h history, * valid, > best, i - internal, 
              r RIB-failure, S Stale, m multipath, b backup-path, f RT-Filter, 
              x best-external, a additional-path, c RIB-compressed, 
Origin codes: i - IGP, e - EGP, ? - incomplete
RPKI validation codes: V valid, I invalid, N Not found

     Network          Next Hop            Metric LocPrf Weight Path
Route Distinguisher: 1:1 (default for vrf a)
 *>  FC00::1/128      ::FFFF:2.2.2.2           0             0 1000 1 ?
 *>  FC00::2/128      ::FFFF:2.2.2.2           0             0 1000 2 ?
 *>  FC00::3/128      ::FFFF:2.2.2.2           0             0 1000 3 ?
 *   FC00::5/128      ::FFFF:2.2.2.2           0             0 1000 5 ?
 *>                   ::                       0         32768 ?
R5#
R5#show bgp ipv6 unicast 
BGP table version is 8, local router ID is 5.5.5.5
Status codes: s suppressed, d damped, h history, * valid, > best, i - internal, 
              r RIB-failure, S Stale, m multipath, b backup-path, f RT-Filter, 
              x best-external, a additional-path, c RIB-compressed, 
Origin codes: i - IGP, e - EGP, ? - incomplete
RPKI validation codes: V valid, I invalid, N Not found

     Network          Next Hop            Metric LocPrf Weight Path
 *>  2001::1/128      ::FFFF:2.2.2.2           0             0 1000 1 ?
 *>  2001::3/128      ::FFFF:2.2.2.2           0             0 1000 3 ?
 *   2001::5/128      ::FFFF:2.2.2.2           0             0 1000 5 ?
 *>                   ::                       0         32768 ?
R5#
```