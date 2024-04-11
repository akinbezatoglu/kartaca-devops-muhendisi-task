# syntax=docker/dockerfile:1.4
FROM python:3.12.3-alpine3.19

WORKDIR /app

COPY requirements.txt ./
RUN --mount=type=cache,target=/root/.cache/pip \
    pip3 install -r requirements.txt

COPY *.py ./

EXPOSE 4444

ENTRYPOINT ["python3"]
CMD ["app.py"]