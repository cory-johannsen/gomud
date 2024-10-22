#!/bin/bash

create_quality() {
  name="$1"
  effect="$2"
  filename=$(echo "$name" | tr '[:upper:]' '[:lower:]' | tr ' ' '_').yaml
cat <<EOL > "qualities/$filename"
name: $name
conditions:
  - Equipped
effects:
  - $effect
targets:
  - single
EOL
}

filename="qualities.txt"
echo "processing $filename"
while read -r line; do
  IFS=':' read -r -a parts <<< "$line"
  name=$(echo "${parts[0]}" | awk '{$1=$1; print}')
  effect=$(echo "${parts[1]}" | awk '{$1=$1; print}')
  echo "creating $name"
  echo "    effect: $effect"
  create_quality "$name" "$effect"
done < "$filename"

create_weapon() {
  name="$1"
  description="$2"
  load="$3"
  handling="$4"
  distance="$5"
  qualities="$6"
  type="$7"
  encumbrance="$8"
  cost="$9"
  skill="${10}"

  IFS=',' read -r -a qualities <<< "$qualities"
  level=""
  category=""
  IFS=' ' read -r level category <<< "$skill"
  filename=$(echo "$name" | tr '[:upper:]' '[:lower:]' | tr ' ' '_').yaml
cat <<EOL > "weapons/$filename"
weapon:
  category: $category
  level: $level
  type: $type
  name: $name
  description: $description
  handling: $handling
  distance: $distance
  skill: $skill
  target: Single Target
  encumbrance: $encumbrance
  cost: $cost
  load: $load
  qualities:
EOL
  for quality in "${qualities[@]}"; do
    q=$(echo "$quality" | awk '{$1=$1; print}')
    echo "    - $q" >> "weapons/$filename"
  done
}

filename="weapons.txt"
echo "processing $filename"
while read -r line; do
  IFS='|' read -r -a parts <<< "$line"
  name=${parts[0]}
  description=${parts[1]}
  load=${parts[2]}
  handling=${parts[3]}
  distance=${parts[4]}
  qualities=${parts[5]}
  type=${parts[6]}
  encumbrance=${parts[7]}
  cost=${parts[8]}
  skill=${parts[9]}
  echo "creating $name"
  echo "    description: $description"
  echo "    load: $load"
  echo "    handling: $handling"
  echo "    distance: $distance"
  echo "    qualities: $qualities"
  echo "    type: $type"
  echo "    encumbrance: $encumbrance"
  echo "    cost: $cost"
  echo "    skill: $skill"
  create_weapon "$name" "$description" "$load" "$handling" "$distance" "$qualities" "$type" "$encumbrance" "$cost" "$skill"
done < "$filename"