import json

def save_posts(posts, filename):
    with open(filename, "w", encoding="utf-8") as file:
        json.dump(posts, file, indent=4)
    print(f"All scraped posts are saved in {filename}.")