package file

//pseudo endpoint for upload SSL cert/key archive file

import (
	"github.com/a10networks/terraform-provider-thunder/thunder/axapi"
	endpnt "github.com/a10networks/terraform-provider-thunder/thunder/axapi/endpoint"
)

type SslCertKey struct {
	Name            string
	Protocol        string
	Host            string
	Path            string
	Username        string
	Password        string
	UseMgmtPort     int
	Overwrite       int
	Secured         int
}

func (p *SslCertKey) GetId() string {
	return p.Name
}

func (p *SslCertKey) Post(authToken string, host string, logger *axapi.ThunderLog) error {
	logger.Println("file.SslCertKey::Post")

	realInst := endpnt.Import{}
	realInst.Inst.SslCertKey = "bulk"
	realInst.Inst.UseMgmtPort = p.UseMgmtPort
	realInst.Inst.Overwrite = p.Overwrite
	realInst.Inst.Secured = p.Secured
	if len(p.Password) > 0 {
		realInst.Inst.Password = p.Password
	}
	realInst.Inst.RemoteFile = axapi.GenUrlForDownload(p.Protocol, p.Host, p.Path, p.Username)

	return realInst.Post(authToken, host, logger)
}

func (p *SslCertKey) Get(authToken string, host, instId string, logger *axapi.ThunderLog) error {
	logger.Println("file.SslCertKey::Get")
	logger.Println("No way to get files imported by 'import cert-key'.")
	return nil
}

func (p *SslCertKey) Delete(authToken string, host string, instId string, logger *axapi.ThunderLog) error {
	logger.Println("file.SslCertKey::Delet")
	logger.Println("No way to delete cert/key files imported by 'import cert-key'.")
	return nil
}
