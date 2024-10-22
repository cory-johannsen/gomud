#!/bin/bash


create_effect() {
  name=$1
  description=$(echo "$2" | tr -d '\n' | tr '"' "'")
  effect=$(echo "$3" | tr -d '\n' | tr '"' "'")
  structName=$4
  varName=$5
  filename=$(echo "$name" | tr '[:upper:]' '[:lower:]' | tr ' ' '_').go
  echo "$name, $filename, $structName"
cat <<EOL > "../server/internal/domain/effect/$filename"
package effect

import (
  "github.com/cory-johannsen/gomud/internal/domain"
  log "github.com/sirupsen/logrus"
)

type $structName struct {
  name string
  description string
}

func New$structName() *$structName {
  return &$structName{
    name: "$name",
    description: "$effect",
  }
}

func (e *$structName) Name() string {
  return e.name
}

func (e *$structName) Description() string {
  return e.description
}

func (e *$structName) Applier() domain.Applier {
  return e.Apply
}

func (e *$structName) Apply(state domain.State) domain.State {
  // $effects
  log.Println("applying $name")
  return state
}

var _ domain.Effect = &$structName{}
EOL
#echo "goType: $structName" >> "$yaml"
  echo "  New$structName," >> "../server/internal/domain/effect/wireset.go"

  echo "  $varName *$structName," >> "../server/internal/domain/effect/effect.go"
}


cat <<EOL > "../server/internal/domain/effect/wireset.go"
// +build wireinject

package effect

import (
	"github.com/google/wire"
)

var EffectsSet = wire.NewSet(
EOL


cat <<EOL > "../server/internal/domain/effect/effect.go"
package effect

import (
	"github.com/cory-johannsen/gomud/internal/domain"
)


func NewEffects (
EOL

varNames=()

function process() {
    name=$(yq .name < "$1")
    description=$(yq .description < "$1")
    effects=$(yq .effects < "$1" | tr -d '\n' | tr '"' "'" )
    structName=$(echo "$name" | tr -d ' ' | tr -d '&' | tr -d "'" | tr -d ';' | tr -d '!' | tr -d '.' | tr -d '?' | tr -d '-' | tr -d '\n' | tr -d "’" | tr -d '(' | tr -d ')')
    varName=$(echo "${structName:0:1}" | tr '[:upper:]' '[:lower:]')${structName:1}

    create_effect "$name" "$description" "${effects:2}" "$structName" "$varName"
    varNames+=("$varName")
}

echo "generating trait effects"
traits=$(ls -l traits/*.yaml | awk '{print $9}')
for trait in $traits; do
  process "$trait"
done

echo "generating talent effects"
talents=$(ls -l talents/*.yaml | awk '{print $9}')
for talent in $talents; do
  process "$talent"
done

echo "generating injury effects"
injuries=$(ls -l injuries/moderate/*.yaml | awk '{print $9}')
for injury in $injuries; do
  process "$injury"
done
injuries=$(ls -l injuries/serious/*.yaml | awk '{print $9}')
for injury in $injuries; do
  process "$injury"
done
injuries=$(ls -l injuries/grievous/*.yaml | awk '{print $9}')
for injury in $injuries; do
  process "$injury"
done

echo "generating drawback effects"
drawbackCount=$(cat appearance/drawbacks.yaml | yq '. | length')
for i in $(seq 0 $(($drawbackCount - 1))); do
  name=$(cat appearance/drawbacks.yaml | yq .["$i"].name)
  description=$(cat appearance/drawbacks.yaml | yq .["$i"].description)
  effect=$(cat appearance/drawbacks.yaml | yq .["$i"].effect | tr -d '\n' | tr '"' "'")
  structName=$(echo "$name" | tr -d ' ' | tr -d '&' | tr -d "'" | tr -d ';' | tr -d '!' | tr -d '.' | tr -d '?' | tr -d '-' | tr -d '\n'| tr -d "’" | tr -d '(' | tr -d ')')
  varName=$(echo "${structName:0:1}" | tr '[:upper:]' '[:lower:]')${structName:1}

  create_effect "$name" "$description" "${effect}" "$structName" "$varName"
  varNames+=("$varName")
done

echo "generating item quality effects"
qualities=$(ls -l equipment/qualities/*.yaml | awk '{print $9}')
for quality in $qualities; do
  process "$quality"
done

# complete the wireset

echo ")" >> "../server/internal/domain/effect/wireset.go"

cat <<EOL >> "../server/internal/domain/effect/effect.go"
) domain.Effects {
  return domain.Effects{
EOL

for varName in "${varNames[@]}"; do
  echo "    $varName," >> "../server/internal/domain/effect/effect.go"
done
echo "  }" >> "../server/internal/domain/effect/effect.go"
echo "}" >> "../server/internal/domain/effect/effect.go"