package exception

type ItemCounting struct{}

func (e *ItemCounting) Error() string {
	return "failed to count items"
}
