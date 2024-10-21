#!/bin/bash


create_trait_effect() {
  name=$1
  description=$2
  effects=$3
  yaml=$4
  filename=$(echo "$name" | tr '[:upper:]' '[:lower:]' | tr ' ' '_').go
  structName=$(echo "$name" | tr -d ' ' | tr -d '&' | tr -d "'" | tr -d ';' | tr -d '!' | tr -d '.' | tr -d '?' | tr -d '-')
  echo "$name, $filename, $structName"
cat <<EOL > "../server/internal/domain/effects/$filename"
package effects

import (
	"github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type $structName struct {
  Name string
  Description string
}

func (e *$structName) Apply(state domain.State) domain.State {
  // $effects
  log.Println("applying $name")
  return state
}
EOL
echo "goType: $structName" >> "$yaml"
}

echo "generating trait effects"
traits=$(ls -l traits/*.yaml | awk '{print $9}')
for trait in $traits; do
  echo "processing $trait"
  name=$(yq .name < $trait)
  description=$(yq .description < $trait)
  effects=$(yq .effects < $trait)
  create_trait_effect "$name" "$description" "$effects" "$trait"
done
