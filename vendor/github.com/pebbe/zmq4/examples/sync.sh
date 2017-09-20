#!/bin/sh
echo "Starting subscribers..."
for i in 1 2 3 4 5 6 7 8 9 10
do
    ./syncsub &
done
echo "Starting publisher..."
./syncpub
# have all subscribers finished?
sleep 1
echo Still running instances of syncsub:
ps | grep syncsub
