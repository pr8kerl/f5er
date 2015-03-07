package main

import (
	"log"
)

type LBDeviceRef struct {
	Link  string          `json:"selfLink"`
	Items []LBDeviceState `json":items"`
}

type LBDeviceState struct {
	Name          string `json:"name"`
	Path          string `json:"fullPath"`
	FailoverState string `json:"failoverState"`
	ManagementIP  string `json:"managementIP"`
}

/*
{
  "kind":"tm:cm:device:devicecollectionstate",
  "selfLink":"https://localhost/mgmt/tm/cm/device?ver=11.6.0",
  "items":[{
    "kind":"tm:cm:device:devicestate",
    "name":"aumel-bigip-1.myob.com.au",
    "partition":"Common",
    "fullPath":"/Common/aumel-bigip-1.myob.com.au",
    "generation":1,
    "selfLink":"https://localhost/mgmt/tm/cm/device/~Common~aumel-bigip-1.myob.com.au?ver=11.6.0",
    "activeModules":["LTM, Base, 1600|TXGAYYB-YHAJFVP|LTM, Base|IPV6 Gateway|Rate Shaping|Ram Cache|50 MBPS COMPRESSION|SSL, 500 TPS Per Core|SSL, CMP|Anti-Virus Checks|Base Endpoint Security Checks|Firewall Checks|Network Access|Secure Virtual Keyboard|APM, Web Application|Machine Certificate Checks|Protected Workspace|Remote Desktop|App Tunnel|Application Acceleration Manager, Core","SSL, Max TPS|TVOMGMS-JRGJFKE"],
    "baseMac":"00:23:e9:51:30:40",
    "build":"0.0.401",
    "cert":"/Common/dtdi.crt",
    "chassisId":"f5-vvyi-zaet",
    "chassisType":"individual",
    "configsyncIp":"172.16.1.241",
    "edition":"Final",
    "failoverState":"standby",
    "haCapacity":0,
    "hostname":"aumel-bigip-1.myob.com.au",
    "key":"/Common/dtdi.key",
    "managementIp":"10.60.99.241",
    "marketingName":"BIG-IP 1600",
    "mirrorIp":"any6",
    "mirrorSecondaryIp":"any6",
    "multicastIp":"224.0.0.245",
    "multicastPort":62960,
    "optionalModules":["AFM, 1600","APM, Base CCU, 1600","APM, Max CCU, 1600","App Mode (TMSH Only, No Root/Bash)","ASM, 1600 Bundle","Client Authentication","Compression, Max MBPS, 1600","DNS Services","External Interface and Network HSM","Global Traffic Manager Module","IPI Subscription, 1Yr, 1600/2000/2200","IPI Subscription, 3Yr, 1600/2000/2200","Link Controller","LTM, GTM, ASM, APM, WAM, WOM (1600)","MSM","PSM","Routing Bundle","SDN Services","SSL, Forward Proxy","WBA","WBA, Bundle, 1600","WOM","WOM, Bundle, 1600"],
    "platformId":"C102",
    "product":"BIG-IP",
    "selfDevice":"true",
    "timeZone":"EST",
    "version":"11.6.0",
    "unicastAddress":[{"effectiveIp":"172.16.1.241","effectivePort":1026,"ip":"172.16.1.241","port":1026}]
}

*/

func showDevice() {

	u := "https://" + f5Host + "/mgmt/tm/cm/device"
	res := LBDeviceRef{}

	err, resp := SendRequest(u, GET, &sessn, nil, &res)
	if err != nil {
		log.Fatalf("%s : %s\n", resp.HttpResponse().Status, err)
	}
	printResponse(&res.Items)

}
