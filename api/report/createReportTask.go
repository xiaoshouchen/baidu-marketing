package report

import (
	"github.com/bububa/baidu-marketing/core"
	"github.com/bububa/baidu-marketing/model"
	"github.com/bububa/baidu-marketing/model/report"
)

// CreateReportTask 创建异步任务
func CreateReportTask(clt *core.SDKClient, auth *model.RequestHeader, createRequest *report.CreateReportTaskRequest) (*model.ResponseHeader, string, error) {
	req := &model.Request{
		Header: auth,
		Body:   createRequest,
	}
	var resp report.CreateReportTaskResponse
	header, err := clt.Do(req, &resp)
	return header, resp.Data.TaskId, err
}
