curl -s https://raw.githubusercontent.com/kigiri/superhero-api/master/api/all.json | jq --arg ID "$HERO_ID" '.[] |select (.id == ($ID|tonumber)) |.connetions.relatives' | sed 's/"//g'
