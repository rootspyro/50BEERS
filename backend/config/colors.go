package config

type ColorsConf struct {
	Reset  string
	Red    string
	Green  string
	Yellow string
	Blue   string
	Purple string
	Cyan   string
}

var Colors ColorsConf = ColorsConf{
	Reset : "\033[0m",
	Red : "\033[31m",
	Green : "\033[32m",
	Yellow : "\033[33m",
	Blue : "\033[34m",
	Purple : "\033[35m",
	Cyan : "\033[36m",
}
