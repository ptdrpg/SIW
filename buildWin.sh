cd frontend
npm run build
cd ..
CC=x86_64-w64-mingw32-gcc CGO_ENABLED=1 wails build -platform windows/amd64
