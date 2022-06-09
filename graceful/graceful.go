package graceful

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func Shutdown(foreground func(), cleanups ...func()) {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGINT)
	go func() {
		foreground()
	}()

	sig := <-sigChan
	fmt.Println("触发信号", sig)
	for i := 0; i < len(cleanups); i++ {
		cleanups[i]()
	}
}
