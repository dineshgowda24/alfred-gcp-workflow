package orchestrator

var _ Handler = (*ErrorHandler)(nil)

type ErrorHandler struct{}

func (e *ErrorHandler) Handle(ctx *Context) error {
	wf := ctx.Workflow
	if ctx.Err != nil {
		wf.NewItem(ctx.Err.Error()).
			Subtitle("Please check the logs for more details").
			Valid(false)
		wf.SendFeedback()
	}
	return nil
}
