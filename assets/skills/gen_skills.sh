#!/bin/bash

create_skill() {
  name=$1
  stat=$2
  type=$3
  focuses=("${@:4}")
  filename=$(echo "$name" | tr '[:upper:]' '[:lower:]' | tr ' ' '_').yaml
  echo "Creating skill: $name"
  echo "  Stat: $stat"
  echo "  Type: $type"
  echo "  Focuses: ${focuses[@]}"
cat <<EOL > "$filename"
name: $1
description: "This is a description of the skill."
stat: $stat
type: $type
focuses:
EOL
  for focus in "${focuses[@]}"; do
    trimmed="${focus%,}"
    echo "  - $trimmed" >> "$filename"
  done
}

while read -r line; do
  # Process each line as an array of strings
  echo "Processing line: $line"
  IFS=$' ' read -r -a skill <<< "$line"

  name=${skill[0]}
  stat=${skill[1]}
  type=${skill[2]}
  focuses=(${skill[@]:3})

  create_skill "$name" "$stat" "$type" "${focuses[@]}"
done < skills_list.txt