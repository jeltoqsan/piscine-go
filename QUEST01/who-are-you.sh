curl -s 'https://raw.githubusercontent.com/kigiri/superhero-api/master/api/all.json' | jq '.[52].name' 

# curl -s 'https://raw.githubusercontent.com/kigiri/superhero-api/master/api/all.json' | jq '.[].name' - выводит весь список имен
# curl - передает данные от указанной ссылки; 
# 52 - имя бэтман по счету 51ый и 52 указан, 
# -s - убирает/скрывает лишние данные, выведенные на экран.
# jq - разбор/поиск строк и обращение к нужному полю/параметру

# второй способ url=https://raw.githubusercontent.com/kigiri/superhero-api/master/api/all.json
# curl -s ${url} | jq '.[51].name' 