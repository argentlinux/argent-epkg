package ui

import (
	"epkg-go/pkg/search"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"strings"
)

func CreateAppLayout(w fyne.Window) fyne.CanvasObject {
	sidebar := container.NewVBox(
		widget.NewLabel("Discover"),
		widget.NewButton("Configure", func() {}),
		widget.NewButton("Develop", func() {}),
		widget.NewButton("Categories", func() {}),
		widget.NewButton("Updates", func() {}),
	)

	searchEntry := widget.NewEntry()
	searchEntry.SetPlaceHolder("Enter application to search")

	searchResultsLabel := widget.NewLabel("Results of epkg package search")
	searchResults := container.NewVBox()

	showInstallDialog := func(appName string) {
		confirmDialog := dialog.NewConfirm("Install Application",
			fmt.Sprintf("Do you want to install %s?", appName),
			func(confirm bool) {
				if confirm {
					output, err := search.RunSearch("epkg autoinstall --skip-check " + appName)
					if err != nil {
						// without these scrolls right here, the UI would be even more terrible than it is
						scrollableOutput := widget.NewLabel(output)
						scrollContainer := container.NewVScroll(scrollableOutput)
						scrollContainer.SetMinSize(fyne.NewSize(400, 300))
						customDialog := dialog.NewCustom("Installation Failed", "Close", scrollContainer, w)
						customDialog.Show()
					} else {
						scrollableOutput := widget.NewLabel(output)
						scrollContainer := container.NewVScroll(scrollableOutput)
						scrollContainer.SetMinSize(fyne.NewSize(400, 300))
						customDialog := dialog.NewCustom("Installation Output", "Close", scrollContainer, w)
						customDialog.Show()
					}
				}
			}, w)
		confirmDialog.Show()
	}

	searchButton := widget.NewButton("Search", func() {
		query := searchEntry.Text
		// completely unconventional to use this directly as fmt.Sprintf hardcoded
		// will implement structs another time, so we can properly use them instead of hardcoded
		command := fmt.Sprintf("epkg search %s", query)
		output, err := search.RunSearch(command)
		if err != nil {
			if strings.Contains(err.Error(), "exit code 1") {
				searchResultsLabel.SetText(fmt.Sprintf("Command failed: %s", output))
			} else {
				searchResultsLabel.SetText(fmt.Sprintf("Error: %s", err))
			}
		} else {
			results := strings.Split(output, "\n")
			searchResults.Objects = nil
			for _, result := range results {
				// one of the most unorthodox ways of selecting a package
				// in the future I'll have a separate package with proper package information returns
				// this if checks if the return string result from the above iteration contains "asterisk"
				// if yes, it's our program, and we trim the first character (which would be our asterisk) out
				if result != "" && strings.HasPrefix(result, "*") {
					appName := strings.TrimSpace(result[1:])
					appLink := widget.NewHyperlink(appName, nil)

					// declaring a button that we call "install"
					// this button starts a yes/no dialog defined above as showInstallDialog(parameter)
					installButton := widget.NewButton("Install", func() {
						showInstallDialog(appName)
					})

					// this positions the newly created button on the right side of the asterisk name, same size as the other buttons
					installButton.Resize(fyne.NewSize(80, 40))
					appHeader := container.NewHBox(
						appLink,
						layout.NewSpacer(),
						installButton,
					)
					searchResults.Add(appHeader)
				} else if result != "" {
					searchResults.Add(widget.NewLabel(result))
				}
			}
			searchResults.Refresh()
		}
	})

	searchBar := container.NewVBox(
		widget.NewLabel("Search Packages"),
		searchEntry,
		searchButton,
	)

	fullSidebar := container.NewVBox(
		searchBar,
		widget.NewLabel("Navigation"),
		sidebar,
	)

	scrollableResults := container.NewVScroll(searchResults)
	scrollableResults.SetMinSize(fyne.NewSize(600, 400))

	content := container.NewVBox(
		searchResultsLabel,
		scrollableResults,
	)

	mainLayout := container.NewBorder(nil, nil, fullSidebar, nil, content)

	return mainLayout
}
