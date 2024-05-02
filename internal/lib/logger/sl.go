package sl

import "log/slog"

func Err(err error)slog.Attr{
	return slog.Attr{
		Key: "erroro",
		Value: slog.StringValue(err.Error()),
	}
}