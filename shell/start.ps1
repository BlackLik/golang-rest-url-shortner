# Установка пути к gcc
$Env:PATH += ";C:\TDM-GCC-64\bin" 
$Env:CONFIG_PATH = "./config/local.yaml"
# Включаем CGO
$Env:CGO_ENABLED = 1

# Собираем бинарник 
go build -gcflags "-N -l" -o main.exe .\cmd\url-sortener\main.go

# Запускаем приложение
.\main.exe