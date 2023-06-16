set -e

tag=$1

if [ -z "$tag" ]
then
  echo "No tag supplied, defaulting to latest"
  tag="latest"
fi

docker build ./ -f Dockerfile -t dag:$tag
