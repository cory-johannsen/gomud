#!/bin/bash

create_area() {
  name=$1
  description=$2
  shift 2
  rooms=("$@")
  filename=$(echo "$name" | tr '[:upper:]' '[:lower:]' | tr ' ' '_').yaml
cat <<EOL > "areas/$filename"
name: $name
description: $description
rooms:
EOL
  for room in "${rooms[@]}"; do
    echo "  - $room" >> "areas/$filename"
  done
}

echo "generating areas"
rm -rf areas/*.yaml

areas=$(ls -l areas/*.txt | awk '{print $9}')
for area in $areas; do
  echo "processing $area"
  while read -r line; do
      IFS='|' read -r -a parts <<< "$line"
      name=$(echo "${parts[0]}" | awk '{$1=$1; print}')
      description=$(echo "${parts[1]}" | awk '{$1=$1; print}')
      rooms=("${parts[@]:2}")
      echo "creating $name"
      create_area "$name" "$description" "${rooms[@]}"
    done < "$area"
done
