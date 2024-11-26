package main

import (
	"github.com/getlantern/systray"
	"github.com/tech-thinker/stikky/menu"
)

func main() {
	runSystray()
}

func runSystray() {
	systray.Run(menu.OnReady, menu.OnExit)
}

// func onReady() {
// 	// Set the tray icon
// 	systray.SetIcon(getIcon("icon/icon.png")) // Replace with your icon path
// 	// systray.SetIcon(theme.FyneLogo().Content())

// 	// Set the tray title and tooltip
// 	systray.SetTitle("Stikky")
// 	systray.SetTooltip("Simple data management app.")

// 	// Add menu items
// 	mOpen := systray.AddMenuItem("Open App", "Open the main application")
// 	mSave := systray.AddMenuItem("Save Context", "Save from context.")
// 	mSetContext := systray.AddMenuItem("Set Context", "Set from context.")
// 	mQuit := systray.AddMenuItem("Quit", "Quit the application")

// 	// Example dynamic submenu items
// 	subMenuItems := []string{"Item 1", "Item 2", "Item 3"}

// 	// Create a submenu
// 	subMenu := systray.AddMenuItem("Dynamic Submenu", "A submenu with dynamic items")
// 	for _, item := range subMenuItems {
// 		subItem := subMenu.AddSubMenuItem(item, fmt.Sprintf("Action for %s", item))
// 		go func(item string) {
// 			for {
// 				select {
// 				case <-subItem.ClickedCh:
// 					// Action to perform when the submenu item is clicked
// 					fmt.Printf("Clicked on %s\n", item)
// 					subMenu.AddSubMenuItem("Another", "Another item")
// 				}
// 			}
// 		}(item)
// 	}

// 	// Handle menu item clicks
// 	go func() {
// 		for {
// 			select {
// 			case <-mOpen.ClickedCh:
// 				// Open app logic (e.g., bring Fyne app window to the front)
// 				// Custom logic can go here
// 				// Example: Logging
// 				println("Open App clicked!")
// 			case <-mSave.ClickedCh:
// 				content, _ := clipboard.ReadAll()
// 				fmt.Println(content)
// 			case <-mSetContext.ClickedCh:
// 				clipboard.WriteAll("This is new context!")
// 			case <-mQuit.ClickedCh:
// 				systray.Quit()
// 				os.Exit(0)
// 			}
// 		}
// 	}()
// 	go func() {
// 		for {
// 			date := time.Now().String()
// 			systray.SetTitle(date)
// 			time.Sleep(1 * time.Second)
// 		}
// 	}()
// }

// func onExit() {
// 	// Cleanup logic when the tray icon is removed
// 	println("Systray exiting...")
// }

// // getIcon reads an icon file and returns its bytes
// func getIcon(filePath string) []byte {
// 	data, err := os.ReadFile(filePath)
// 	if err != nil {
// 		panic(err)
// 	}
// 	return data
// }
