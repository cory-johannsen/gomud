#!/bin/bash

create_background_trait() {
  name=$1
  description=$2
  effect=$3
  filename=$4
cat <<EOL > "traits/$filename"
name: $name
description: $description
effects:
  - $effect
EOL
}
echo "generating background traits"
while read -r line; do
  IFS='|' read -r -a parts <<< "$line"
  name=$(echo "${parts[0]}" | awk '{$1=$1; print}')
  description=$(echo "${parts[1]}" | awk '{$1=$1; print}')
  effect=$(echo "${parts[2]}" | awk '{$1=$1; print}')
  filename=$(echo "$name" | tr '[:upper:]' '[:lower:]' | tr ' ' '_').yaml
  echo "creating $name"
  create_background_trait "$name" "$description" "$effect", "$filename"
done < background_traits.txt


