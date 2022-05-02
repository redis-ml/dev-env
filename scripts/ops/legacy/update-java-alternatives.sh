for f in `find /etc/alternatives/ -name "*.*" -prune -o -name "j*"  -print`; do
    f=$(basename $f)
    sudo update-alternatives --config $f
done
