package feed

import (
	reportApi "github.com/bububa/baidu-marketing/api/report"
	"github.com/bububa/baidu-marketing/core"
	"github.com/bububa/baidu-marketing/model"
	"github.com/bububa/baidu-marketing/model/report"
)

// Creative 创意报告
func Creative(clt *core.SDKClient, auth *model.RequestHeader, reportRequest *report.GetReportDataRequest) (*model.ResponseHeader, []report.ReportData, error) {
	reportRequest.ReportType = 2094816
	return reportApi.GetReportData(clt, auth, reportRequest)
}
