package f5

import (
	"encoding/json"
	"strings"
)

type LBVirtualPolicy struct {
	Name      string `json:"name"`
	Partition string `json:"partition"`
	FullPath  string `json:"fullPath"`
}

type LBVirtualPoliciesRef struct {
	Items []LBVirtualPolicy `json":items"`
}

type LBVirtualProfile struct {
	Name      string `json:"name"`
	Partition string `json:"partition"`
	FullPath  string `json:"fullPath"`
	Context   string `json:"context"`
}

type LBVirtualPersistProfile struct {
	Name      string `json:"name"`
	Partition string `json:"partition"`
	TmDefault string `json:"tmDefault"`
}

type LBVirtualProfileRef struct {
	Items []LBVirtualProfile `json":items"`
}

type LBVirtual struct {
	Name             string                    `json:"name"`
	FullPath         string                    `json:"fullPath"`
	Partition        string                    `json:"partition"`
	Destination      string                    `json:"destination"`
	Pool             string                    `json:"pool"`
	AddressStatus    string                    `json:"addressStatus"`
	AutoLastHop      string                    `json:"autoLasthop"`
	CmpEnabled       string                    `json:"cmpEnabled"`
	ConnectionLimit  int                       `json:"connectionLimit"`
	Enabled          bool                      `json:"enabled"`
	IpProtocol       string                    `json:"ipProtocol"`
	Source           string                    `json:"source"`
	SourcePort       string                    `json:"sourcePort"`
	SynCookieStatus  string                    `json:"synCookieStatus"`
	TranslateAddress string                    `json:"translateAddress"`
	TranslatePort    string                    `json:"translatePort"`
	Profiles         LBVirtualProfileRef       `json:"profilesReference"`
	Policies         LBVirtualPoliciesRef      `json:"policiesReference"`
	Rules            []string                  `json:"rules"`
	Persist          []LBVirtualPersistProfile `json:"persist"`
}

type LBVirtuals struct {
	Items []LBVirtual
}

type LBVirtualStatsDescription struct {
	Description string `json:"description"`
}

type LBVirtualStatsValue struct {
	Value int `json:"value"`
}

type LBVirtualStatsInnerEntries struct {
	Clientside_bitsIn             LBVirtualStatsValue       `json:"clientside.bitsIn"`
	Clientside_bitsOut            LBVirtualStatsValue       `json:"clientside.bitsOut"`
	Clientside_curConns           LBVirtualStatsValue       `json:"clientside.curConns"`
	Clientside_evictedConns       LBVirtualStatsValue       `json:"clientside.evictedConns"`
	Clientside_maxConns           LBVirtualStatsValue       `json:"clientside.maxConns"`
	Clientside_pktsIn             LBVirtualStatsValue       `json:"clientside.pktsIn"`
	Clientside_pktsOut            LBVirtualStatsValue       `json:"clientside.pktsOut"`
	Clientside_slowKilled         LBVirtualStatsValue       `json:"clientside.slowKilled"`
	Clientside_totConns           LBVirtualStatsValue       `json:"clientside.totConns"`
	CmpEnableMode                 LBVirtualStatsDescription `json:"cmpEnableMode"`
	CmpEnabled                    LBVirtualStatsDescription `json:"cmpEnabled"`
	CsMaxConnDur                  LBVirtualStatsValue       `json:"csMaxConnDur"`
	CsMeanConnDur                 LBVirtualStatsValue       `json:"csMeanConnDur"`
	CsMinConnDur                  LBVirtualStatsValue       `json:"csMinConnDur"`
	Destination                   LBVirtualStatsDescription `json:"destination"`
	Ephemeral_bitsIn              LBVirtualStatsValue       `json:"ephemeral.bitsIn"`
	Ephemeral_bitsOut             LBVirtualStatsValue       `json:"ephemeral.bitsOut"`
	Ephemeral_curConns            LBVirtualStatsValue       `json:"ephemeral.curConns"`
	Ephemeral_evictedConns        LBVirtualStatsValue       `json:"ephemeral.evictedConns"`
	Ephemeral_maxConns            LBVirtualStatsValue       `json:"ephemeral.maxConns"`
	Ephemeral_pktsIn              LBVirtualStatsValue       `json:"ephemeral.pktsIn"`
	Ephemeral_pktsOut             LBVirtualStatsValue       `json:"ephemeral.pktsOut"`
	Ephemeral_slowKilled          LBVirtualStatsValue       `json:"ephemeral.slowKilled"`
	Ephemeral_totConns            LBVirtualStatsValue       `json:"ephemeral.totConns"`
	FiveMinAvgUsageRatio          LBVirtualStatsValue       `json:"fiveMinAvgUsageRatio"`
	FiveSecAvgUsageRatio          LBVirtualStatsValue       `json:"fiveSecAvgUsageRatio"`
	TmName                        LBVirtualStatsDescription `json:"tmName"`
	OneMinAvgUsageRatio           LBVirtualStatsValue       `json:"oneMinAvgUsageRatio"`
	Status_availabilityState      LBVirtualStatsDescription `json:"status.availabilityState"`
	Status_enabledState           LBVirtualStatsDescription `json:"status.enabledState"`
	Status_statusReason           LBVirtualStatsDescription `json:"status.statusReason"`
	SyncookieStatus               LBVirtualStatsDescription `json:"syncookieStatus"`
	Syncookie_accepts             LBVirtualStatsValue       `json:"syncookie.accepts"`
	Syncookie_hwAccepts           LBVirtualStatsValue       `json:"syncookie.hwAccepts"`
	Syncookie_hwSyncookies        LBVirtualStatsValue       `json:"syncookie.hwSyncookies"`
	Syncookie_hwsyncookieInstance LBVirtualStatsValue       `json:"syncookie.hwsyncookieInstance"`
	Syncookie_rejects             LBVirtualStatsValue       `json:"syncookie.rejects"`
	Syncookie_swsyncookieInstance LBVirtualStatsValue       `json:"syncookie.swsyncookieInstance"`
	Syncookie_syncacheCurr        LBVirtualStatsValue       `json:"syncookie.syncacheCurr"`
	Syncookie_syncacheOver        LBVirtualStatsValue       `json:"syncookie.syncacheOver"`
	Syncookie_syncookies          LBVirtualStatsValue       `json:"syncookie.syncookies"`
	TotRequests                   LBVirtualStatsValue       `json:"totRequests"`
}

type LBVirtualStatsNestedStats struct {
	Kind     string                     `json:"kind"`
	SelfLink string                     `json:"selfLink"`
	Entries  LBVirtualStatsInnerEntries `json:"entries"`
}

type LBVirtualURLKey struct {
	NestedStats LBVirtualStatsNestedStats `json:"nestedStats"`
}
type LBVirtualStatsOuterEntries map[string]LBVirtualURLKey

type LBVirtualStats struct {
	Kind       string                     `json:"kind"`
	Generation int                        `json:"generation"`
	SelfLink   string                     `json:"selfLink"`
	Entries    LBVirtualStatsOuterEntries `json:"entries"`
}

func (f *Device) ShowVirtuals() (error, *LBVirtuals) {

	u := f.Proto + "://" + f.Hostname + "/mgmt/tm/ltm/virtual"
	res := LBVirtuals{}

	err, _ := f.sendRequest(u, GET, nil, &res)
	if err != nil {
		return err, nil
	} else {
		return nil, &res
	}

}

func (f *Device) ShowVirtual(vname string) (error, *LBVirtual) {

	vname = strings.Replace(vname, "/", "~", -1)
	u := f.Proto + "://" + f.Hostname + "/mgmt/tm/ltm/virtual/" + vname + "?expandSubcollections=true"
	res := LBVirtual{}

	err, _ := f.sendRequest(u, GET, nil, &res)
	if err != nil {
		return err, nil
	} else {
		return nil, &res
	}

}

func (f *Device) ShowVirtualStats(vname string) (error, *LBVirtualStats) {

	vname = strings.Replace(vname, "/", "~", -1)
	u := f.Proto + "://" + f.Hostname + "/mgmt/tm/ltm/virtual/" + vname + "/stats"
	res := LBVirtualStats{}

	err, _ := f.sendRequest(u, GET, nil, &res)
	if err != nil {
		return err, nil
	} else {
		return nil, &res
	}
}

func (f *Device) AddVirtual(virt *json.RawMessage) (error, *LBVirtual) {

	u := f.Proto + "://" + f.Hostname + "/mgmt/tm/ltm/virtual"
	res := LBVirtual{}

	// post the request
	err, _ := f.sendRequest(u, POST, virt, &res)
	if err != nil {
		return err, nil
	} else {
		return nil, &res
	}

}

func (f *Device) UpdateVirtual(vname string, body *json.RawMessage) (error, *LBVirtual) {

	vname = strings.Replace(vname, "/", "~", -1)
	u := f.Proto + "://" + f.Hostname + "/mgmt/tm/ltm/virtual/" + vname
	res := LBVirtual{}

	// put the request
	err, _ := f.sendRequest(u, PUT, &body, &res)
	if err != nil {
		return err, nil
	} else {
		return nil, &res
	}

}

func (f *Device) DeleteVirtual(vname string) (error, *Response) {

	vname = strings.Replace(vname, "/", "~", -1)
	u := f.Proto + "://" + f.Hostname + "/mgmt/tm/ltm/virtual/" + vname
	res := json.RawMessage{}

	err, resp := f.sendRequest(u, DELETE, nil, &res)
	if err != nil {
		return err, nil
	} else {
		return nil, resp
	}

}
