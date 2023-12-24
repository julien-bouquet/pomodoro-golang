package main

import (
	"fmt"
	"log/slog"
	"os"
	"time"

	"github.com/caseymrm/menuet"
	"github.com/gen2brain/beeep"
)

const (
	workDurationSession             = 25 * time.Minute
	shortBreakDuration              = 5 * time.Minute
	longBreakDuration               = 20 * time.Minute
	countWorkSessionBeforeLongBreak = 4
	AppName                         = "Pomodoro"
)

func main() {
	slog.Info(fmt.Sprintf("%s start", AppName))
	defer slog.Info(fmt.Sprintf("%s stop", AppName))
	countWorkSession := 0
	go timerMenu()
	app := menuet.App()
	app.Name = "Not a Fan"
	app.Label = "com.github.caseymrm.notafan"
	app.RunApplication()
	for {
		runWorkSession()
		countWorkSession = (countWorkSession + 1) % (countWorkSessionBeforeLongBreak + 1)
		runBreakSession(countWorkSession)
	}

}

// Run work session with notification at the beginning and at the end
func runWorkSession() {
	notificationMessage := fmt.Sprintf("Focus start for %s", workDurationSession)
	err := beeep.Notify(AppName, notificationMessage, "assets/information.png")
	defer beeep.Notify(AppName, "Take a break", "assets/information.png")
	if err != nil {
		slog.Error("Error notify: ", err)
		os.Exit(1)
	}
	sleep(workDurationSession, "Focus for %s")
}

// Run break session with notification at the beginning and at the end
func runBreakSession(countWorkSession int) {
	breakDuration := shortBreakDuration
	if countWorkSession == countWorkSessionBeforeLongBreak {
		breakDuration = longBreakDuration
	}
	notificationMessage := fmt.Sprintf("Break start for %s", breakDuration)
	err := beeep.Notify(AppName, notificationMessage, "assets/information.png")
	defer beeep.Notify(AppName, "Go back to work", "assets/information.png")
	if err != nil {
		slog.Error("Error notify: ", err)
		os.Exit(1)
	}
	sleep(breakDuration, "Break for %s")
}

// Sleep and logs the remaining time
func sleep(waitDuration time.Duration, message string) {
	for waitDuration > 0 {
		slog.Info(fmt.Sprintf(message, waitDuration))
		time.Sleep(time.Second)
		waitDuration -= time.Second
	}
}

func timerMenu() {
	for {
		menuTitle := fmt.Sprintf("%s - %s", AppName, time.Now().Format("15:04:05"))
		menuet.App().SetMenuState(&menuet.MenuState{
			Title: menuTitle,
		})
	}
}
