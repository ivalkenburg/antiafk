package main

import (
	"context"
	"fmt"
	"math/rand"
	"os/signal"
	"syscall"
	"time"

	"github.com/micmonay/keybd_event"
)

var keys = [...]int{
	keybd_event.VK_W,
	keybd_event.VK_A,
	keybd_event.VK_S,
	keybd_event.VK_D,
	keybd_event.VK_SPACE,
}

func main() {
	ctx, _ := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)

	go func() {
		kb, err := keybd_event.NewKeyBonding()
		if err != nil {
			panic(err)
		}

		for {
			time.Sleep(time.Second * 5)

			rand.Seed(time.Now().Unix())
			kb.SetKeys(keys[rand.Intn(len(keys))])

			fmt.Println("Pressing key...")

			kb.Press()
			time.Sleep(time.Second / 4)
			kb.Release()

			kb.Clear()
		}
	}()

	<-ctx.Done()
	fmt.Println("Exiting")
}
