package app

import "github.com/leebrouse/GoMcp/internal/llm/service"

type ResearchAppService struct {
	svc service.ResearchService
}

func (r *ResearchAppService) Run(query string) string {
	// 调用领域服务
	return "research result"
}
