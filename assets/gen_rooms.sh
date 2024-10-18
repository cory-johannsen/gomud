#!/bin/bash

create_room() {
  name=$1
  description=$2
  shift 2
  exits=("$@")
  filename=$(echo "$name" | tr '[:upper:]' '[:lower:]' | tr ' ' '_').yaml

cat <<EOL > "rooms/$filename"
name: $name
description: $description
exits:
EOL
  for exit in "${exits[@]}"; do
    IFS=':' read -r -a parts <<< "$exit"
    direction=$(echo "${parts[0]}" | awk '{$1=$1; print}')
    room=$(echo "${parts[1]}" | awk '{$1=$1; print}')
    echo "  $direction:" >> "rooms/$filename"
    echo "    direction: $direction" >> "rooms/$filename"
    echo "    name: $room" >> "rooms/$filename"
    echo "    description: $room" >> "rooms/$filename"
    echo "    target: $room" >> "rooms/$filename"
  done
}

echo "generating rooms"
rm -rf rooms/*.yaml
rooms=$(ls -l rooms/*.txt | awk '{print $9}')
for room in $rooms; do
  echo "processing $room"
  while read -r line; do
      IFS='|' read -r -a parts <<< "$line"
      name=$(echo "${parts[0]}" | awk '{$1=$1; print}')
      description=$(echo "${parts[1]}" | awk '{$1=$1; print}')
      exits=("${parts[@]:2}")

      create_room "$name" "$description" "${exits[@]}"
    done < "$room"
done
