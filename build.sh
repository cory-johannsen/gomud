#!/bin/zsh

base_dir=$(pwd)
gcp_project_id=big-formula-443619-k0
git_sha=$(git rev-parse --short HEAD)

goarch=$GOARCH
if [ -z "$goarch" ]; then
  goarch=amd64
fi

echo "Building for $goarch"

rm -rf $base_dir/bin
mkdir -p $base_dir/bin
cd $base_dir/server/cmd
GOARCH="$goarch" go build -o $base_dir/bin/gunchete."$goarch" .
cp $base_dir/bin/gunchete."$goarch" $base_dir/bin/gunchete
cd $base_dir

# Note: you must be logged into GCP artifact registry with a service account that is authorized to push
docker build -t gunchete."$goarch" .

sha_tag=us-west1-docker.pkg.dev/"$gcp_project_id"/gunchete/gunchete."$goarch":"$git_sha"
latest_tag=us-west1-docker.pkg.dev/"$gcp_project_id"/gunchete/gunchete."$goarch":latest

docker tag gunchete."$goarch" $sha_tag
docker tag gunchete."$goarch" $latest_tag
docker push $sha_tag
docker push $latest_tag

echo "pushed $latest_tag"