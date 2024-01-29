FROM python:3.8
WORKDIR /app
COPY requirements.txt requirements.txt
COPY app.py app.py 
RUN pip install -r requirements.txt
EXPOSE 8000
CMD ["python","app.py"]