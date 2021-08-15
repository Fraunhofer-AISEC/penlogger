// +build !linux

package penlogger

func (l *Logger) outputJournal(msg map[string]interface{}) {
	panic("systemd journal is not available on this plattform")
}

func pollJournal() bool {
	panic("systemd journal is not available on this plattform")
}
