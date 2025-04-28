package main

import (
    "encoding/json"
    "fmt"
    "log"
    "math/rand"
    "net/http"
    "strconv"
)

type RecommendationResponse struct {
    UserID           int     `json:"user_id"`
    RecommendedItems []int   `json:"recommended_items"`
}

func recommendHandler(w http.ResponseWriter, r *http.Request) {
    userIDStr := r.URL.Path[len("/recommend/"):] 
    userID, err := strconv.Atoi(userIDStr)
    if err != nil {
        http.Error(w, "Invalid user ID", http.StatusBadRequest)
        return
    }

    postIDs := []int{101, 102, 103, 104, 105}
    rand.Shuffle(len(postIDs), func(i, j int) { postIDs[i], postIDs[j] = postIDs[j], postIDs[i] })
    recommendedItems := postIDs[:3]

    response := RecommendationResponse{
        UserID:           userID,
        RecommendedItems: recommendedItems,
    }
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(response)
}

func main() {
    http.HandleFunc("/recommend/", recommendHandler)
    fmt.Println("Server is running on port 8080...")
    log.Fatal(http.ListenAndServe(":8080", nil))
}





