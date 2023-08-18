# Установка пути к gcc
# $Env:PATH += ";C:\TDM-GCC-64\bin" 
# $Env:CONFIG_PATH = "./config/local.yaml"
# # Включаем CGO
# $Env:CGO_ENABLED = 1

swag init

# Собираем бинарник -gcflags "-N -l"
go build -o main.exe ./main.go

# Запускаем приложение
.\main.exe