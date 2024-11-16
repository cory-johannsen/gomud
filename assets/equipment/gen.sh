#!/bin/bash

# Create all the qualities
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

# create the weapons

create_weapon() {
  name="$1"
  description="$2"
  load="$3"
  handling="$4"
  distance="$5"
  weaponQualities="$6"
  type="$7"
  encumbrance="$8"
  mass="$9"
  cost="${10}"
  skill="${11}"

  IFS=',' read -r -a weaponQualities <<< "$weaponQualities"
  level=""
  category=""
  IFS=' ' read -r level category <<< "$skill"
  filename=$(echo "$name" | tr '[:upper:]' '[:lower:]' | tr ' ' '_').yaml
cat <<EOL > "weapons/$filename"
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
mass: $mass
cost: $cost
load: $load
qualities:
EOL
  for quality in "${weaponQualities[@]}"; do
    q=$(echo "$quality" | awk '{$1=$1; print}')
    echo "  - $q" >> "weapons/$filename"
  done
}


rm -rf weapons
mkdir weapons
filename="weapons.txt"
echo "processing $filename"
while read -r line; do
  IFS='|' read -r -a parts <<< "$line"
  name=${parts[0]}
  description=${parts[1]}
  load=${parts[2]}
  handling=${parts[3]}
  distance=${parts[4]}
  wquals=${parts[5]}
  type=${parts[6]}
  encumbrance=${parts[7]}
  mass=${parts[8]}
  cost=${parts[9]}
  skill=${parts[10]}
  echo "creating $name"
  echo "    description: $description"
  echo "    load: $load"
  echo "    handling: $handling"
  echo "    distance: $distance"
  echo "    qualities: $wquals"
  echo "    type: $type"
  echo "    encumbrance: $encumbrance"
  echo "    mass: $mass"
  echo "    cost: $cost"
  echo "    skill: $skill"
  create_weapon "$name" "$description" "$load" "$handling" "$distance" "$wquals" "$type" "$encumbrance" "$mass" "$cost" "$skill"
done < "$filename"

# create the armor

create_armor() {
  name="$1"
  description="$2"
  damageThresholdModifier="$3"
  encumbrance="$4"
  mass="$5"
  cost="$6"
  armorQualities="$7"
  filename=$(echo "$name" | tr '[:upper:]' '[:lower:]' | tr ' ' '_').yaml
  IFS=',' read -r -a armorQualities <<< "$armorQualities"
cat <<EOL > "armor/$filename"
name: $name
description: $description
damageThresholdModifier: $damageThresholdModifier
encumbrance: $encumbrance
mass: $mass
cost: $cost
qualities:
EOL
for quality in "${armorQualities[@]}"; do
    q=$(echo "$quality" | awk '{$1=$1; print}')
    if [ -z "$q" ]; then
      continue
    fi
    echo "  - $q" >> "armor/$filename"
  done
}


rm -rf armor
mkdir armor
filename="armor.txt"
echo "processing $filename"
while read -r line; do
  IFS='|' read -r -a parts <<< "$line"
  name=${parts[0]}
  description=${parts[1]}
  damageThresholdModifier=${parts[2]}
  aquals=${parts[3]}
  encumbrance=${parts[4]}
  mass=${parts[5]}
  cost=${parts[6]}
  echo "creating $name"
  echo "    description: $description"
  echo "    damageThresholdModifier: $damageThresholdModifier"
  echo "    qualities: $aquals"
  echo "    encumbrance: $encumbrance"
  echo "    mass: $mass"
  echo "    cost: $cost"
  create_armor "$name" "$description" "$damageThresholdModifier" "$encumbrance" "$mass" "$cost" "$aquals"
done < "$filename"

# create the shields

create_shield() {
  name="$1"
  description="$2"
  handedness="$3"
  encumbrance="$4"
  mass="$5"
  cost="$6"
  shieldQualities="$7"
  filename=$(echo "$name" | tr '[:upper:]' '[:lower:]' | tr ' ' '_').yaml
  IFS=',' read -r -a shieldQualities <<< "$shieldQualities"
cat <<EOL > "shields/$filename"
name: $name
description: $description
handling: $handedness
encumbrance: $encumbrance
mass: $mass
cost: $cost
qualities:
EOL
for quality in "${shieldQualities[@]}"; do
    q=$(echo "$quality" | awk '{$1=$1; print}')
    if [ -z "$q" ]; then
      continue
    fi
    echo "  - $q" >> "shields/$filename"
  done
}

rm -rf shields
mkdir shields
filename="shields.txt"
echo "processing $filename"
#Name|Description|handedness|qualities|encumbrance|mass|cost
while read -r line; do
  IFS='|' read -r -a parts <<< "$line"
  name=${parts[0]}
  description=${parts[1]}
  handedness=${parts[2]}
  squals=${parts[3]}
  encumbrance=${parts[4]}
  mass=${parts[5]}
  cost=${parts[6]}
  echo "creating $name"
  echo "    description: $description"
  echo "    handedness: $handedness"
  echo "    qualities: $aquals"
  echo "    encumbrance: $encumbrance"
  echo "    mass: $mass"
  echo "    cost: $cost"
  create_shield "$name" "$description" "$handedness" "$encumbrance" "$mass" "$cost" "$squals"
done < "$filename"


# create the misc equipment

create_misc() {
  name="$1"
  description="$2"
  encumbrance="$4"
  mass="$5"
  cost="$6"
  filename=$(echo "$name" | tr '[:upper:]' '[:lower:]' | tr ' ' '_').yaml
cat <<EOL > "misc/$filename"
name: $name
description: $description
encumbrance: $encumbrance
mass: $mass
cost: $cost
EOL
}

rm -rf misc
mkdir misc
filename="misc.txt"
echo "processing $filename"
#Name|Description|encumbrance|mass|cost
while read -r line; do
  IFS='|' read -r -a parts <<< "$line"
  name=${parts[0]}
  description=${parts[1]}
  encumbrance=${parts[2]}
  mass=${parts[3]}
  cost=${parts[4]}
  echo "creating $name"
  echo "    description: $description"
  echo "    encumbrance: $encumbrance"
  echo "    mass: $mass"
  echo "    cost: $cost"
  create_misc "$name" "$description" "$encumbrance" "$mass" "$cost"
done < "$filename"