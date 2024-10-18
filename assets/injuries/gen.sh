#!/bin/bash
severities=(
  "moderate"
  "serious"
  "grievous"
)



create_injury() {
  path=$1
  echo "creating to $path"
  severity=$(echo "$path" | awk '{print toupper(substr($0, 1, 1)) tolower(substr($0, 2))}')
  echo "creating $severity injury"
  name="$2"
  effect="$3"
  filename=$(echo "$name" | tr '[:upper:]' '[:lower:]' | tr ' ' '_').yaml
cat <<EOL > "$path/$filename"
name: $name
severity: $severity
effects:
  - $effect
EOL
}

process_injuries() {
  severity=$1
  filename="$severity.txt"
  echo "processing $filename"
  while read -r line; do
    IFS='|' read -r -a parts <<< "$line"
    name=$(echo "${parts[0]}" | awk '{$1=$1; print}')
    effect=$(echo "${parts[1]}" | awk '{$1=$1; print}')
    echo "creating $name"
    echo "    severity: $severity"
    echo "    effect: $effect"
    create_injury "$severity" "$name" "$effect"
  done < "$filename"
}

echo "generating injuries"
for severity in "${severities[@]}"; do
  echo "generating $severity injuries"
  rm -rf "$severity"
  mkdir -p "$severity"
  process_injuries "$severity"
done
