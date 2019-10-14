curl -s https://raw.githubusercontent.com/kigiri/superhero-api/master/api/all.json | jq --arg HERO_ID "$HERO_ID" '.[] |select (.id == ($HERO_ID|tonumber)) |.connetions.relatives' | sed 's/"//g'
