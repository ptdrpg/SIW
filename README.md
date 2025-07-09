# SIW – Human Resource Management System for Churches

**SIW** is a Human Resource Management System (HRMS) specifically designed for churches. Built with **Go**, **Wails**, **GORM**, and a **React frontend**, this application aims to streamline the management of church members, sacraments, and community areas (called *faritra*).

---

## ✨ Features

- 📋 Member registration and management
- 🎓 Sacramental record tracking (Baptism, Communion, Confession, Confirmation, etc.)
- 💍 Marriage & couple tracking
- 📅 Birth and death records
- 📤 Excel export of member data
- 🔒 Admin login system with token authentication
- ⚙️ Local SQLite database via GORM
- 🌐 Frontend built with React and TypeScript

---

## 🧰 Technologies Used

| Stack      | Tools                        |
|------------|------------------------------|
| Backend    | Go, GORM, SQLite        |
| Frontend   | React, TypeScript, Vite      |
| Desktop UI | Wails (Go + WebView)         |
| Others     | bcrypt, excelize (Excel lib) |

---

## 🚀 Getting Started

### 🛠 Prerequisites

- Go 1.21+
- Node.js (for frontend)
- Wails v2 (`wails doctor` to verify setup)

---

### 📦 Installation

```bash
# Clone the repository
git clone https://github.com/ptdrpg/SIW.git
cd SIW

# Install backend dependencies
go mod tidy

# Install frontend dependencies
cd frontend
npm install

# Build the for Linux
./buildLinux.sh

# Build the for Linux
./buildWin.sh

# Run the app in dev mode
wails dev
```
