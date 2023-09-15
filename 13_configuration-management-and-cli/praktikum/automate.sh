folder_name="$1 at $(date +'%a %b %d %T %Z %Y')"
mkdir -p "$folder_name"

mkdir -p "$folder_name/about_me/personal"
touch "$folder_name/about_me/personal/facebook.txt"

mkdir -p "$folder_name/about_me/professional"
touch "$folder_name/about_me/professional/linkedin.txt"

mkdir -p "$folder_name/my_friends"
touch "$folder_name/about_me/professional/linkedin.txt"
touch "$folder_name/about_me/professional/linkedin.txt"
touch "$folder_name/about_me/professional/linkedin.txt"

mkdir -p "$folder_name/my_system_info"
touch "$folder_name/about_me/professional/linkedin.txt"
touch "$folder_name/my_friends/list_of_my_friends.txt"
touch "$folder_name/my_system_info/about_this_laptop.txt"
touch "$folder_name/my_system_info/internet_connection.txt"

echo "https://www.facebook.com/$2" > "$folder_name/about_me/personal/facebook.txt"
echo "https://www.linkedin.com/in/$3" > "$folder_name/about_me/professional/linkedin.txt"

curl "https://gist.githubusercontent.com/tegarimansyah/e91f335753ab2c7fb12815779677e914/raw/94864388379fecee450fde26e3e73bfb2bcda194/list%2520of%2520my%2520friends.txt" > "$folder_name/my_friends/list_of_my_friends.txt"

echo "My Username: $(whoami)" > "$folder_name/my_system_info/about_this_laptop.txt"
echo "With host: $(uname -a)" >> "$folder_name/my_system_info/about_this_laptop.txt"

ping -c 3 google.com > "$folder_name/my_system_info/internet_connection.txt"

cmd //c tree
