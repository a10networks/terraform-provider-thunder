

package endpoint
import (
    "github.com/a10networks/terraform-provider-thunder/thunder/axapi"
    "github.com/clarketm/json"
)

//based on ACOS 6_0_2_P1-37
type RuleSetAppStats struct {
    
    Stats RuleSetAppStatsStats `json:"stats"`
    Name string 

}
type DataRuleSetAppStats struct {
    DtRuleSetAppStats RuleSetAppStats `json:"app"`
}


type RuleSetAppStatsStats struct {
    Appstat1 int `json:"appstat1"`
    Appstat2 int `json:"appstat2"`
    Appstat3 int `json:"appstat3"`
    Appstat4 int `json:"appstat4"`
    Appstat5 int `json:"appstat5"`
    Appstat6 int `json:"appstat6"`
    Appstat7 int `json:"appstat7"`
    Appstat8 int `json:"appstat8"`
    Appstat9 int `json:"appstat9"`
    Appstat10 int `json:"appstat10"`
    Appstat11 int `json:"appstat11"`
    Appstat12 int `json:"appstat12"`
    Appstat13 int `json:"appstat13"`
    Appstat14 int `json:"appstat14"`
    Appstat15 int `json:"appstat15"`
    Appstat16 int `json:"appstat16"`
    Appstat17 int `json:"appstat17"`
    Appstat18 int `json:"appstat18"`
    Appstat19 int `json:"appstat19"`
    Appstat20 int `json:"appstat20"`
    Appstat21 int `json:"appstat21"`
    Appstat22 int `json:"appstat22"`
    Appstat23 int `json:"appstat23"`
    Appstat24 int `json:"appstat24"`
    Appstat25 int `json:"appstat25"`
    Appstat26 int `json:"appstat26"`
    Appstat27 int `json:"appstat27"`
    Appstat28 int `json:"appstat28"`
    Appstat29 int `json:"appstat29"`
    Appstat30 int `json:"appstat30"`
    Appstat31 int `json:"appstat31"`
    Appstat32 int `json:"appstat32"`
    Appstat33 int `json:"appstat33"`
    Appstat34 int `json:"appstat34"`
    Appstat35 int `json:"appstat35"`
    Appstat36 int `json:"appstat36"`
    Appstat37 int `json:"appstat37"`
    Appstat38 int `json:"appstat38"`
    Appstat39 int `json:"appstat39"`
    Appstat40 int `json:"appstat40"`
    Appstat41 int `json:"appstat41"`
    Appstat42 int `json:"appstat42"`
    Appstat43 int `json:"appstat43"`
    Appstat44 int `json:"appstat44"`
    Appstat45 int `json:"appstat45"`
    Appstat46 int `json:"appstat46"`
    Appstat47 int `json:"appstat47"`
    Appstat48 int `json:"appstat48"`
    Appstat49 int `json:"appstat49"`
    Appstat50 int `json:"appstat50"`
    Appstat51 int `json:"appstat51"`
    Appstat52 int `json:"appstat52"`
    Appstat53 int `json:"appstat53"`
    Appstat54 int `json:"appstat54"`
    Appstat55 int `json:"appstat55"`
    Appstat56 int `json:"appstat56"`
    Appstat57 int `json:"appstat57"`
    Appstat58 int `json:"appstat58"`
    Appstat59 int `json:"appstat59"`
    Appstat60 int `json:"appstat60"`
    Appstat61 int `json:"appstat61"`
    Appstat62 int `json:"appstat62"`
    Appstat63 int `json:"appstat63"`
    Appstat64 int `json:"appstat64"`
    Appstat65 int `json:"appstat65"`
    Appstat66 int `json:"appstat66"`
    Appstat67 int `json:"appstat67"`
    Appstat68 int `json:"appstat68"`
    Appstat69 int `json:"appstat69"`
    Appstat70 int `json:"appstat70"`
    Appstat71 int `json:"appstat71"`
    Appstat72 int `json:"appstat72"`
    Appstat73 int `json:"appstat73"`
    Appstat74 int `json:"appstat74"`
    Appstat75 int `json:"appstat75"`
    Appstat76 int `json:"appstat76"`
    Appstat77 int `json:"appstat77"`
    Appstat78 int `json:"appstat78"`
    Appstat79 int `json:"appstat79"`
    Appstat80 int `json:"appstat80"`
    Appstat81 int `json:"appstat81"`
    Appstat82 int `json:"appstat82"`
    Appstat83 int `json:"appstat83"`
    Appstat84 int `json:"appstat84"`
    Appstat85 int `json:"appstat85"`
    Appstat86 int `json:"appstat86"`
    Appstat87 int `json:"appstat87"`
    Appstat88 int `json:"appstat88"`
    Appstat89 int `json:"appstat89"`
    Appstat90 int `json:"appstat90"`
    Appstat91 int `json:"appstat91"`
    Appstat92 int `json:"appstat92"`
    Appstat93 int `json:"appstat93"`
    Appstat94 int `json:"appstat94"`
    Appstat95 int `json:"appstat95"`
    Appstat96 int `json:"appstat96"`
    Appstat97 int `json:"appstat97"`
    Appstat98 int `json:"appstat98"`
    Appstat99 int `json:"appstat99"`
    Appstat100 int `json:"appstat100"`
    Appstat101 int `json:"appstat101"`
    Appstat102 int `json:"appstat102"`
    Appstat103 int `json:"appstat103"`
    Appstat104 int `json:"appstat104"`
    Appstat105 int `json:"appstat105"`
    Appstat106 int `json:"appstat106"`
    Appstat107 int `json:"appstat107"`
    Appstat108 int `json:"appstat108"`
    Appstat109 int `json:"appstat109"`
    Appstat110 int `json:"appstat110"`
    Appstat111 int `json:"appstat111"`
    Appstat112 int `json:"appstat112"`
    Appstat113 int `json:"appstat113"`
    Appstat114 int `json:"appstat114"`
    Appstat115 int `json:"appstat115"`
    Appstat116 int `json:"appstat116"`
    Appstat117 int `json:"appstat117"`
    Appstat118 int `json:"appstat118"`
    Appstat119 int `json:"appstat119"`
    Appstat120 int `json:"appstat120"`
    Appstat121 int `json:"appstat121"`
    Appstat122 int `json:"appstat122"`
    Appstat123 int `json:"appstat123"`
    Appstat124 int `json:"appstat124"`
    Appstat125 int `json:"appstat125"`
    Appstat126 int `json:"appstat126"`
    Appstat127 int `json:"appstat127"`
    Appstat128 int `json:"appstat128"`
    Appstat129 int `json:"appstat129"`
    Appstat130 int `json:"appstat130"`
    Appstat131 int `json:"appstat131"`
    Appstat132 int `json:"appstat132"`
    Appstat133 int `json:"appstat133"`
    Appstat134 int `json:"appstat134"`
    Appstat135 int `json:"appstat135"`
    Appstat136 int `json:"appstat136"`
    Appstat137 int `json:"appstat137"`
    Appstat138 int `json:"appstat138"`
    Appstat139 int `json:"appstat139"`
    Appstat140 int `json:"appstat140"`
    Appstat141 int `json:"appstat141"`
    Appstat142 int `json:"appstat142"`
    Appstat143 int `json:"appstat143"`
    Appstat144 int `json:"appstat144"`
    Appstat145 int `json:"appstat145"`
    Appstat146 int `json:"appstat146"`
    Appstat147 int `json:"appstat147"`
    Appstat148 int `json:"appstat148"`
    Appstat149 int `json:"appstat149"`
    Appstat150 int `json:"appstat150"`
    Appstat151 int `json:"appstat151"`
    Appstat152 int `json:"appstat152"`
    Appstat153 int `json:"appstat153"`
    Appstat154 int `json:"appstat154"`
    Appstat155 int `json:"appstat155"`
    Appstat156 int `json:"appstat156"`
    Appstat157 int `json:"appstat157"`
    Appstat158 int `json:"appstat158"`
    Appstat159 int `json:"appstat159"`
    Appstat160 int `json:"appstat160"`
    Appstat161 int `json:"appstat161"`
    Appstat162 int `json:"appstat162"`
    Appstat163 int `json:"appstat163"`
    Appstat164 int `json:"appstat164"`
    Appstat165 int `json:"appstat165"`
    Appstat166 int `json:"appstat166"`
    Appstat167 int `json:"appstat167"`
    Appstat168 int `json:"appstat168"`
    Appstat169 int `json:"appstat169"`
    Appstat170 int `json:"appstat170"`
    Appstat171 int `json:"appstat171"`
    Appstat172 int `json:"appstat172"`
    Appstat173 int `json:"appstat173"`
    Appstat174 int `json:"appstat174"`
    Appstat175 int `json:"appstat175"`
    Appstat176 int `json:"appstat176"`
    Appstat177 int `json:"appstat177"`
    Appstat178 int `json:"appstat178"`
    Appstat179 int `json:"appstat179"`
    Appstat180 int `json:"appstat180"`
    Appstat181 int `json:"appstat181"`
    Appstat182 int `json:"appstat182"`
    Appstat183 int `json:"appstat183"`
    Appstat184 int `json:"appstat184"`
    Appstat185 int `json:"appstat185"`
    Appstat186 int `json:"appstat186"`
    Appstat187 int `json:"appstat187"`
    Appstat188 int `json:"appstat188"`
    Appstat189 int `json:"appstat189"`
    Appstat190 int `json:"appstat190"`
    Appstat191 int `json:"appstat191"`
    Appstat192 int `json:"appstat192"`
    Appstat193 int `json:"appstat193"`
    Appstat194 int `json:"appstat194"`
    Appstat195 int `json:"appstat195"`
    Appstat196 int `json:"appstat196"`
    Appstat197 int `json:"appstat197"`
    Appstat198 int `json:"appstat198"`
    Appstat199 int `json:"appstat199"`
    Appstat200 int `json:"appstat200"`
    Appstat201 int `json:"appstat201"`
    Appstat202 int `json:"appstat202"`
    Appstat203 int `json:"appstat203"`
    Appstat204 int `json:"appstat204"`
    Appstat205 int `json:"appstat205"`
    Appstat206 int `json:"appstat206"`
    Appstat207 int `json:"appstat207"`
    Appstat208 int `json:"appstat208"`
    Appstat209 int `json:"appstat209"`
    Appstat210 int `json:"appstat210"`
    Appstat211 int `json:"appstat211"`
    Appstat212 int `json:"appstat212"`
    Appstat213 int `json:"appstat213"`
    Appstat214 int `json:"appstat214"`
    Appstat215 int `json:"appstat215"`
    Appstat216 int `json:"appstat216"`
    Appstat217 int `json:"appstat217"`
    Appstat218 int `json:"appstat218"`
    Appstat219 int `json:"appstat219"`
    Appstat220 int `json:"appstat220"`
    Appstat221 int `json:"appstat221"`
    Appstat222 int `json:"appstat222"`
    Appstat223 int `json:"appstat223"`
    Appstat224 int `json:"appstat224"`
    Appstat225 int `json:"appstat225"`
    Appstat226 int `json:"appstat226"`
    Appstat227 int `json:"appstat227"`
    Appstat228 int `json:"appstat228"`
    Appstat229 int `json:"appstat229"`
    Appstat230 int `json:"appstat230"`
    Appstat231 int `json:"appstat231"`
    Appstat232 int `json:"appstat232"`
    Appstat233 int `json:"appstat233"`
    Appstat234 int `json:"appstat234"`
    Appstat235 int `json:"appstat235"`
    Appstat236 int `json:"appstat236"`
    Appstat237 int `json:"appstat237"`
    Appstat238 int `json:"appstat238"`
    Appstat239 int `json:"appstat239"`
    Appstat240 int `json:"appstat240"`
    Appstat241 int `json:"appstat241"`
    Appstat242 int `json:"appstat242"`
    Appstat243 int `json:"appstat243"`
    Appstat244 int `json:"appstat244"`
    Appstat245 int `json:"appstat245"`
    Appstat246 int `json:"appstat246"`
    Appstat247 int `json:"appstat247"`
    Appstat248 int `json:"appstat248"`
    Appstat249 int `json:"appstat249"`
    Appstat250 int `json:"appstat250"`
    Appstat251 int `json:"appstat251"`
    Appstat252 int `json:"appstat252"`
    Appstat253 int `json:"appstat253"`
    Appstat254 int `json:"appstat254"`
    Appstat255 int `json:"appstat255"`
    Appstat256 int `json:"appstat256"`
    Appstat257 int `json:"appstat257"`
    Appstat258 int `json:"appstat258"`
    Appstat259 int `json:"appstat259"`
    Appstat260 int `json:"appstat260"`
    Appstat261 int `json:"appstat261"`
    Appstat262 int `json:"appstat262"`
    Appstat263 int `json:"appstat263"`
    Appstat264 int `json:"appstat264"`
    Appstat265 int `json:"appstat265"`
    Appstat266 int `json:"appstat266"`
    Appstat267 int `json:"appstat267"`
    Appstat268 int `json:"appstat268"`
    Appstat269 int `json:"appstat269"`
    Appstat270 int `json:"appstat270"`
    Appstat271 int `json:"appstat271"`
    Appstat272 int `json:"appstat272"`
    Appstat273 int `json:"appstat273"`
    Appstat274 int `json:"appstat274"`
    Appstat275 int `json:"appstat275"`
    Appstat276 int `json:"appstat276"`
    Appstat277 int `json:"appstat277"`
    Appstat278 int `json:"appstat278"`
    Appstat279 int `json:"appstat279"`
    Appstat280 int `json:"appstat280"`
    Appstat281 int `json:"appstat281"`
    Appstat282 int `json:"appstat282"`
    Appstat283 int `json:"appstat283"`
    Appstat284 int `json:"appstat284"`
    Appstat285 int `json:"appstat285"`
    Appstat286 int `json:"appstat286"`
    Appstat287 int `json:"appstat287"`
    Appstat288 int `json:"appstat288"`
    Appstat289 int `json:"appstat289"`
    Appstat290 int `json:"appstat290"`
    Appstat291 int `json:"appstat291"`
    Appstat292 int `json:"appstat292"`
    Appstat293 int `json:"appstat293"`
    Appstat294 int `json:"appstat294"`
    Appstat295 int `json:"appstat295"`
    Appstat296 int `json:"appstat296"`
    Appstat297 int `json:"appstat297"`
    Appstat298 int `json:"appstat298"`
    Appstat299 int `json:"appstat299"`
    Appstat300 int `json:"appstat300"`
    Appstat301 int `json:"appstat301"`
    Appstat302 int `json:"appstat302"`
    Appstat303 int `json:"appstat303"`
    Appstat304 int `json:"appstat304"`
    Appstat305 int `json:"appstat305"`
    Appstat306 int `json:"appstat306"`
    Appstat307 int `json:"appstat307"`
    Appstat308 int `json:"appstat308"`
    Appstat309 int `json:"appstat309"`
    Appstat310 int `json:"appstat310"`
    Appstat311 int `json:"appstat311"`
    Appstat312 int `json:"appstat312"`
    Appstat313 int `json:"appstat313"`
    Appstat314 int `json:"appstat314"`
    Appstat315 int `json:"appstat315"`
    Appstat316 int `json:"appstat316"`
    Appstat317 int `json:"appstat317"`
    Appstat318 int `json:"appstat318"`
    Appstat319 int `json:"appstat319"`
    Appstat320 int `json:"appstat320"`
    Appstat321 int `json:"appstat321"`
    Appstat322 int `json:"appstat322"`
    Appstat323 int `json:"appstat323"`
    Appstat324 int `json:"appstat324"`
    Appstat325 int `json:"appstat325"`
    Appstat326 int `json:"appstat326"`
    Appstat327 int `json:"appstat327"`
    Appstat328 int `json:"appstat328"`
    Appstat329 int `json:"appstat329"`
    Appstat330 int `json:"appstat330"`
    Appstat331 int `json:"appstat331"`
    Appstat332 int `json:"appstat332"`
    Appstat333 int `json:"appstat333"`
    Appstat334 int `json:"appstat334"`
    Appstat335 int `json:"appstat335"`
    Appstat336 int `json:"appstat336"`
    Appstat337 int `json:"appstat337"`
    Appstat338 int `json:"appstat338"`
    Appstat339 int `json:"appstat339"`
    Appstat340 int `json:"appstat340"`
    Appstat341 int `json:"appstat341"`
    Appstat342 int `json:"appstat342"`
    Appstat343 int `json:"appstat343"`
    Appstat344 int `json:"appstat344"`
    Appstat345 int `json:"appstat345"`
    Appstat346 int `json:"appstat346"`
    Appstat347 int `json:"appstat347"`
    Appstat348 int `json:"appstat348"`
    Appstat349 int `json:"appstat349"`
    Appstat350 int `json:"appstat350"`
    Appstat351 int `json:"appstat351"`
    Appstat352 int `json:"appstat352"`
    Appstat353 int `json:"appstat353"`
    Appstat354 int `json:"appstat354"`
    Appstat355 int `json:"appstat355"`
    Appstat356 int `json:"appstat356"`
    Appstat357 int `json:"appstat357"`
    Appstat358 int `json:"appstat358"`
    Appstat359 int `json:"appstat359"`
    Appstat360 int `json:"appstat360"`
    Appstat361 int `json:"appstat361"`
    Appstat362 int `json:"appstat362"`
    Appstat363 int `json:"appstat363"`
    Appstat364 int `json:"appstat364"`
    Appstat365 int `json:"appstat365"`
    Appstat366 int `json:"appstat366"`
    Appstat367 int `json:"appstat367"`
    Appstat368 int `json:"appstat368"`
    Appstat369 int `json:"appstat369"`
    Appstat370 int `json:"appstat370"`
    Appstat371 int `json:"appstat371"`
    Appstat372 int `json:"appstat372"`
    Appstat373 int `json:"appstat373"`
    Appstat374 int `json:"appstat374"`
    Appstat375 int `json:"appstat375"`
    Appstat376 int `json:"appstat376"`
    Appstat377 int `json:"appstat377"`
    Appstat378 int `json:"appstat378"`
    Appstat379 int `json:"appstat379"`
    Appstat380 int `json:"appstat380"`
    Appstat381 int `json:"appstat381"`
    Appstat382 int `json:"appstat382"`
    Appstat383 int `json:"appstat383"`
    Appstat384 int `json:"appstat384"`
    Appstat385 int `json:"appstat385"`
    Appstat386 int `json:"appstat386"`
    Appstat387 int `json:"appstat387"`
    Appstat388 int `json:"appstat388"`
    Appstat389 int `json:"appstat389"`
    Appstat390 int `json:"appstat390"`
    Appstat391 int `json:"appstat391"`
    Appstat392 int `json:"appstat392"`
    Appstat393 int `json:"appstat393"`
    Appstat394 int `json:"appstat394"`
    Appstat395 int `json:"appstat395"`
    Appstat396 int `json:"appstat396"`
    Appstat397 int `json:"appstat397"`
    Appstat398 int `json:"appstat398"`
    Appstat399 int `json:"appstat399"`
    Appstat400 int `json:"appstat400"`
    Appstat401 int `json:"appstat401"`
    Appstat402 int `json:"appstat402"`
    Appstat403 int `json:"appstat403"`
    Appstat404 int `json:"appstat404"`
    Appstat405 int `json:"appstat405"`
    Appstat406 int `json:"appstat406"`
    Appstat407 int `json:"appstat407"`
    Appstat408 int `json:"appstat408"`
    Appstat409 int `json:"appstat409"`
    Appstat410 int `json:"appstat410"`
    Appstat411 int `json:"appstat411"`
    Appstat412 int `json:"appstat412"`
    Appstat413 int `json:"appstat413"`
    Appstat414 int `json:"appstat414"`
    Appstat415 int `json:"appstat415"`
    Appstat416 int `json:"appstat416"`
    Appstat417 int `json:"appstat417"`
    Appstat418 int `json:"appstat418"`
    Appstat419 int `json:"appstat419"`
    Appstat420 int `json:"appstat420"`
    Appstat421 int `json:"appstat421"`
    Appstat422 int `json:"appstat422"`
    Appstat423 int `json:"appstat423"`
    Appstat424 int `json:"appstat424"`
    Appstat425 int `json:"appstat425"`
    Appstat426 int `json:"appstat426"`
    Appstat427 int `json:"appstat427"`
    Appstat428 int `json:"appstat428"`
    Appstat429 int `json:"appstat429"`
    Appstat430 int `json:"appstat430"`
    Appstat431 int `json:"appstat431"`
    Appstat432 int `json:"appstat432"`
    Appstat433 int `json:"appstat433"`
    Appstat434 int `json:"appstat434"`
    Appstat435 int `json:"appstat435"`
    Appstat436 int `json:"appstat436"`
    Appstat437 int `json:"appstat437"`
    Appstat438 int `json:"appstat438"`
    Appstat439 int `json:"appstat439"`
    Appstat440 int `json:"appstat440"`
    Appstat441 int `json:"appstat441"`
    Appstat442 int `json:"appstat442"`
    Appstat443 int `json:"appstat443"`
    Appstat444 int `json:"appstat444"`
    Appstat445 int `json:"appstat445"`
    Appstat446 int `json:"appstat446"`
    Appstat447 int `json:"appstat447"`
    Appstat448 int `json:"appstat448"`
    Appstat449 int `json:"appstat449"`
    Appstat450 int `json:"appstat450"`
    Appstat451 int `json:"appstat451"`
    Appstat452 int `json:"appstat452"`
    Appstat453 int `json:"appstat453"`
    Appstat454 int `json:"appstat454"`
    Appstat455 int `json:"appstat455"`
    Appstat456 int `json:"appstat456"`
    Appstat457 int `json:"appstat457"`
    Appstat458 int `json:"appstat458"`
    Appstat459 int `json:"appstat459"`
    Appstat460 int `json:"appstat460"`
    Appstat461 int `json:"appstat461"`
    Appstat462 int `json:"appstat462"`
    Appstat463 int `json:"appstat463"`
    Appstat464 int `json:"appstat464"`
    Appstat465 int `json:"appstat465"`
    Appstat466 int `json:"appstat466"`
    Appstat467 int `json:"appstat467"`
    Appstat468 int `json:"appstat468"`
    Appstat469 int `json:"appstat469"`
    Appstat470 int `json:"appstat470"`
    Appstat471 int `json:"appstat471"`
    Appstat472 int `json:"appstat472"`
    Appstat473 int `json:"appstat473"`
    Appstat474 int `json:"appstat474"`
    Appstat475 int `json:"appstat475"`
    Appstat476 int `json:"appstat476"`
    Appstat477 int `json:"appstat477"`
    Appstat478 int `json:"appstat478"`
    Appstat479 int `json:"appstat479"`
    Appstat480 int `json:"appstat480"`
    Appstat481 int `json:"appstat481"`
    Appstat482 int `json:"appstat482"`
    Appstat483 int `json:"appstat483"`
    Appstat484 int `json:"appstat484"`
    Appstat485 int `json:"appstat485"`
    Appstat486 int `json:"appstat486"`
    Appstat487 int `json:"appstat487"`
    Appstat488 int `json:"appstat488"`
    Appstat489 int `json:"appstat489"`
    Appstat490 int `json:"appstat490"`
    Appstat491 int `json:"appstat491"`
    Appstat492 int `json:"appstat492"`
    Appstat493 int `json:"appstat493"`
    Appstat494 int `json:"appstat494"`
    Appstat495 int `json:"appstat495"`
    Appstat496 int `json:"appstat496"`
    Appstat497 int `json:"appstat497"`
    Appstat498 int `json:"appstat498"`
    Appstat499 int `json:"appstat499"`
    Appstat500 int `json:"appstat500"`
    Appstat501 int `json:"appstat501"`
    Appstat502 int `json:"appstat502"`
    Appstat503 int `json:"appstat503"`
    Appstat504 int `json:"appstat504"`
    Appstat505 int `json:"appstat505"`
    Appstat506 int `json:"appstat506"`
    Appstat507 int `json:"appstat507"`
    Appstat508 int `json:"appstat508"`
    Appstat509 int `json:"appstat509"`
    Appstat510 int `json:"appstat510"`
    Appstat511 int `json:"appstat511"`
}

func (p *RuleSetAppStats) GetId() string{
    return "1"
}

func (p *RuleSetAppStats) getPath() string{
    
    return "rule-set/" +p.Name + "/app/stats"
}

func (p *RuleSetAppStats) Get(authToken string, host string, instId string, logger *axapi.ThunderLog) (DataRuleSetAppStats,error) {
logger.Println("RuleSetAppStats::Get")
    headers := axapi.GenRequestHeader(authToken)
    _, axResp, err := axapi.SendGet(host, p.getPath(), "", nil, headers, logger)
    var payload DataRuleSetAppStats
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
