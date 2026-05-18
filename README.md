\# VortexCLI 🌀



\*\*VortexCLI\*\* is a high-performance, cross-platform dependency management tool written in \*\*Go\*\*. It provides a unified, interactive interface to manage updates for both Python (`pip`) and Web (`npm`) projects.



Stop switching between different package managers. One command to rule them all.





\## 🚀 Features



\- \*\*Unified Interface\*\*: Use a single tool to check and update dependencies across different ecosystems.

\- \*\*Interactive Shell\*\*: A beautiful, custom-built interactive CLI with its own command loop.

\- \*\*Smart Scanning\*\*: Automatically detects `package.json` or `requirements.txt` before running commands.

\- \*\*Dependency Insight\*\*: Lists found modules before performing update scans.

\- \*\*Cross-Platform\*\*: Native support for Windows, Linux, and macOS (Intel/M1/M2).

\- \*\*Fast \& Lightweight\*\*: Built with Go for maximum performance and zero external dependencies.



\## 🖥️ Preview



┌────────────────────────────────────────────────────────────────────────────┬─────────────────────────────────────────────────────────────────────────────────┐

│██╗   ██╗ ██████╗ ██████╗ ████████╗███████╗██╗  ██╗       ██████╗██╗     ██╗│                              Commands:                                          │

│██║   ██║██╔═══██╗██╔══██╗╚══██╔══╝██╔════╝╚██╗██╔╝      ██╔════╝██║     ██║│           ^vort check <web/py> - Check for updates                              │

│██║   ██║██║   ██║██████╔╝   ██║   █████╗   ╚███╔╝       ██║     ██║     ██║│           ^vort get <web/py>   - Update all modules                             │

│╚██╗ ██╔╝██║   ██║██╔══██╗   ██║   ██╔══╝   ██╔██╗       ██║     ██║     ██║│           ^vort all            - Full system check                              │

│ ╚████╔╝ ╚██████╔╝██║  ██║   ██║   ███████╗██╔╝ ██╗██╗   ╚██████╗███████╗██║│           ^vort exit           - Leave the program                              │

│  ╚═══╝   ╚═════╝ ╚═╝  ╚═╝   ╚═╝   ╚══════╝╚═╝  ╚═╝╚═╝    ╚═════╝╚══════╝╚═╝│                                                                                 │

└────────────────────────────────────────────────────────────────────────────┴─────────────────────────────────────────────────────────────────────────────────┘

🛠️ Commands

You can use VortexCLI in Interactive Mode or as a Direct Command.



Command	Action

^vort check py	Scans requirements.txt and checks for outdated Python modules via pip.

^vort check web	Scans package.json and checks for outdated JS modules via npm.

^vort get py	Upgrades all Python dependencies listed in requirements.txt.

^vort get web	Runs npm update to synchronize your web project dependencies.

^vort all	Performs a full scan of both Python and Web environments.

^vort exit	Safely closes the VortexCLI shell.

📦 Installation \& Build

Prerequisites

Go 1.18 or higher (only for building from source).

Python/Pip and Node.js/NPM installed in your system PATH.

Build from source

Bash



git clone https://github.com/YOUR\_USERNAME/VortexCLI.git

cd VortexCLI

go build -o vort.exe

Cross-Platform Compilation

VortexCLI can be compiled for other systems easily:



Linux: GOOS=linux GOARCH=amd64 go build -o vort-linux

macOS (M1/M2): GOOS=darwin GOARCH=arm64 go build -o vort-mac-arm

🤝 Contributing

Feel free to fork this project, report bugs, or submit pull requests. Every "Senior" started somewhere!



📜 License

This project is open-source and available under the MIT License.





