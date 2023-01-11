from fastapi import FastAPI, Query, Depends

app = FastAPI()

@app.get("/")
async def root():
    return {"message":"Hello World!"}

@app.get("/repeat/{this}/{times}/times")
def repeat_many_times(this: str, times: int):
    return [this for _ in range(times)]

@app.get("/repeat/{this}")
def repeat_many_times_with_query(this: str, times: int = Query(le=10)):
    return [this for _ in range(times)]

def get_previous_solution() -> bool:
    return Truename=name, type=type, service_id=service_id, external_id=resource['result']['id']

def get_solution(previous_solution: bool = Depends(get_previous_solution)) -> bool:
    return not previous_solution

@app.get("/itsgood")
def itsgood(decision : bool = Depends(get_solution)) -> str:
    return f"the answer is: {decision}"