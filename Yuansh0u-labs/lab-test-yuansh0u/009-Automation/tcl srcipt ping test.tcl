
##// Add new motherfucker folder on flash:

R1#mkdir flash:motherfucker
R1#dir flash:
Directory of bootflash:/
92744  drwx             4096  May 29 2017 02:10:29 +00:00  motherfucker

1494446080 bytes total (518746112 bytes free)
R1#   

R1#rmdir flash:motherfucker
Remove directory filename [motherfucker]? 
Delete bootflash:motherfucker? [confirm]
Removed dir bootflash:motherfucker

##

tclsh
## STANDARD PING OPERATION
 foreach pingloopbacks {
 	2.2.2.2
 	2.2.2.3
 	2.2.2.4
 	2.2.2.5
 } {ping $pingloopbacks repeat 1 size 1300 source loopback 0}
write memory
tclquit
##
tclsh
## Write iosconfig into flash:motherfucker/motherfucker.tcl
puts [ open "flash:motherfucker/motherfucker.tcl" w+ ] {
## Set up basic configuration...
 ios_config "hostname R1"
 ios_config "interface loopback110" "ip address 123.123.123.123 255.255.255.255" "ip ospf 1 area 0" "description loopback0"
 ios_config "router ospf 1" "router-id 123.123.123.123" "do clear ip ospf process" 
## If tap "do clear ip ospf process" and Reset ALL OSPF processes? [no]:
## Input "yes" to finish this command.  
 typeahead "yes"
 ios_config "do clear ip ospf process" 
 ping 2.2.2.2 source loopback 0 repeat 10
 ping 2.2.2.3 source loopback 0 repeat 10
	write memory
}
tclquit	

tclsh flash:/motherfucker/motherfucker.tcl  
##
tclsh
puts [ open "flash:motherfucker/DEATH.tcl" w+ ] {
## If tap "do debug all" and prompt messages is comming up.
## This may severely impact network performance. Continue? (yes/[no]):
## Input "yes" to finish this command.   
   typeahead "yes"
   ios_config "do debug all"
   write memory
## If tap "reload " and Proceed with reload? [confirm]
## Input "implicit enter" to finish this command. \n MEANS implicit enter!!  
 typeahead "\n"
 reload
}
tclquit	

tclsh flash:/motherfucker/DEATH.tcl
##
motherfucker#show flash: | sec tcl
245        274 May 29 2017 02:25:43 +00:00 /bootflash/motherfucker/motherfucker.tcl
246         59 May 29 2017 02:28:19 +00:00 /bootflash/motherfucker/DEATH.tcl
motherfucker#
motherfucker#more flash:motherfucker/DEATH.tcl
 
 typeahead "yes"
 ios_config "debug all"
 
 write memory

motherfucker#tclsh flash:/motherfucker/DEATH.tclsh 

##