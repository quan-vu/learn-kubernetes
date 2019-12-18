from fastapi import FastAPI

app = FastAPI()


@app.get("/")
def read_root():
    return {"Message": "demo-python-fastapi-app"}


@app.get("/items/{item_id}")
def read_item(item_id: int, q: str = None):
    return {"item_id": item_id, "q": q}