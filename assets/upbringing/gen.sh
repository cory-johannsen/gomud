#!/bin/bash

create_upbringing() {
  name="$1"
  stat="$2"
  description="$3"
  filename=$(echo "$name" | tr '[:upper:]' '[:lower:]' | tr ' ' '_').yaml
cat <<EOL > "$filename"
name: $name
stat: $stat
description: $description
EOL
}

filename="upbringing.txt"
echo "processing $filename"
while read -r line; do
  IFS='|' read -r -a parts <<< "$line"
  name=$(echo "${parts[0]}" | awk '{$1=$1; print}')
  stat=$(echo "${parts[1]}" | awk '{$1=$1; print}')
  description=$(echo "${parts[2]}" | awk '{$1=$1; print}')
  echo "creating $name"
  echo "    stat: $stat"
  echo "    desc: $description"
  create_upbringing "$name" "$stat" "$description"
done < "$filename"
