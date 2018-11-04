package apperr

type Code string

const (
	// Common
	ServerError        Code = "server_error"
	DBError            Code = "db_error"
	RequestError       Code = "request_error"
	MaintenanceOnGoing Code = "maintenance_on_going"
	NotFound           Code = "not_found"
)

func (c Code) status() int {
	var status int = 500

	// Common
	switch c {
	case ServerError:
		status = 500
	case DBError:
		status = 500
	case RequestError:
		status = 400
	case MaintenanceOnGoing:
		status = 503
	case NotFound:
		status = 404
	}

	return status
}

func (c Code) message() string {
	var message string

	switch c {
	case ServerError:
		message = "サーバーエラー"
	case DBError:
		message = "アクセスが集中指定ます。しばらく経ってからお試しください。"
	case RequestError:
		message = "リクエストエラー"
	case MaintenanceOnGoing:
		message = "現在メンテナンスを行なっております。サイト復旧までしばらくお待ちください"
	case NotFound:
		message = "お探しのページが見つかりませんでした。"
	}
	return message
}
