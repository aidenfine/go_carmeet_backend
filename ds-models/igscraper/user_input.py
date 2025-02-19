def get_usernames():
    user_input = input("Enter Instagram usernames (comma-separated): ")
    usernames = [user.strip() for user in user_input.split(',') if user.strip()]
    
    if not usernames:
        print("Please enter a valid username")
    
    return usernames