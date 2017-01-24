package runner

type Scaler interface {
	ProcessStat(cpu, ram int, app string)
}
