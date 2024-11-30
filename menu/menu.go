package menu

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"

	"github.com/atotto/clipboard"
	"github.com/getlantern/systray"
	"github.com/tech-thinker/stikky/config"
	"github.com/tech-thinker/stikky/res"
	"github.com/tech-thinker/stikky/tasks"
)

func OnReady() {
	cfg := config.NewAppConfig()
	tasks := tasks.NewTask(cfg)
	ctx := context.Background()

	// Set the tray icon
	goos := runtime.GOOS
	if goos == "windows" {
		iconData, err := res.GetIconIco()
		if err != nil {
			log.Fatalf("Failed to load icon: %v", err)
		}
		systray.SetIcon(iconData)
	} else {
		iconData, err := res.GetIcon()
		if err != nil {
			log.Fatalf("Failed to load icon: %v", err)
		}
		systray.SetIcon(iconData)
	}

	// Set the tray title and tooltip
	systray.SetTitle("")
	systray.SetTooltip("Stikky")

	// Add menu items
	mBase64 := systray.AddMenuItem("Base64", "")
	mBase64Encode := mBase64.AddSubMenuItem("Encode", "")
	mBase64Decode := mBase64.AddSubMenuItem("Decode", "")
	mEncryption := systray.AddMenuItem("Encryption", "")
	mEncryptionEncrypt := mEncryption.AddSubMenuItem("Encrypt", "")
	mEncryptionDecrypt := mEncryption.AddSubMenuItem("Decrypt", "")
	mUUIDGen := systray.AddMenuItem("Generate UUID", "")

	mOpen := systray.AddMenuItem("Open", "Open the application")

	mQuit := systray.AddMenuItem("Quit", "Quit the application")

	// Example dynamic submenu items
	// subMenuItems := []string{"Item 1", "Item 2", "Item 3"}

	// // Create a submenu
	// subMenu := systray.AddMenuItem("Dynamic Submenu", "A submenu with dynamic items")
	// for _, item := range subMenuItems {
	// 	subItem := subMenu.AddSubMenuItem(item, fmt.Sprintf("Action for %s", item))
	// 	go func(item string) {
	// 		for {
	// 			select {
	// 			case <-subItem.ClickedCh:
	// 				// Action to perform when the submenu item is clicked
	// 				fmt.Printf("Clicked on %s\n", item)
	// 				subMenu.AddSubMenuItem("Another", "Another item")
	// 			}
	// 		}
	// 	}(item)
	// }

	// Handle menu item clicks
	go func() {
		for {
			select {
			case <-mBase64Encode.ClickedCh:
				clip, _ := clipboard.ReadAll()
				encoded, _ := tasks.Base64Encode(ctx, clip)
				clipboard.WriteAll(encoded)
			case <-mBase64Decode.ClickedCh:
				clip, _ := clipboard.ReadAll()
				decoded, _ := tasks.Base64Decode(ctx, clip)
				clipboard.WriteAll(decoded)
			case <-mEncryptionEncrypt.ClickedCh:
				clip, _ := clipboard.ReadAll()
				encoded, _ := tasks.Encrypt(ctx, clip)
				clipboard.WriteAll(encoded)
			case <-mEncryptionDecrypt.ClickedCh:
				clip, _ := clipboard.ReadAll()
				decoded, _ := tasks.Decrypt(ctx, clip)
				clipboard.WriteAll(decoded)
			case <-mUUIDGen.ClickedCh:
				uuid, _ := tasks.UUIDGenerate(ctx)
				clipboard.WriteAll(uuid)
			case <-mOpen.ClickedCh:
				self, err := os.Executable()
				fmt.Println(err)
				cmd := exec.Command(self, "--ui")
				err = cmd.Run()
				fmt.Println(err)
			case <-mQuit.ClickedCh:
				systray.Quit()
				os.Exit(0)
			}
		}
	}()
}

func OnExit() {
	// Cleanup logic when the tray icon is removed
	println("Systray exiting...")
}
