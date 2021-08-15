package penlogger

import "github.com/coreos/go-systemd/journal"

func (l *Logger) outputJournal(msg map[string]interface{}) {
	var (
		data string
		prio = -1
		vars = convertVarsForJournal(msg)
	)
	if rawVal, ok := msg["priority"]; ok {
		if val, ok := rawVal.(Prio); ok {
			prio = int(val)
		}
	}
	if prio == -1 {
		prio = int(PrioInfo)
	}
	if rawData, ok := msg["data"]; ok {
		if val, ok := rawData.(string); ok {
			data = val
		}
	}
	if err := journal.Send(data, journal.Priority(prio), vars); err != nil {
		panic(err)
	}
}

func pollJournal() bool {
	return journal.Enabled()
}
