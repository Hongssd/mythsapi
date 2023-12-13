package futures_code

import "github.com/Hongssd/mythsapi/codesdef"

type FuturesCode codesdef.Code

type FuturesCodes []FuturesCode

func (fcs *FuturesCodes) ConvertToCodes() codesdef.Codes {
	var codes codesdef.Codes
	for _, fc := range *fcs {
		codes = append(codes, codesdef.Code(fc))
	}
	return codes
}

const (
	GC0W_CMX = "@GC0W.CMX" //纽约金主连
	SI0W_CMX = "@SI0W.CMX" //纽约银主连
	HG0W_CMX = "@HG0W.CMX" //纽约铜主连
	PL0W_NMX = "@PL0W.NMX" //铂主连
	CL0W_NMX = "@CL0W.NMX" //WTI原油主连
	NG0W_NMX = "@NG0W.NMX" //天然气主连
	RB0W_NMX = "@RB0W.NMX" //RBOB汽油主连
)
