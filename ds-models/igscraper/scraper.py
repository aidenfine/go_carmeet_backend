from instagrapi import Client


def scrape_instagram_posts(usernames, limit=6):
    client = Client()
    username = input("Enter your Instagram username? ")
    password = input("your Instagram password? ")

    try:
        client.login(username, password)
        print("Logged in.")
    except Exception as e:
        print(f"login failed: {e}")
        return []

    all_posts = []
    for user in usernames:
        print(f"Checking out @{user}...")
        try:
            user_id = client.user_id_from_username(user)
            posts = client.user_medias(user_id, limit)
            for post in posts:
                all_posts.append(
                    {
                        "username": user,
                        "caption": post.caption_text or "No caption found.",
                        "image_url": str(post.thumbnail_url)
                        if post.thumbnail_url
                        else "No image found.",
                    }
                )
            print(f"Scraped {len(posts)} posts from @{user}.")
        except Exception:
            print(f"@{user} Failed. Skipping...")

    return all_posts
