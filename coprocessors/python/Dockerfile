FROM python:3-alpine

WORKDIR /app

COPY requirements.txt ./

RUN python -m pip install -r requirements.txt

COPY . .

CMD python main.py
