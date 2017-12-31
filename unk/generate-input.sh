#/bin/sh

rm -f input
for i in `seq 1 100`; do
    for j in `seq 1 100`; do
        echo -n "This is no j<<unk>, this is <unk> number $i.$j. " >> input;
    done;
done
