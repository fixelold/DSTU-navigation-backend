search_dir="./migration_up/"
migration_script="./migration.sh"
for entry in "$search_dir"*
do 
    /bin/zsh "$migration_script" "$entry"
done