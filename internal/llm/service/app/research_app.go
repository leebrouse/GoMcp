package app

type ResearchAppService struct {
}

func (r *ResearchAppService) Run(query string) string {
	// 调用领域服务
	return "research result"
}
