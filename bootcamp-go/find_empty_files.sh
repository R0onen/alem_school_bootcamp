print_usage() {
  echo "usage: find_empty_files.sh dir_path"
}

# Check if an argument is provided
if [ -z "$1" ]; then
  print_usage
  exit 0
fi

# Check if the provided argument exists and is a directory
if [ ! -e "$1" ]; then
  echo "error: directory $1 not found"
  exit 0
elif [ ! -d "$1" ]; then
  echo "error: $1 is not a directory"
  exit 0
fi

# Find empty files or files with only spaces
find "$1" -type f -exec sh -c '
  for file do
    if [ ! -s "$file" ]; then
      echo "$file"
    elif [ -z "$(grep -v "^[[:space:]]*$" "$file")" ]; then
      echo "$file"
    fi
  done
' sh {} +

exit 0
