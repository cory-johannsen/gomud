#!/bin/bash

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
