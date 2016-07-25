
F_L=$(find ./ -maxdepth 1 -mindepth 1 -type d)

for D in $F_L; do
    echo $D;
    sh -c "cd $D; sh install.sh"
done
