from user_input import get_usernames
from scraper import scrape_instagram_posts
from storage import save_posts
import json


def main():
    print("Instagram scraper starting..")
    insta_users = get_usernames()

    if not insta_users:
        print("Please enter username")
        return

    all_posts = scrape_instagram_posts(insta_users)

    if not all_posts:
        print("No posts found.")
        return

    print("\n You scraped :")
    print(json.dumps(all_posts, indent=4))
    save_posts(all_posts, "instagram_scraped_posts.json")
    print("Posts saved.")


if __name__ == "__main__":
    main()
