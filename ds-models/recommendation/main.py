import pandas as pd
import numpy as np
import tensorflow as tf
from sklearn.model_selection import train_test_split

data = {
    "user_id": [1, 2, 3, 4, 5],
    "post_id": [101, 102, 103, 104, 105],
    "interaction_type": ["like", "share", "repost", "comment", "search"],
}

df = pd.DataFrame(data)
interaction_amount = {
    "like": 4,
    "repost": 3,
    "comment": 2,
    "share": 3,
    "search": 1,
}

df["interaction_value"] = df["interaction_type"].map(interaction_amount)

df["user_id"] = df["user_id"].astype("category").cat.codes
df["post_id"] = df["post_id"].astype("category").cat.codes

train, test = train_test_split(df, test_size=0.2, random_state=42)
embedding_dim = 16
user_input = tf.keras.layers.Input(shape=(1,))
post_input = tf.keras.layers.Input(shape=(1,))
user_embedding = tf.keras.layers.Embedding(
    input_dim=len(df["user_id"].unique()) + 1, output_dim=embedding_dim
)(user_input)
post_embedding = tf.keras.layers.Embedding(
    input_dim=len(df["post_id"].unique()) + 1, output_dim=embedding_dim
)(post_input)

user_flat = tf.keras.layers.Flatten()(user_embedding)
post_flat = tf.keras.layers.Flatten()(post_embedding)

concat = tf.keras.layers.Concatenate()([user_flat, post_flat])
dense1 = tf.keras.layers.Dense(64, activation="relu")(concat)
dense2 = tf.keras.layers.Dense(32, activation="relu")(dense1)
output = tf.keras.layers.Dense(1, activation="sigmoid")(dense2)

model = tf.keras.Model(inputs=[user_input, post_input], outputs=output)
model.compile(optimizer="adam", loss="binary_crossentropy", metrics=["accuracy"])
print("Columns in train dataset:", train.columns)
model.fit(
    [train["user_id"], train["post_id"]],
    train["interaction_value"],
    epochs=10,
    batch_size=16,
)
model.save("recommendation_model.h5")

print("Model training complete and saved!")
