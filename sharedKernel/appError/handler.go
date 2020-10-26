package appError

//var slack = notification.NewSlack()

func Handle(err error) AppError {
	e, ok := err.(AppError)
	if !ok {
		return NewApi(InternalServer(), 500)
	}
	if e.GetType() == api {
		return e
	}
	
	return NewApi(InternalServer(), 500)
}
