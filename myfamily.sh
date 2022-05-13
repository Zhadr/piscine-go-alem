curl -s https://01.alem.school/assets/superhero/all.json \
    | jq --arg MAN_ID "$HERO_ID" '.[] | select(.id==($MAN_ID | tonumber)) | .connections.relatives' | tr -d '""'
