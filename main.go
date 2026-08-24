package main

import (
	"fmt"
	"os"
	"strings"
)

type Widget struct {
	Title string
	Type  string
	Size  string
	Data  string
}

func main() {
	widgets := []Widget{
		{"CPU Usage", "gauge", "1x1", "67%"},
		{"Memory", "gauge", "1x1", "4.2/8 GB"},
		{"Requests/min", "sparkline", "2x1", "142,138,155,160,149"},
		{"Error Rate", "number", "1x1", "0.3%"},
		{"Uptime", "stat", "1x1", "99.97%"},
		{"Recent Logs", "table", "3x2", "12 entries"},
	}
	if len(os.Args) > 1 && os.Args[1] == "--json" {
		fmt.Println(`{"dashboard":"ops","widgets":%d}`, len(widgets))
		return
	}
	fmt.Println("+-------------------------------+")
	fmt.Println("|      OPS DASHBOARD            |")
	fmt.Println("+-------------------------------+")
	for _, w := range widgets {
		padding := 25 - len(w.Title)
		if padding < 0 {
			padding = 0
		}
		fmt.Printf("| %-25s %s |\n", w.Title, strings.Repeat(" ", padding))
		fmt.Printf("|   %-30s  |\n", w.Data)
		fmt.Println("|                               |")
	}
	fmt.Println("+-------------------------------+")
	fmt.Printf("  Widgets: %d | Refresh: 30s\n", len(widgets))
}