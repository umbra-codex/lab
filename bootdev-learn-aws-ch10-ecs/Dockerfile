FROM python:slim
WORKDIR /app
RUN pip install --no-cache-dir boto3
COPY app.py /app/app.py
ENTRYPOINT ["python", "/app/app.py"]

