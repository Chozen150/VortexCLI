package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

type PackageData struct {
	Dependencies map[string]string `json:"dependencies"`
}

const (
	ColorReset  = "\033[0m"
	ColorGreen  = "\033[32m"
	ColorYellow = "\033[33m"
	ColorRed    = "\033[31m"
	ColorCyan   = "\033[36m"
)

const vortLogo = `
┌────────────────────────────────────────────────────────────────────────────┬─────────────────────────────────────────────────────────────────────────────────┐
│██╗   ██╗ ██████╗ ██████╗ ████████╗███████╗██╗  ██╗       ██████╗██╗     ██╗│                              Commands:                                          │
│██║   ██║██╔═══██╗██╔══██╗╚══██╔══╝██╔════╝╚██╗██╔╝      ██╔════╝██║     ██║│           ^vort check <web/py> - Check have it an updates                       │
│██║   ██║██║   ██║██████╔╝   ██║   █████╗   ╚███╔╝       ██║     ██║     ██║│           ^vort exit - leave the program                                        │
│╚██╗ ██╔╝██║   ██║██╔══██╗   ██║   ██╔══╝   ██╔██╗       ██║     ██║     ██║│                                                                                 │
│ ╚████╔╝ ╚██████╔╝██║  ██║   ██║   ███████╗██╔╝ ██╗██╗   ╚██████╗███████╗██║│                                                                                 │
│  ╚═══╝   ╚═════╝ ╚═╝  ╚═╝   ╚═╝   ╚══════╝╚═╝  ╚═╝╚═╝    ╚═════╝╚══════╝╚═╝│                                                                                 │
└────────────────────────────────────────────────────────────────────────────┴─────────────────────────────────────────────────────────────────────────────────┘

[^VortexCLI] Enter the command:`

func getPythonModules() []string {
	file, err := os.Open("requirements.txt")
	if err != nil {
		return nil
	}
	defer file.Close()

	var modules []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		parts := strings.FieldsFunc(line, func(r rune) bool {
			return r == '=' || r == '>' || r == '<' || r == '~'
		})
		if len(parts) > 0 {
			modules = append(modules, strings.TrimSpace(parts[0]))
		}
	}
	return modules
}

func getWebModules() []string {
	file, err := os.ReadFile("package.json")
	if err != nil {
		return nil
	}

	var data PackageData
	err = json.Unmarshal(file, &data)
	if err != nil {
		return nil
	}

	var modules []string
	for name := range data.Dependencies {
		modules = append(modules, name)
	}
	return modules
}

func checkWeb() {
	if _, err := os.Stat("package.json"); os.IsNotExist(err) {
		fmt.Println(ColorYellow + "\n[^VortexCLI] package.json not found" + ColorReset)
		return
	} else {
		modules := getWebModules()
		fmt.Printf(ColorCyan+"\n[^VortexCLI] Found %d modules in package.json:\n"+ColorReset, len(modules))
		for _, m := range modules {
			fmt.Printf(ColorCyan+"  • %s\n"+ColorReset, m)
		}

		fmt.Println(ColorCyan + "\n[^VortexCLI] Scanning for updates..." + ColorReset)

		cmd := exec.Command("npm", "outdated")

		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		err := cmd.Run()
		if err != nil {
			fmt.Printf(ColorRed+"\n[^VortexCLI] Command failed: %v\n", err)
			fmt.Println("\n[^VortexCLI] Error: make sure you have installed NodeJS." + ColorReset)
		}

		fmt.Println(ColorGreen + "\n[^VortexCLI] Version check complete. Use '^vort get web' to update." + ColorReset)
	}
}

func checkPython() {
	if _, err := os.Stat("requirements.txt"); os.IsNotExist(err) {
		fmt.Println(ColorYellow + "\n[^VortexCLI] requirements.txt not found" + ColorReset)
		return
	} else {
		modules := getPythonModules()
		fmt.Printf(ColorCyan+"\n[^VortexCLI] Found %d modules in requirements.txt:\n"+ColorReset, len(modules))
		for _, m := range modules {
			fmt.Printf("  • %s\n", m)
		}

		fmt.Println(ColorCyan + "\n[^VortexCLI] Scanning for updates..." + ColorReset)

		cmd := exec.Command("pip", "list", "--outdated")

		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		err := cmd.Run()
		if err != nil {
			fmt.Printf(ColorRed+"\n[^VortexCLI] Command failed: %v\n", err)
			fmt.Println("\n[^VortexCLI] Error: make sure you have installed Python." + ColorReset)
		}
		fmt.Println(ColorGreen + "\n[^VortexCLI] Version check complete. Use '^vort get py' to update." + ColorReset)
	}
}

func runCheck(target string) {
	switch target {
	case "py":
		checkPython()
	case "web":
		checkWeb()
	default:
		fmt.Println(ColorYellow + "\n[^VortexCLI] That manager doesn't support in VortexCLI." + ColorReset)
	}
}

func getPython() {
	if _, err := os.Stat("requirements.txt"); os.IsNotExist(err) {
		fmt.Println(ColorYellow + "\n[^VortexCLI] requirements.txt not found" + ColorReset)
	} else {
		cmd := exec.Command("pip", "install", "--upgrade", "-r", "requirements.txt")

		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		err := cmd.Run()
		if err != nil {
			fmt.Printf(ColorRed+"\n[^VortexCLI] Command failed: %v\n", err)
			fmt.Println("\n[^VortexCLI] Error: make sure you have installed Python." + ColorReset)
		}
	}
}

func getWeb() {
	if _, err := os.Stat("package.json"); os.IsNotExist(err) {
		fmt.Println(ColorYellow + "\n[^VortexCLI] package.json not found" + ColorReset)
	} else {
		cmd := exec.Command("npm", "update")

		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		err := cmd.Run()
		if err != nil {
			fmt.Printf(ColorRed+"\n[^VortexCLI] Command failed: %v\n", err)
			fmt.Println("\n[^VortexCLI] Error: make sure you have installed NodeJS." + ColorReset)
		}
	}
}

func runGet(target string) {
	switch target {
	case "py":
		getPython()
	case "web":
		getWeb()
	default:
		fmt.Println(ColorYellow + "\n[^VortexCLI] That manager doesn't support in VortexCLI." + ColorReset)
	}
}

func main() {
	if len(os.Args) >= 3 {
		command := os.Args[1]
		target := os.Args[2]

		if command == "check" {
			runCheck(target)
		}
		return
	}

	fmt.Print(vortLogo)
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("\n	> ^vort ")

		if !scanner.Scan() {
			break
		}

		input := scanner.Text()
		parts := strings.Fields(input)

		if len(parts) == 0 {
			continue
		}

		var command string
		var args []string

		if parts[0] == "^vort" {
			if len(parts) < 2 {
				continue
			}
			command = parts[1]
			args = parts[2:]
		} else {
			command = parts[0]
			args = parts[1:]
		}

		switch command {
		case "check":
			if len(args) < 1 {
				fmt.Println(ColorYellow + "\n[^VortexCLI] Usage: check <web/py>" + ColorReset)
			} else {
				runCheck(args[0])
			}
		case "get":
			if len(args) < 1 {
				fmt.Println(ColorYellow + "\n[^VortexCLI] Usage: get <web/py>" + ColorReset)
			} else {
				runGet(args[0])
			}
		case "all":
			fmt.Println(ColorCyan + "[^VortexCLI] Full system check..." + ColorReset)
			checkPython()
			checkWeb()
		case "exit":
			return
		default:
			fmt.Printf(ColorRed+"\n[^VortexCLI] Unknown command: %s\n"+ColorReset, command)
		}
	}
}
