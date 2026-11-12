# Open Companion Listener v0.1.4

**Control your Windows PC using Bitfocus Companion, Stream Deck, or any other Companion-compatible control device.**

Open Companion Listener is a **free, open-source Windows application** that allows [Bitfocus Companion](https://bitfocus.io/companion) to send keyboard commands, mouse actions, and application launch requests to your computer over a local network.

Want to launch a program with one button, send a keyboard shortcut, scroll through a document, switch between applications, or run your own automation scripts?

**Open Companion Listener makes that possible.**

No subscriptions. No cloud accounts. No locked features. No license fees.

**100% free and open source under the MIT License.**

![Open Companion Listener Dashboard](https://github.com/user-attachments/assets/6da70034-3ae5-4462-8339-6f7085c11d63)

---

## 💡 What is Open Companion Listener?

If you've never used [Bitfocus Companion](https://bitfocus.io/companion), here's the basic idea.

Bitfocus Companion is software that lets you create programmable buttons to control other devices, applications, and services. It's commonly used with hardware such as the Elgato Stream Deck, but it supports other control methods too.

**Open Companion Listener adds Windows desktop control to those buttons.**

It works through two components:

- **Open Companion Listener** — a small application running on the Windows computer you want to control.
- **Open Companion Listener Companion Module** — a module installed in Bitfocus Companion that sends commands to the Windows application.

When you press a configured Companion button, the module sends a command over your private network. The Windows listener receives that command and performs the requested action.

Both components can run on the same computer or on separate computers connected to the same trusted network.

### Why would you use it?

Normally, controlling Windows applications requires reaching for your keyboard or mouse.

With Open Companion Listener, you can assign commonly used commands to dedicated buttons and execute them instantly.

It can be useful for:

- Controlling applications with custom keyboard shortcuts
- Creating one-button workflows
- Launching frequently used programs, files, and folders
- Automating repetitive desktop tasks
- Controlling a Windows computer from another location on your local network
- Building custom control surfaces with Bitfocus Companion
- Integrating Windows desktop actions into larger Companion workflows

You decide what each button does.

---

## 🎛️ Features

Open Companion Listener supports several types of desktop control.

| Feature | Description |
|---|---|
| ⌨️ **Keyboard shortcuts** | Send combinations such as Ctrl+C, Ctrl+V, Alt+Tab, or custom shortcuts |
| 🖱️ **Mouse buttons** | Simulate left, right, and middle mouse clicks |
| ↕️ **Mouse scrolling** | Scroll vertically or horizontally |
| ↔️ **Mouse movement** | Move the mouse pointer |
| 🚀 **Launch applications** | Start installed Windows programs |
| 📁 **Open files and folders** | Open documents, files, and folders |
| 🌐 **Open websites** | Launch URLs in your default browser |
| ⚙️ **Run scripts** | Launch scripts through installed interpreters or configured file associations |
| 🎛️ **Companion presets** | Choose from 38 optional ready-made presets |
| 🖥️ **Web dashboard** | Configure the listener through a local browser interface |
| 🔔 **System tray integration** | Keep the listener running without an open application window |
| 🔐 **Authenticated commands** | Use a pairing key and HMAC-SHA256 authentication |

### A practical example

Imagine having five buttons on your Stream Deck:

| Button | Action |
|---|---|
| **1 — Open App** | Launch your favorite Windows application |
| **2 — Save As** | Send Ctrl+Shift+S to the active application |
| **3 — Scroll Down** | Scroll through a document or webpage |
| **4 — Switch Window** | Send Alt+Tab to switch applications |
| **5 — Run Script** | Launch a PowerShell, batch, or AutoHotkey script |

Instead of using your keyboard and mouse for these tasks, you simply press the corresponding Companion button.

You can combine Open Companion Listener actions with other Companion actions to create more advanced workflows.

**Important:** Keyboard and mouse input is sent to the active Windows desktop session. Keyboard shortcuts generally affect whichever application currently has focus.

For entering entire text strings, you can use an external script such as AutoHotkey. A dedicated native Send Text action is not currently included.

---

## 🖥️ Windows application

Open Companion Listener is a lightweight, native Windows application.

It runs in the system tray and provides a browser-based dashboard for configuration.

### Windows listener features

- Native Windows x64 executable
- Runs under your normal Windows user account
- System tray operation
- Local web-based configuration dashboard
- Configurable network listening address and port
- Trusted Companion host configuration
- Pairing-key authentication
- Option to start automatically with Windows
- Controls for opening the interface, restarting the listener, and exiting the application

The default API port is **17992**.

The dashboard is accessed locally on the Windows computer.

No cloud service or separate account is required.

---

## 📥 Downloads

**For normal use, you only need two files: the Windows application and the Companion module.**

| Download | Description |
|---|---|
| **[Windows Application — v0.1.4](https://github.com/ninoleto/OpenCompanionListener/releases/download/v0.1.4/OpenCompanionListener-v0.1.4-Windows-x64.zip)** | The Windows listener executable and documentation |
| **[Companion Module — v0.1.4](https://github.com/ninoleto/OpenCompanionListener/releases/download/v0.1.4/community-openlistener-0.1.4.tgz)** | Ready-to-import Bitfocus Companion module |
| [Complete Source Code](https://github.com/ninoleto/OpenCompanionListener/releases/download/v0.1.4/OpenCompanionListener-v0.1.4-Source.zip) | Source files for developers and contributors |
| [Build Information](https://github.com/ninoleto/OpenCompanionListener/releases/download/v0.1.4/BUILD-INFO.json) | Build metadata and source revision |
| [SHA-256 Checksums](https://github.com/ninoleto/OpenCompanionListener/releases/download/v0.1.4/SHA256SUMS.txt) | File integrity verification |

### System requirements

**Windows application:**

- Windows 10 or Windows 11
- 64-bit (x64) computer
- Network access to your Companion installation

**Companion module:**

- Bitfocus Companion 4.1 or newer with the Node.js 22 module runtime
- Support for importing custom module packages

**You do not need to install Node.js or npm separately to use the prebuilt Companion module.**

The Windows listener does not require Node.js.

---

## 🚀 Installation and setup

Getting started requires installing the Windows listener and importing its module into Bitfocus Companion.

### Step 1 — Download and run the Windows listener

1. Download **OpenCompanionListener-v0.1.4-Windows-x64.zip**.
2. Extract the archive into a permanent folder on your Windows PC.
3. Run `OpenCompanionListener.exe`.
4. The listener starts in the Windows system tray.
5. Right-click the tray icon and select **Open GUI** to access the settings dashboard.

You do not need to install additional runtime software for normal use.

**Tip:** Keep the application in a permanent folder, especially if you want it to start automatically with Windows.

### Step 2 — Configure the listener

Open the dashboard and configure:

1. **Listen IP address** — the Windows network address on which the listener accepts commands.
2. **Trusted Companion IP** — the address of the computer running Bitfocus Companion.
3. **API port** — default: `17992`.
4. **Pairing key** — the secret used to authenticate commands from Companion.

Save your settings.

Your Windows firewall must allow the Companion computer to reach the configured API port.

**Security reminder:** Never share your pairing key or expose the listener's API to the public Internet.

### Step 3 — Install the Bitfocus Companion module

On the computer running Bitfocus Companion:

1. Download **community-openlistener-0.1.4.tgz**.
2. Open the Companion web interface.
3. Navigate to **Modules → Import module package**.
4. Select the downloaded `.tgz` file.
5. Import the module.

The module includes its required runtime dependencies.

**No command-line setup, separate Node.js installation, or developer folder is necessary.**

### Step 4 — Connect Companion to the Windows listener

Inside Bitfocus Companion:

1. Add a new **Open Companion Listener** connection.
2. Enter the IP address of your Windows computer.
3. Enter the API port configured in the listener (default `17992`).
4. Enter the listener's pairing key.
5. Save the connection.

Once the connection is established, Companion can send commands to the Windows listener.

### Step 5 — Create your first button

1. Open a button page in Companion.
2. Select an empty button.
3. Add an **Open Companion Listener** action.
4. Choose an action, such as a keyboard shortcut.
5. Configure the required options.
6. Press the button to test it.

You can also explore the **38 included optional presets** for ready-made button configurations.

**That's it! Your Windows computer can now receive commands from Bitfocus Companion.**

For additional setup information, see the [Companion Installation Guide](https://github.com/ninoleto/OpenCompanionListener/blob/main/docs/COMPANION-INSTALL.md).

---

## ✅ Testing and verification

Open Companion Listener has been tested in a real Windows and Bitfocus Companion environment.

### Confirmed functional tests

- Windows executable starts successfully
- System tray integration works
- Local configuration dashboard works
- Companion module imports through the Companion interface
- Existing Companion connection settings can be retained
- Companion successfully connects to the Windows listener
- Tested keyboard, mouse, and application-control buttons work
- Windows executable and dashboard display the approved application branding

### Automated verification

Development and release verification included:

- Go automated tests, including race checks
- Go static analysis (`go vet`)
- Companion module automated tests
- Windows executable build verification
- Embedded Windows icon resource checks
- Companion module package validation
- Release archive integrity checks
- Reproducible release builds
- SHA-256 checksum verification

**Testing does not guarantee that every configuration, action, or operating environment will work without issues.**

Not all 38 presets or every possible Windows configuration have been individually tested.

Full verification details:

**[Open Companion Listener v0.1.4 Validation Report](https://github.com/ninoleto/OpenCompanionListener/blob/main/docs/VALIDATION-v0.1.4.md)**

---

## 🔒 Security and limitations

Open Companion Listener is intended for **trusted local networks and private network overlays**.

It is not designed to be exposed directly to the public Internet.

### Important security information

- Commands use **HMAC-SHA256 authentication**.
- Communication is **not encrypted** by the listener itself.
- Keep the listener and Companion on a trusted private network.
- Restrict Windows firewall access to your trusted Companion host.
- Never expose the listener API or dashboard to the public Internet.
- Never share or publish your pairing key.
- Run the listener under a normal Windows user account, not as Administrator.
- Treat paired Companion hosts as trusted devices.

**A paired Companion host can request actions with the permissions of the Windows desktop user, including launching applications and scripts.**

### Windows limitations

Windows security and desktop behavior still apply.

For example:

- Keyboard shortcuts affect the application currently receiving input.
- Windows UAC prompts and protected desktop sessions may prevent simulated input.
- Some applications handle simulated keyboard or mouse input differently.
- Successfully requesting that Windows launch a program does not guarantee that the program completed its task.

### Unsigned executable

Open Companion Listener is currently distributed without a commercial code-signing certificate.

Windows may display a security warning when launching the application.

Download only from the official project repository and verify the published checksum.

Read the complete [Security Documentation](https://github.com/ninoleto/OpenCompanionListener/blob/v0.1.4/SECURITY.md) for additional details.

---

## 🔍 Verify your download

Each release file is accompanied by a published SHA-256 checksum.

You can use these checksums to verify that downloaded files match the official release artifacts.

1. Download [SHA256SUMS.txt](https://github.com/ninoleto/OpenCompanionListener/releases/download/v0.1.4/SHA256SUMS.txt).
2. Open PowerShell in the folder containing your downloaded file.
3. Run:

   ```powershell
   Get-FileHash -Algorithm SHA256 ".\OpenCompanionListener-v0.1.4-Windows-x64.zip"
   ```

4. Compare the result with the matching entry in `SHA256SUMS.txt`.

The verified v0.1.4 release artifacts were built from source commit:

`30be74c8841594d67a4aaf49ee5cc4f6db36c827`

The source code, release documentation, and validation records are publicly available on GitHub.

---

## 🛠️ Open source and development

Open Companion Listener is **free and open-source software released under the MIT License**.

You are welcome to:

- Use it for personal or professional projects
- Inspect and study the source code
- Modify it to suit your needs
- Contribute improvements and bug fixes
- Fork the project and develop your own version
- Redistribute it under the terms of the MIT License

The goal is simple: **make useful desktop automation tools available to everyone, without mandatory subscriptions or unnecessary restrictions.**

### 🤖 AI-assisted development — Yes, this is vibe-coded!

**Open Companion Listener is a vibe-coded project, developed with extensive assistance from ChatGPT and OpenAI Codex.**

The project was conceived, directed, tested, and maintained by a human creator, with AI-assisted tools used to write, modify, debug, and document the software.

This project is an example of how an individual can use modern AI tools to develop useful applications without relying on a traditional software development team.

We're open about how this software was created.

**AI-generated code is not automatically correct, secure, or reliable.**

AI-assisted development can introduce bugs, unexpected behavior, and security vulnerabilities. For that reason, the project includes automated tests, practical testing, and release verification.

However, no amount of testing can guarantee that software is completely free of defects.

The complete source code is available for public inspection, and feedback from experienced developers, security researchers, and other AI-assisted creators is welcome.

**Vibe-coded. Human-directed. Tested in the real world. Free and open source.**

### 🤝 Contributing and reporting issues

Found a bug? Have an idea for improvement? Want to contribute?

You're welcome to participate.

- **[GitHub Repository](https://github.com/ninoleto/OpenCompanionListener)**
- **[Developer Documentation](https://github.com/ninoleto/OpenCompanionListener/blob/v0.1.4/docs/DEVELOPMENT.md)**
- **[Report a Bug or Suggest a Feature](https://github.com/ninoleto/OpenCompanionListener/issues)**

When reporting a problem, please include:

- Your Windows version
- Your Bitfocus Companion version
- Your Open Companion Listener version
- A short description of the problem
- Steps to reproduce it

**Never include pairing keys, passwords, private IP configuration details, or other sensitive information in public reports.**

---

## ❤️ Free software, built to be shared

Open Companion Listener exists because controlling your own computer shouldn't require an expensive subscription.

If you find it useful, feel free to share it, improve it, contribute to its development, or build something new from its source code.

**No subscriptions. No cloud dependency. No paid feature unlocks.**

Just a practical tool for controlling your Windows desktop through Bitfocus Companion.

---

### Open Companion Listener v0.1.4

**Free. Local. Open source. AI-assisted.**

[**Download Open Companion Listener**](https://github.com/ninoleto/OpenCompanionListener/releases/tag/v0.1.4) | [Source Code](https://github.com/ninoleto/OpenCompanionListener) | [Report an Issue](https://github.com/ninoleto/OpenCompanionListener/issues)

**MIT Licensed — Made with ChatGPT, OpenAI Codex, and a lot of real-world testing.**
