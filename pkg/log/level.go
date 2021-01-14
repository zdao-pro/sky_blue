package log

//Level log level
type Level int

const (
	_debugLevel Level = iota
	_infoLevel
	_warnLevel
	_accessLevel
	_errorLevel
	_fetalLevel
	purple = "\x1b[1;35m" //紫色
	yellow = "\x1b[1;33m"
	red    = "\x1b[97;41m"
	white  = "\x1b[0;00m"
	green  = "\x1b[1;32m"
)

var levelName = [...]string{
	_debugLevel:  "DEBUG",
	_infoLevel:   "INFO",
	_warnLevel:   "WARN",
	_accessLevel: "ACCESS",
	_errorLevel:  "ERROR",
	_fetalLevel:  "FETAL",
}

var prefixD = [...]D{
	_debugLevel:  KVString("debug", white),
	_infoLevel:   KVString("info", green),
	_warnLevel:   KVString("warn", purple),
	_accessLevel: KVString("access", ""),
	_errorLevel:  KVString("error", yellow),
	_fetalLevel:  KVString("fetal", red),
}

var tailD = [...]D{
	_debugLevel:  KVString("debug", "\x1b[0m\n"),
	_infoLevel:   KVString("info", "\x1b[0m\n"),
	_warnLevel:   KVString("warn", "\x1b[0m\n"),
	_accessLevel: KVString("access", "\n"),
	_errorLevel:  KVString("error", "\x1b[0m\n"),
	_fetalLevel:  KVString("fetal", "\x1b[0m\n"),
}

func (l Level) String() string {
	return levelName[l]
}
