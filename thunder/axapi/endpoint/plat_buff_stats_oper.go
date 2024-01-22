

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type PlatBuffStatsOper struct {
    
    Oper PlatBuffStatsOperOper `json:"oper"`

}
type DataPlatBuffStatsOper struct {
    DtPlatBuffStatsOper PlatBuffStatsOper `json:"plat-buff-stats"`
}


type PlatBuffStatsOperOper struct {
    BuffThrData []PlatBuffStatsOperOperBuffThrData `json:"buff-thr-data"`
    TotalBufAState int `json:"total-buf-a-state"`
    TotalInappCp int `json:"total-inapp-cp"`
    TotalBIncacheCp int `json:"total-b-incache-cp"`
    TotalBIncache int `json:"total-b-incache"`
    TotalBInthrq int `json:"total-b-inthrq"`
    TotalBDcmsgQ int `json:"total-b-dcmsg-q"`
    TotalBMisc int `json:"total-b-misc"`
    TotalBDfree int `json:"total-b-dfree"`
    TotalFree int `json:"total-free"`
    TotalInhw int `json:"total-inhw"`
    G_num_thr_poll int `json:"g_num_thr_poll"`
    G_num_domains int `json:"g_num_domains"`
    ThreadsDomain0 []PlatBuffStatsOperOperThreadsDomain0 `json:"threads-domain0"`
    CapsulesDomain0 []PlatBuffStatsOperOperCapsulesDomain0 `json:"capsules-domain0"`
    ThreadsDomain1 []PlatBuffStatsOperOperThreadsDomain1 `json:"threads-domain1"`
    CapsulesDomain1 []PlatBuffStatsOperOperCapsulesDomain1 `json:"capsules-domain1"`
    GBuffPool []PlatBuffStatsOperOperGBuffPool `json:"g-buff-pool"`
    TotalNumBuffers int `json:"total-num-buffers"`
    DebugCount int `json:"debug-count"`
    DbgNcache int `json:"dbg-ncache"`
    DbgNcap int `json:"dbg-ncap"`
    DbgNpart int `json:"dbg-npart"`
    DbgNddom int `json:"dbg-nddom"`
    DbgNbuf int `json:"dbg-nbuf"`
    DbgBufTotalCachePfpga int `json:"dbg-buf-total-cache-pfpga"`
    DbgBufTotalCachePcpu int `json:"dbg-buf-total-cache-pcpu"`
    DbgBufInCachePerFpga []PlatBuffStatsOperOperDbgBufInCachePerFpga `json:"dbg-buf-in-cache-per-fpga"`
    DbgBufInCachePerCpu []PlatBuffStatsOperOperDbgBufInCachePerCpu `json:"dbg-buf-in-cache-per-cpu"`
}


type PlatBuffStatsOperOperBuffThrData struct {
    ThrNum int `json:"thr-num"`
    NumBuffCache int `json:"num-buff-cache"`
    BuffAppState int `json:"buff-app-state"`
    BuffAppInq int `json:"buff-app-inq"`
    BuffMisc int `json:"buff-misc"`
}


type PlatBuffStatsOperOperThreadsDomain0 struct {
    ThrNameBufStat string `json:"thr-name-buf-stat"`
}


type PlatBuffStatsOperOperCapsulesDomain0 struct {
    DomainNum int `json:"domain-num"`
    BufFpga []PlatBuffStatsOperOperCapsulesDomain0BufFpga `json:"buf-fpga"`
}


type PlatBuffStatsOperOperCapsulesDomain0BufFpga struct {
    BufFpgaNum int `json:"buf-fpga-num"`
    BufFpgaS []PlatBuffStatsOperOperCapsulesDomain0BufFpgaBufFpgaS `json:"buf-fpga-s"`
}


type PlatBuffStatsOperOperCapsulesDomain0BufFpgaBufFpgaS struct {
    BufStatFpga int `json:"buf-stat-fpga"`
}


type PlatBuffStatsOperOperThreadsDomain1 struct {
    ThrNameBufStat string `json:"thr-name-buf-stat"`
}


type PlatBuffStatsOperOperCapsulesDomain1 struct {
    DomainNum int `json:"domain-num"`
    BufFpga []PlatBuffStatsOperOperCapsulesDomain1BufFpga `json:"buf-fpga"`
}


type PlatBuffStatsOperOperCapsulesDomain1BufFpga struct {
    BufFpgaNum int `json:"buf-fpga-num"`
    BufFpgaS []PlatBuffStatsOperOperCapsulesDomain1BufFpgaBufFpgaS `json:"buf-fpga-s"`
}


type PlatBuffStatsOperOperCapsulesDomain1BufFpgaBufFpgaS struct {
    BufStatFpga int `json:"buf-stat-fpga"`
}


type PlatBuffStatsOperOperGBuffPool struct {
    GStatFpgaName string `json:"g-stat-fpga-name"`
    GStatGets0 string `json:"g-stat-gets0"`
    GStatPuts0 string `json:"g-stat-puts0"`
    GStatGets1 string `json:"g-stat-gets1"`
    GStatPuts1 string `json:"g-stat-puts1"`
}


type PlatBuffStatsOperOperDbgBufInCachePerFpga struct {
    DbgFpgaNum int `json:"dbg-fpga-num"`
    DbgFpgaNcache int `json:"dbg-fpga-ncache"`
    DbgFpgaNcap int `json:"dbg-fpga-ncap"`
    DbgFpgaNpart int `json:"dbg-fpga-npart"`
    DbgFpgaNddom int `json:"dbg-fpga-nddom"`
}


type PlatBuffStatsOperOperDbgBufInCachePerCpu struct {
    DbgCpuNum int `json:"dbg-cpu-num"`
    DbgCpuNcache int `json:"dbg-cpu-ncache"`
    DbgCpuNcap int `json:"dbg-cpu-ncap"`
    DbgCpuNpart int `json:"dbg-cpu-npart"`
    DbgCpuNddom int `json:"dbg-cpu-nddom"`
}

func (p *PlatBuffStatsOper) GetId() string{
    return "1"
}

func (p *PlatBuffStatsOper) getPath() string{
    return "plat-buff-stats/oper"
}

func (p *PlatBuffStatsOper) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataPlatBuffStatsOper,error) {
logger.Println("PlatBuffStatsOper::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataPlatBuffStatsOper
    if err == nil {
        if len(axResp) > 0{
        err = json.Unmarshal(axResp, &p)
        }
        if err != nil {
            logger.Println("json.Unmarshal() failed with error", err)
        }
    }
    return payload,err
}
