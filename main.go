package main

import (
	"fmt"
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"
	time2 "time"
)

func main() {
	var weatherInfo WeatherInfo
	getWeatherForecast(&weatherInfo)

	setupUi(weatherInfo)
}

func setupUi(weatherInfo WeatherInfo) {
	app := app.New()

	w := app.NewWindow("Программа для просмотра погоды")

	var vBox = widget.NewVBox()

	for i := 0; i < len(weatherInfo.List); i++ {
		var weatherForDay = weatherInfo.List[i]
		var weatherMainGroup = widget.NewVBox(
			widget.NewLabel(fmt.Sprintf("Температура: %.2f °C", weatherForDay.Main.Temp)),
			widget.NewLabel(fmt.Sprintf("Ощущается как: %.2f °C", weatherForDay.Main.FeelsLike)),
			widget.NewLabel(fmt.Sprintf("Влажность: %d%%", weatherForDay.Main.Humidity)),
		)

		var weatherTypeGroup = widget.NewVBox()
		for weatherTypeI := 0; weatherTypeI < len(weatherForDay.Weather); weatherTypeI++ {
			var resource, _ = fyne.LoadResourceFromURLString(fmt.Sprintf("http://openweathermap.org/img/wn/%s.png", weatherForDay.Weather[weatherTypeI].Icon))
			var icon = widget.NewIcon(resource)
			weatherTypeGroup.Append(icon)
		}

		var time = time2.Unix(int64(weatherInfo.List[i].Dt), 0).String()
		vBox.Append(widget.NewGroup(time))
		vBox.Append(widget.NewHBox(weatherMainGroup, weatherTypeGroup))
	}
	vBox.Append(widget.NewButton("Закрыть", func() {
		app.Quit()
	}))

	w.SetContent(vBox)

	w.ShowAndRun()
}
