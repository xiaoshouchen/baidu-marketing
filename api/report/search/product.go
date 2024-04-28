package search

import (
	reportApi "github.com/bububa/baidu-marketing/api/report"
	"github.com/bububa/baidu-marketing/core"
	"github.com/bububa/baidu-marketing/model"
	"github.com/bububa/baidu-marketing/model/report"
)

// Product 搜索商品报告
func Product(clt *core.SDKClient, auth *model.RequestHeader, reportRequest *report.GetReportDataRequest) (*model.ResponseHeader, []report.ReportData, error) {
	reportRequest.ReportType = 248654
	return reportApi.GetReportData(clt, auth, reportRequest)
}
