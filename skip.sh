ls -l | gsed -n '1~2p'

#ls -l | sed -n '1,${p;n;n;n;}'

# ls -l Вывод подробной информации о файлах 