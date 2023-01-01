package utils

import (
	"fmt"
	"os"
	"time"

	"github.com/gookit/color"
)

func GetTime() string {
	return time.Now().Format("15:04:05")
}

func LogError(format string, content ...any) {
	color.Printf("<fg=white>[</><fg=magenta;op=bold>%s</><fg=white>] </><fg=red;op=bold>ERROR   </><fg=white>| </><fg=blue>%s</><fg=white></>\n", GetTime(), fmt.Sprintf(format, content...))
}

func LogSuccess(format string, content ...any) {
	color.Printf("<fg=white>[</><fg=magenta;op=bold>%s</><fg=white>] </><fg=green;op=bold>SUCCESS </><fg=white>| </><fg=blue>%s</><fg=white></>\n", GetTime(), fmt.Sprintf(format, content...))
}

func LogInfo(format string, content ...any) {
	color.Printf("<fg=white>[</><fg=magenta;op=bold>%s</><fg=white>] </><fg=yellow;op=bold>INFO    </><fg=white>| </><fg=blue>%s</><fg=white></>\n", GetTime(), fmt.Sprintf(format, content...))
}

func LogPanic(format string, content ...any) {
	color.Printf("<fg=white>[</><fg=magenta;op=bold>%s</><fg=white>] </><fg=red;op=bold>PANIC   </><fg=white>| </><fg=blue>%s</><fg=white></>\n", GetTime(), fmt.Sprintf(format, content...))
	time.Sleep(1 * time.Second)
	os.Exit(0)
}
