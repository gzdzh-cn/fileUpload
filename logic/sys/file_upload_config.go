package sys

import (
	"dzhgo/addons/fileUpload/dao"
	"dzhgo/addons/fileUpload/model"
	"dzhgo/addons/fileUpload/model/entity"
	"dzhgo/addons/fileUpload/service"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"

	"github.com/gzdzh/dzhcore"
)

type sFileUploadConfigService struct {
	*dzhcore.Service
}

func NewsFileUploadConfigService() *sFileUploadConfigService {
	return &sFileUploadConfigService{
		Service: &dzhcore.Service{
			Dao:   &dao.AddonsFileUploadConfig,
			Model: model.NewFileUploadConfig(),
			PageQueryOp: &dzhcore.QueryOp{
				ModifyResult: func(ctx g.Ctx, data interface{}) interface{} {

					type Pagination struct {
						Page  int `json:"page"`
						Size  int `json:"size"`
						Total int `json:"total"`
					}
					type PageData struct {
						List       []*entity.AddonsFileUploadConfig `json:"list"`
						Pagination *Pagination                      `json:"pagination"`
					}

					list := gconv.Map(data)["list"]
					if len(gconv.SliceAny(list)) > 0 {
						pageData := &PageData{}
						_ = gconv.Struct(data, pageData)
						for _, row := range pageData.List {
							task := service.TaskManager().GetTask(row.ItemId)
							if task != nil {
								row.Percent = task.Percent
							}
						}
						data = pageData
					}

					return data
				},
			},
		},
	}
}
