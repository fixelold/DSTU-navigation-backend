search_dir="./migration_up/"
migration_script="./migration.sh"
for entry in "$search_dir"*
do 
    /bin/bash "$migration_script" "$entry"
done
