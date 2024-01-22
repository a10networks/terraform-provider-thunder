

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type RuleSetTagStats struct {
    
    Stats RuleSetTagStatsStats `json:"stats"`
    Name string 

}
type DataRuleSetTagStats struct {
    DtRuleSetTagStats RuleSetTagStats `json:"tag"`
}


type RuleSetTagStatsStats struct {
    Categorystat1 int `json:"categorystat1"`
    Categorystat2 int `json:"categorystat2"`
    Categorystat3 int `json:"categorystat3"`
    Categorystat4 int `json:"categorystat4"`
    Categorystat5 int `json:"categorystat5"`
    Categorystat6 int `json:"categorystat6"`
    Categorystat7 int `json:"categorystat7"`
    Categorystat8 int `json:"categorystat8"`
    Categorystat9 int `json:"categorystat9"`
    Categorystat10 int `json:"categorystat10"`
    Categorystat11 int `json:"categorystat11"`
    Categorystat12 int `json:"categorystat12"`
    Categorystat13 int `json:"categorystat13"`
    Categorystat14 int `json:"categorystat14"`
    Categorystat15 int `json:"categorystat15"`
    Categorystat16 int `json:"categorystat16"`
    Categorystat17 int `json:"categorystat17"`
    Categorystat18 int `json:"categorystat18"`
    Categorystat19 int `json:"categorystat19"`
    Categorystat20 int `json:"categorystat20"`
    Categorystat21 int `json:"categorystat21"`
    Categorystat22 int `json:"categorystat22"`
    Categorystat23 int `json:"categorystat23"`
    Categorystat24 int `json:"categorystat24"`
    Categorystat25 int `json:"categorystat25"`
    Categorystat26 int `json:"categorystat26"`
    Categorystat27 int `json:"categorystat27"`
    Categorystat28 int `json:"categorystat28"`
    Categorystat29 int `json:"categorystat29"`
    Categorystat30 int `json:"categorystat30"`
    Categorystat31 int `json:"categorystat31"`
    Categorystat32 int `json:"categorystat32"`
    Categorystat33 int `json:"categorystat33"`
    Categorystat34 int `json:"categorystat34"`
    Categorystat35 int `json:"categorystat35"`
    Categorystat36 int `json:"categorystat36"`
    Categorystat37 int `json:"categorystat37"`
    Categorystat38 int `json:"categorystat38"`
    Categorystat39 int `json:"categorystat39"`
    Categorystat40 int `json:"categorystat40"`
    Categorystat41 int `json:"categorystat41"`
    Categorystat42 int `json:"categorystat42"`
    Categorystat43 int `json:"categorystat43"`
    Categorystat44 int `json:"categorystat44"`
    Categorystat45 int `json:"categorystat45"`
    Categorystat46 int `json:"categorystat46"`
    Categorystat47 int `json:"categorystat47"`
    Categorystat48 int `json:"categorystat48"`
    Categorystat49 int `json:"categorystat49"`
    Categorystat50 int `json:"categorystat50"`
    Categorystat51 int `json:"categorystat51"`
    Categorystat52 int `json:"categorystat52"`
    Categorystat53 int `json:"categorystat53"`
    Categorystat54 int `json:"categorystat54"`
    Categorystat55 int `json:"categorystat55"`
    Categorystat56 int `json:"categorystat56"`
    Categorystat57 int `json:"categorystat57"`
    Categorystat58 int `json:"categorystat58"`
    Categorystat59 int `json:"categorystat59"`
    Categorystat60 int `json:"categorystat60"`
    Categorystat61 int `json:"categorystat61"`
    Categorystat62 int `json:"categorystat62"`
    Categorystat63 int `json:"categorystat63"`
    Categorystat64 int `json:"categorystat64"`
    Categorystat65 int `json:"categorystat65"`
    Categorystat66 int `json:"categorystat66"`
    Categorystat67 int `json:"categorystat67"`
    Categorystat68 int `json:"categorystat68"`
    Categorystat69 int `json:"categorystat69"`
    Categorystat70 int `json:"categorystat70"`
    Categorystat71 int `json:"categorystat71"`
    Categorystat72 int `json:"categorystat72"`
    Categorystat73 int `json:"categorystat73"`
    Categorystat74 int `json:"categorystat74"`
    Categorystat75 int `json:"categorystat75"`
    Categorystat76 int `json:"categorystat76"`
    Categorystat77 int `json:"categorystat77"`
    Categorystat78 int `json:"categorystat78"`
    Categorystat79 int `json:"categorystat79"`
    Categorystat80 int `json:"categorystat80"`
    Categorystat81 int `json:"categorystat81"`
    Categorystat82 int `json:"categorystat82"`
    Categorystat83 int `json:"categorystat83"`
    Categorystat84 int `json:"categorystat84"`
    Categorystat85 int `json:"categorystat85"`
    Categorystat86 int `json:"categorystat86"`
    Categorystat87 int `json:"categorystat87"`
    Categorystat88 int `json:"categorystat88"`
    Categorystat89 int `json:"categorystat89"`
    Categorystat90 int `json:"categorystat90"`
    Categorystat91 int `json:"categorystat91"`
    Categorystat92 int `json:"categorystat92"`
    Categorystat93 int `json:"categorystat93"`
    Categorystat94 int `json:"categorystat94"`
    Categorystat95 int `json:"categorystat95"`
    Categorystat96 int `json:"categorystat96"`
    Categorystat97 int `json:"categorystat97"`
    Categorystat98 int `json:"categorystat98"`
    Categorystat99 int `json:"categorystat99"`
    Categorystat100 int `json:"categorystat100"`
    Categorystat101 int `json:"categorystat101"`
    Categorystat102 int `json:"categorystat102"`
    Categorystat103 int `json:"categorystat103"`
    Categorystat104 int `json:"categorystat104"`
    Categorystat105 int `json:"categorystat105"`
    Categorystat106 int `json:"categorystat106"`
    Categorystat107 int `json:"categorystat107"`
    Categorystat108 int `json:"categorystat108"`
    Categorystat109 int `json:"categorystat109"`
    Categorystat110 int `json:"categorystat110"`
    Categorystat111 int `json:"categorystat111"`
    Categorystat112 int `json:"categorystat112"`
    Categorystat113 int `json:"categorystat113"`
    Categorystat114 int `json:"categorystat114"`
    Categorystat115 int `json:"categorystat115"`
    Categorystat116 int `json:"categorystat116"`
    Categorystat117 int `json:"categorystat117"`
    Categorystat118 int `json:"categorystat118"`
    Categorystat119 int `json:"categorystat119"`
    Categorystat120 int `json:"categorystat120"`
    Categorystat121 int `json:"categorystat121"`
    Categorystat122 int `json:"categorystat122"`
    Categorystat123 int `json:"categorystat123"`
    Categorystat124 int `json:"categorystat124"`
    Categorystat125 int `json:"categorystat125"`
    Categorystat126 int `json:"categorystat126"`
    Categorystat127 int `json:"categorystat127"`
    Categorystat128 int `json:"categorystat128"`
    Categorystat129 int `json:"categorystat129"`
    Categorystat130 int `json:"categorystat130"`
    Categorystat131 int `json:"categorystat131"`
    Categorystat132 int `json:"categorystat132"`
    Categorystat133 int `json:"categorystat133"`
    Categorystat134 int `json:"categorystat134"`
    Categorystat135 int `json:"categorystat135"`
    Categorystat136 int `json:"categorystat136"`
    Categorystat137 int `json:"categorystat137"`
    Categorystat138 int `json:"categorystat138"`
    Categorystat139 int `json:"categorystat139"`
    Categorystat140 int `json:"categorystat140"`
    Categorystat141 int `json:"categorystat141"`
    Categorystat142 int `json:"categorystat142"`
    Categorystat143 int `json:"categorystat143"`
    Categorystat144 int `json:"categorystat144"`
    Categorystat145 int `json:"categorystat145"`
    Categorystat146 int `json:"categorystat146"`
    Categorystat147 int `json:"categorystat147"`
    Categorystat148 int `json:"categorystat148"`
    Categorystat149 int `json:"categorystat149"`
    Categorystat150 int `json:"categorystat150"`
    Categorystat151 int `json:"categorystat151"`
    Categorystat152 int `json:"categorystat152"`
    Categorystat153 int `json:"categorystat153"`
    Categorystat154 int `json:"categorystat154"`
    Categorystat155 int `json:"categorystat155"`
    Categorystat156 int `json:"categorystat156"`
    Categorystat157 int `json:"categorystat157"`
    Categorystat158 int `json:"categorystat158"`
    Categorystat159 int `json:"categorystat159"`
    Categorystat160 int `json:"categorystat160"`
    Categorystat161 int `json:"categorystat161"`
    Categorystat162 int `json:"categorystat162"`
    Categorystat163 int `json:"categorystat163"`
    Categorystat164 int `json:"categorystat164"`
    Categorystat165 int `json:"categorystat165"`
    Categorystat166 int `json:"categorystat166"`
    Categorystat167 int `json:"categorystat167"`
    Categorystat168 int `json:"categorystat168"`
    Categorystat169 int `json:"categorystat169"`
    Categorystat170 int `json:"categorystat170"`
    Categorystat171 int `json:"categorystat171"`
    Categorystat172 int `json:"categorystat172"`
    Categorystat173 int `json:"categorystat173"`
    Categorystat174 int `json:"categorystat174"`
    Categorystat175 int `json:"categorystat175"`
    Categorystat176 int `json:"categorystat176"`
    Categorystat177 int `json:"categorystat177"`
    Categorystat178 int `json:"categorystat178"`
    Categorystat179 int `json:"categorystat179"`
    Categorystat180 int `json:"categorystat180"`
    Categorystat181 int `json:"categorystat181"`
    Categorystat182 int `json:"categorystat182"`
    Categorystat183 int `json:"categorystat183"`
    Categorystat184 int `json:"categorystat184"`
    Categorystat185 int `json:"categorystat185"`
    Categorystat186 int `json:"categorystat186"`
    Categorystat187 int `json:"categorystat187"`
    Categorystat188 int `json:"categorystat188"`
    Categorystat189 int `json:"categorystat189"`
    Categorystat190 int `json:"categorystat190"`
    Categorystat191 int `json:"categorystat191"`
    Categorystat192 int `json:"categorystat192"`
    Categorystat193 int `json:"categorystat193"`
    Categorystat194 int `json:"categorystat194"`
    Categorystat195 int `json:"categorystat195"`
    Categorystat196 int `json:"categorystat196"`
    Categorystat197 int `json:"categorystat197"`
    Categorystat198 int `json:"categorystat198"`
    Categorystat199 int `json:"categorystat199"`
    Categorystat200 int `json:"categorystat200"`
    Categorystat201 int `json:"categorystat201"`
    Categorystat202 int `json:"categorystat202"`
    Categorystat203 int `json:"categorystat203"`
    Categorystat204 int `json:"categorystat204"`
    Categorystat205 int `json:"categorystat205"`
    Categorystat206 int `json:"categorystat206"`
    Categorystat207 int `json:"categorystat207"`
    Categorystat208 int `json:"categorystat208"`
    Categorystat209 int `json:"categorystat209"`
    Categorystat210 int `json:"categorystat210"`
    Categorystat211 int `json:"categorystat211"`
    Categorystat212 int `json:"categorystat212"`
    Categorystat213 int `json:"categorystat213"`
    Categorystat214 int `json:"categorystat214"`
    Categorystat215 int `json:"categorystat215"`
    Categorystat216 int `json:"categorystat216"`
    Categorystat217 int `json:"categorystat217"`
    Categorystat218 int `json:"categorystat218"`
    Categorystat219 int `json:"categorystat219"`
    Categorystat220 int `json:"categorystat220"`
    Categorystat221 int `json:"categorystat221"`
    Categorystat222 int `json:"categorystat222"`
    Categorystat223 int `json:"categorystat223"`
    Categorystat224 int `json:"categorystat224"`
    Categorystat225 int `json:"categorystat225"`
    Categorystat226 int `json:"categorystat226"`
    Categorystat227 int `json:"categorystat227"`
    Categorystat228 int `json:"categorystat228"`
    Categorystat229 int `json:"categorystat229"`
    Categorystat230 int `json:"categorystat230"`
    Categorystat231 int `json:"categorystat231"`
    Categorystat232 int `json:"categorystat232"`
    Categorystat233 int `json:"categorystat233"`
    Categorystat234 int `json:"categorystat234"`
    Categorystat235 int `json:"categorystat235"`
    Categorystat236 int `json:"categorystat236"`
    Categorystat237 int `json:"categorystat237"`
    Categorystat238 int `json:"categorystat238"`
    Categorystat239 int `json:"categorystat239"`
    Categorystat240 int `json:"categorystat240"`
    Categorystat241 int `json:"categorystat241"`
    Categorystat242 int `json:"categorystat242"`
    Categorystat243 int `json:"categorystat243"`
    Categorystat244 int `json:"categorystat244"`
    Categorystat245 int `json:"categorystat245"`
    Categorystat246 int `json:"categorystat246"`
    Categorystat247 int `json:"categorystat247"`
    Categorystat248 int `json:"categorystat248"`
    Categorystat249 int `json:"categorystat249"`
    Categorystat250 int `json:"categorystat250"`
    Categorystat251 int `json:"categorystat251"`
    Categorystat252 int `json:"categorystat252"`
    Categorystat253 int `json:"categorystat253"`
    Categorystat254 int `json:"categorystat254"`
    Categorystat255 int `json:"categorystat255"`
    Categorystat256 int `json:"categorystat256"`
}

func (p *RuleSetTagStats) GetId() string{
    return "1"
}

func (p *RuleSetTagStats) getPath() string{
    
    return "rule-set/" +p.Name + "/tag/stats"
}

func (p *RuleSetTagStats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataRuleSetTagStats,error) {
logger.Println("RuleSetTagStats::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataRuleSetTagStats
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
