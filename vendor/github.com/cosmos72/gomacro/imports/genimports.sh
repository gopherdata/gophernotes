exec > genimports.gomacro

echo "#!/usr/bin/env gomacro"
echo

find /usr/local/go/src -type d | \
  sed -e 's,/usr/local/go/src/,,' -e 's,/usr/local/go/src,,' | \
  grep "[a-z]" | grep -v 'cmd\|internal\|testdata\|vendor' | \
  sort |
while read i; do
  echo "import _b \"$i\""
done

