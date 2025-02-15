from fastapi import FastAPI
import tensorflow as tf
import numpy as np
import pandas as pd
from sklearn.preprocessing import LabelEncoder

model = tf.keras.models.load_model("recommendation_model.h5")

df = pd.DataFrame({
    'post_id': [101, 102, 103, 104, 105]
})

post_encoder = LabelEncoder()
df['encoded_post_id'] = post_encoder.fit_transform(df['post_id'])

app = FastAPI()

@app.get("/recommend/{user_id}")
async def recommend(user_id: int):
    print(f"Received request for user_id: {user_id}")
    post_ids = np.array(df['encoded_post_id'])
    user_ids = np.array([user_id] * len(post_ids))
    print("Encoded Post IDs:", post_ids)
    print("User IDs:", user_ids)
    predictions = model.predict([user_ids, post_ids])
    recommended_indices = np.argsort(predictions[:, 0])[-5:][::-1]
    recommended_items = df.iloc[recommended_indices]['post_id'].tolist()

    return {"user_id": user_id, "recommended_items": recommended_items}





