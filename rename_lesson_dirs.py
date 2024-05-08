import os

for folder_name in os.listdir():
    if folder_name.startswith("l") and os.path.isdir(folder_name):
        print("moving " + folder_name)
        new_folder_name = folder_name[1:]
        os.rename(folder_name, new_folder_name)

print("Renamed successfully.")
