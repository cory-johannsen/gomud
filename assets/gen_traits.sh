#!/bin/bash
traits=(
  "Blessing in Disguise"
  "Danger Sense"
  "Dauntless"
  "Esoteric Memory"
  "Fortune's Wheel"
  "Grim Resolve"
  "Manifest Destiny"
  "Mixed Family"
  "Mountain Amongst Men"
  "Natural Selection"
  "Noble Savage"
  "Seventh Sense"
  "Cavesight"
  "Children of the Earth"
  "Consume Alcohol"
  "Mountain Warfare"
  "Grudgebearer"
  "Ironclad"
  "Oathkeeper"
  "Physical Prowess"
  "Rune-Marked Glory"
  "Stentorian Voice"
  "Stoneheaded"
  "Strength of the Mountain"
  "Clockworks of War"
  "Crag Fighting"
  "Denizen of Stone"
  "Dungeons Deep"
  "Escape Artist"
  "Goldbergian"
  "Hocus Pocus"
  "Metropolitan"
  "Thieving Ways"
  "Tunnel Vision"
  "Underfoot"
  "Wretched Prankster"
  "Beguiler"
  "Cat-like Reflexes"
  "Craven"
  "Farsight"
  "Fettered Chaos"
  "Fieldwarden"
  "Hijinks"
  "Kleptomania"
  "Low Blow"
  "Memento"
  "Pintsized"
  "Bewitching"
  "Beyond the Veil"
  "Deadly Aim"
  "Enduring Mortality"
  "In Crowd Treachery"
  "Firstborn"
  "Meditative Healing"
  "Nature's Own"
  "Nighteyes"
  "Warrior's Tattoo"
  "Broad Bellied"
  "Cast Iron Stomach"
  "Cruisin' for a Bruisin'"
  "Frightening Bellow"
  "Gut-plate"
  "Hunger Pangs"
  "Mighty Thews"
  "Odd Couple"
  "Rotgut Spray"
  "Slamdance"
  "Thick Lining"
  "Wendigo"
)

create_background_trait() {
  filename=$(echo "$1" | tr '[:upper:]' '[:lower:]' | tr ' ' '_').yaml
cat <<EOL > "traits/$filename"
name: $1
description: "This is a description of the trait."
effects:
  - Effect1
EOL
}

create_job_trait() {
  filename=$(echo "$1" | tr '[:upper:]' '[:lower:]' | tr ' ' '_').yaml
  echo "  creating $filename"
cat <<EOL > "traits/$filename"
name: $1
description: $2
effects:
  - $3
EOL
}

rewrite_job() {
  filename=$1
  job=$2
  archetype=$3
  description=$4
  tier=$5
  rewardPoints=$6
  trait=$7
  echo "  rewriting $job"
cat <<EOL > "rw/$filename"
name: $job
archetype: $archetype
description: $description
ties: $tier
reward_point_cost: $rewardPoints
traits:
  - $trait
EOL
}

PROCESS_BACKGROUNDS="false"
PROCESS_JOBS="true"
PROCESS_TEAMS="false"

if [[ $PROCESS_BACKGROUNDS == "true" ]]; then
echo "generating background traits"
for trait in "${traits[@]}"; do
  echo "creating $trait"
  filename=$(echo "$trait" | tr '[:upper:]' '[:lower:]' | tr ' ' '_')
  create_background_trait "$trait"
done
fi

if [[ $PROCESS_JOBS == "true" ]]; then
echo "generating job traits"
jobs=$(ls -l jobs/*.yaml | awk '{print $9}')
for job in $jobs; do
  echo "processing $job"
  jobName=$(yq .name < $job)
  jobArechtype=$(yq .archetype < $job)
  jobDescription=$(yq .description < $job)
  jobTier=$(yq .tier < $job)
  jobRewardPoints=$(yq .reward_point_cost < $job)

  traitName=$(yq .trait.name < $job)
  traitDescription=$(yq .trait.description < $job)
  traitEffect=$(yq .trait.effect < $job)
  create_job_trait "$traitName" "$traitDescription" "$traitEffect"
  rewrite_job "$job" "$jobName" "$jobArechtype" "$jobDescription" "$jobTier" "$jobRewardPoints" "$traitName"
done
fi


if [[ $PROCESS_TEAMS == "true" ]]; then
echo "generating team traits"
teams=$(ls -l teams/*.yaml | awk '{print $9}')
for team in $teams; do
  echo "processing $team"
  if [[ "$team" == "*.tmpl.yaml" ]]; then
    echo "skipping template"
    continue
  fi
  teamTraits=$(yq .traits < "$team")
  echo "team traits: $teamTraits"
  for trait in $teamTraits; do
    echo "  $trait"
#    traitName=$(echo $trait | tr -d '"')
#    traitDescription=$(yq ."$trait".description < $trait)
#    traitEffect=$(yq ."$trait".effect < $trait)
#    create_job_trait "$traitName" "$traitDescription" "$traitEffect"
  done
done
fi
