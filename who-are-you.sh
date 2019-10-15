curl -s 'https://raw.githubusercontent.com/kigiri/superhero-api/master/api/all.json' | jq '.[51].name' 

# curl -s 'https://raw.githubusercontent.com/kigiri/superhero-api/master/api/all.json' | jq '.[].name' - выводит весь список имен
# curl - передает данные от указанной ссылки; 
# 51 - имя бэтман по счету 51ый указан, 
# -s - убирает/скрывает лишние данные, выведенные на экран.
# jq - разбор/поиск строк и обращение к нужному полю/параметру

